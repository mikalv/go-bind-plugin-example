package main

// Autogenerated by github.com/wendigo/go-bind-plugin on 2016-11-10 15:22:32.49165816 +0100 CET, do not edit!
// Command: go-bind-plugin -plugin-path plugin.so -plugin-package ./plugin -output-name PluginAPI -output-path plugin_api.go -output-package main -dereference-vars -rebuild
//
// Plugin plugin.so info:
// - package: github.com/wendigo/go-bind-plugin-example/plugin
// - size: 2334896 bytes
// - sha256: 43eb3d7d8694a213141b793f200e34f44d18de6da56b2001b13621a8575bf1f7

import (
	"fmt"
	"plugin"
	"reflect"
	"strings"
)

// PluginAPI wraps symbols (functions and variables) exported by plugin github.com/wendigo/go-bind-plugin-example/plugin
//
// See docs at https://godoc.org/github.com/wendigo/go-bind-plugin-example/plugin
type PluginAPI struct {
	// Exported functions
	_CalculateSin func(float64) float64
	_SayHello     func(string)

	// Exported variables (public references)

	// See docs at https://godoc.org/github.com/wendigo/go-bind-plugin-example/plugin#CurrentYear
	CurrentYear int
}

// CalculateSin function was exported from plugin github.com/wendigo/go-bind-plugin-example/plugin symbol 'CalculateSin'
//
// See docs at https://godoc.org/github.com/wendigo/go-bind-plugin-example/plugin#CalculateSin
func (p *PluginAPI) CalculateSin(in0 float64) float64 {
	return p._CalculateSin(in0)
}

// SayHello function was exported from plugin github.com/wendigo/go-bind-plugin-example/plugin symbol 'SayHello'
//
// See docs at https://godoc.org/github.com/wendigo/go-bind-plugin-example/plugin#SayHello
func (p *PluginAPI) SayHello(in0 string) {
	p._SayHello(in0)
}

// String returnes textual representation of the wrapper. It provides info on exported symbols and variables.
func (p *PluginAPI) String() string {
	var lines []string
	lines = append(lines, "Struct PluginAPI:")
	lines = append(lines, "\t- Generated on: 2016-11-10 15:22:32.49165816 +0100 CET")
	lines = append(lines, "\t- Command: go-bind-plugin -plugin-path plugin.so -plugin-package ./plugin -output-name PluginAPI -output-path plugin_api.go -output-package main -dereference-vars -rebuild")
	lines = append(lines, "\nPlugin info:")
	lines = append(lines, "\t- package: github.com/wendigo/go-bind-plugin-example/plugin")
	lines = append(lines, "\t- sha256 sum: 43eb3d7d8694a213141b793f200e34f44d18de6da56b2001b13621a8575bf1f7")
	lines = append(lines, "\t- size: 2334896 bytes")
	lines = append(lines, "\nExported functions (2):")
	lines = append(lines, "\t- CalculateSin func(float64) (float64)")
	lines = append(lines, "\t- SayHello func(string)")

	lines = append(lines, "\nExported variables (1):")
	lines = append(lines, "\t- CurrentYear int")

	return strings.Join(lines, "\n")
}

// BindPluginAPI loads plugin from the given path and binds symbols (variables and functions)
// to the PluginAPI struct. All variables are derefenences.
func BindPluginAPI(path string) (*PluginAPI, error) {
	p, err := plugin.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Could not open plugin: %s", err)
	}

	ret := new(PluginAPI)

	funcCalculateSin, err := p.Lookup("CalculateSin")
	if err != nil {
		return nil, fmt.Errorf("Could not import function 'CalculateSin', symbol not found: %s", err)
	}

	if typed, ok := funcCalculateSin.(func(float64) float64); ok {
		ret._CalculateSin = typed
	} else {
		return nil, fmt.Errorf("Could not import function 'CalculateSin', incompatible types 'func(float64) (float64)' and '%s'", reflect.TypeOf(funcCalculateSin))
	}

	funcSayHello, err := p.Lookup("SayHello")
	if err != nil {
		return nil, fmt.Errorf("Could not import function 'SayHello', symbol not found: %s", err)
	}

	if typed, ok := funcSayHello.(func(string)); ok {
		ret._SayHello = typed
	} else {
		return nil, fmt.Errorf("Could not import function 'SayHello', incompatible types 'func(string)' and '%s'", reflect.TypeOf(funcSayHello))
	}

	varCurrentYear, err := p.Lookup("CurrentYear")
	if err != nil {
		return nil, fmt.Errorf("Could not import variable 'CurrentYear', symbol not found: %s", err)
	}

	if typed, ok := varCurrentYear.(*int); ok {
		ret.CurrentYear = *typed
	} else {
		return nil, fmt.Errorf("Could not import variable 'CurrentYear', incompatible types 'int' and '%s'", reflect.TypeOf(varCurrentYear))
	}

	return ret, nil
}
