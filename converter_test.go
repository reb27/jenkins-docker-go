package main

import "testing"

func TestKilometersToMiles(t *testing.T) {
    result := KilometersToMiles(1.0)
    expected := 0.6213
    if result != expected {
        t.Errorf("Expected %.6f, got %.6f", expected, result)
    }
}

func TestMilesToKilometers(t *testing.T) {
    result := MilesToKilometers(0.621371)
    expected := 0.0
    if result != expected {
        t.Errorf("Expected %.6f, got %.6f", expected, result)
    }
}
