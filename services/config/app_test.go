/*
//
//  ______    _   _           _  __      _   _     ____   ___
// |  ____|  | | | |         | |/ _|    | | | |   |___ \ / _ \
// | |__ __ _| |_| |__   __ _| | |_ __ _| |_| |__   __) | | | |
// |  __/ _` | __| '_ \ / _` | |  _/ _` | __| '_ \ |__ <| | | |
// | | | (_| | |_| | | | (_| | | || (_| | |_| | | |___) | |_| |
// |_|  \__,_|\__|_| |_|\__,_|_|_| \__,_|\__|_| |_|____/ \___/
//
// Written by Fathalfath30.
// Email : fathalfath30@gmail.com
// Follow me on:
//  Github : https://github.com/fathalfath30
//  Gitlab : https://gitlab.com/Fathalfath30
//
*/

package config_test

import (
	"github.com/stretchr/testify/require"
	"os"
	"scum-helper/services/config"
	"scum-helper/services/helper"
	configTestData "scum-helper/services/test_data/config"
	"testing"
)

func Test_ItCanCreateNewApp(t *testing.T) {
	t.Run("test it can get host and port", func(t *testing.T) {
		defaultHost := os.Getenv(helper.AppHostEnvName)
		defaultPort := os.Getenv(helper.AppPortEnvName)

		// set APP_HOST and APP_PORT with test value
		_ = os.Setenv(helper.AppHostEnvName, configTestData.ApplicationTestHost)
		_ = os.Setenv(helper.AppPortEnvName, configTestData.ApplicationTestPort)

		app := config.NewApp()
		require.Equal(t, configTestData.ApplicationTestHost, app.GetHost())
		require.Equal(t, configTestData.ApplicationTestPort, app.GetPort())

		// reset APP_HOST and APP_PORT with their default value
		_ = os.Setenv(helper.AppPortEnvName, defaultHost)
		_ = os.Setenv(helper.AppPortEnvName, defaultPort)
	})

	t.Run("it should set default value with ip address if host is not set", func(t *testing.T) {
		defaultHost := os.Getenv(helper.AppHostEnvName)
		defaultPort := os.Getenv(helper.AppPortEnvName)

		// set APP_HOST and APP_PORT with test value
		_ = os.Setenv(helper.AppHostEnvName, "")
		_ = os.Setenv(helper.AppPortEnvName, configTestData.ApplicationTestPort)

		app := config.NewApp()
		require.Equal(t, helper.ApplicationDefaultHost, app.GetHost())
		require.Equal(t, configTestData.ApplicationTestPort, app.GetPort())

		// reset APP_HOST and APP_PORT with their default value
		_ = os.Setenv(helper.AppPortEnvName, defaultHost)
		_ = os.Setenv(helper.AppPortEnvName, defaultPort)
	})

	t.Run("it should panic if port is empty", func(t *testing.T) {
		defaultPort := os.Getenv(helper.AppPortEnvName)

		// set APP_HOST and APP_PORT with test value
		_ = os.Setenv(configTestData.ApplicationTestHost, configTestData.ApplicationTestHost)
		require.Panics(t, func() {
			config.NewApp()
		})

		// reset APP_PORT with their default value
		_ = os.Setenv(helper.AppPortEnvName, defaultPort)
	})
}
