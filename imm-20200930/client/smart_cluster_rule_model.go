// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartClusterRule interface {
	dara.Model
	String() string
	GoString() string
	SetKeywords(v []*string) *SmartClusterRule
	GetKeywords() []*string
	SetSensitivity(v float32) *SmartClusterRule
	GetSensitivity() *float32
}

type SmartClusterRule struct {
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	// example:
	//
	// 0.5
	Sensitivity *float32 `json:"Sensitivity,omitempty" xml:"Sensitivity,omitempty"`
}

func (s SmartClusterRule) String() string {
	return dara.Prettify(s)
}

func (s SmartClusterRule) GoString() string {
	return s.String()
}

func (s *SmartClusterRule) GetKeywords() []*string {
	return s.Keywords
}

func (s *SmartClusterRule) GetSensitivity() *float32 {
	return s.Sensitivity
}

func (s *SmartClusterRule) SetKeywords(v []*string) *SmartClusterRule {
	s.Keywords = v
	return s
}

func (s *SmartClusterRule) SetSensitivity(v float32) *SmartClusterRule {
	s.Sensitivity = &v
	return s
}

func (s *SmartClusterRule) Validate() error {
	return dara.Validate(s)
}
