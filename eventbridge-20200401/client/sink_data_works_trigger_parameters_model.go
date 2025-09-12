// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSinkDataWorksTriggerParameters interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *SinkDataWorksTriggerParameters
	GetEnable() *string
}

type SinkDataWorksTriggerParameters struct {
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s SinkDataWorksTriggerParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkDataWorksTriggerParameters) GoString() string {
	return s.String()
}

func (s *SinkDataWorksTriggerParameters) GetEnable() *string {
	return s.Enable
}

func (s *SinkDataWorksTriggerParameters) SetEnable(v string) *SinkDataWorksTriggerParameters {
	s.Enable = &v
	return s
}

func (s *SinkDataWorksTriggerParameters) Validate() error {
	return dara.Validate(s)
}
