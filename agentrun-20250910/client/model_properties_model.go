// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelProperties interface {
	dara.Model
	String() string
	GoString() string
	SetContextSize(v int32) *ModelProperties
	GetContextSize() *int32
}

type ModelProperties struct {
	ContextSize *int32 `json:"contextSize,omitempty" xml:"contextSize,omitempty"`
}

func (s ModelProperties) String() string {
	return dara.Prettify(s)
}

func (s ModelProperties) GoString() string {
	return s.String()
}

func (s *ModelProperties) GetContextSize() *int32 {
	return s.ContextSize
}

func (s *ModelProperties) SetContextSize(v int32) *ModelProperties {
	s.ContextSize = &v
	return s
}

func (s *ModelProperties) Validate() error {
	return dara.Validate(s)
}
