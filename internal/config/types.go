// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package config

const (
	defaultRefreshRate  = 2
	defaultMaxConnRetry = 5

	// CPU tracks cpu usage.
	CPU = "cpu"

	// MEM tracks memory usage.
	MEM = "memory"
)

// ClockConfig represents a single timezone clock shown in the header.
type ClockConfig struct {
	// Timezone is an IANA timezone name, e.g. "America/New_York".
	Timezone string `json:"timezone" yaml:"timezone"`

	// Label is an optional short display label. Defaults to the timezone name.
	Label string `json:"label" yaml:"label,omitempty"`
}

// UI tracks ui specific configs.
type UI struct {
	// EnableMouse toggles mouse support.
	EnableMouse bool `json:"enableMouse" yaml:"enableMouse"`

	// Headless toggles top header display.
	Headless bool `json:"headless" yaml:"headless"`

	// LogoLess toggles k9s logo.
	Logoless bool `json:"logoless" yaml:"logoless"`

	// Crumbsless toggles nav crumb display.
	Crumbsless bool `json:"crumbsless" yaml:"crumbsless"`

	// Splashless disables the splash screen on startup.
	Splashless bool `json:"splashless" yaml:"splashless"`

	// Reactive toggles reactive ui changes.
	Reactive bool `json:"reactive" yaml:"reactive"`

	// NoIcons toggles icons display.
	NoIcons bool `json:"noIcons" yaml:"noIcons"`

	// Invert inverts all skin colors using Oklch lightness inversion.
	Invert bool `json:"invert" yaml:"invert"`

	// Skin reference the general k9s skin name.
	// Can be overridden per context.
	Skin string `json:"skin" yaml:"skin,omitempty"`

	// Clocks defines optional timezone clocks shown in the header when logoless is true.
	Clocks []ClockConfig `json:"clocks" yaml:"clocks,omitempty"`

	// DefaultsToFullScreen toggles fullscreen on views like logs, yaml, details.
	DefaultsToFullScreen bool `json:"defaultsToFullScreen" yaml:"defaultsToFullScreen"`

	// UseFullGVRTitle toggles the display of full GVR (group/version/resource) vs R in views title.
	UseFullGVRTitle bool `json:"useFullGVRTitle" yaml:"useFullGVRTitle"`

	manualHeadless   *bool
	manualLogoless   *bool
	manualCrumbsless *bool
	manualSplashless *bool
	manualInvert     *bool
}
