package robot_in_circle_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestRobotInCircle(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "robot_in_circle Suite")
}
