// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFalconSeekRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *UpdateFalconSeekRequest
	GetEnable() *bool
}

type UpdateFalconSeekRequest struct {
	// This parameter is required.
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s UpdateFalconSeekRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFalconSeekRequest) GoString() string {
	return s.String()
}

func (s *UpdateFalconSeekRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateFalconSeekRequest) SetEnable(v bool) *UpdateFalconSeekRequest {
	s.Enable = &v
	return s
}

func (s *UpdateFalconSeekRequest) Validate() error {
	return dara.Validate(s)
}
