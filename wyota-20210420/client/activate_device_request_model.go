// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuid(v string) *ActivateDeviceRequest
	GetUuid() *string
}

type ActivateDeviceRequest struct {
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ActivateDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateDeviceRequest) GoString() string {
	return s.String()
}

func (s *ActivateDeviceRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ActivateDeviceRequest) SetUuid(v string) *ActivateDeviceRequest {
	s.Uuid = &v
	return s
}

func (s *ActivateDeviceRequest) Validate() error {
	return dara.Validate(s)
}
