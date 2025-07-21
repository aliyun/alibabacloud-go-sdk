// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOtaAutoStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientType(v int32) *GetDeviceOtaAutoStatusRequest
	GetClientType() *int32
}

type GetDeviceOtaAutoStatusRequest struct {
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
}

func (s GetDeviceOtaAutoStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaAutoStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaAutoStatusRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *GetDeviceOtaAutoStatusRequest) SetClientType(v int32) *GetDeviceOtaAutoStatusRequest {
	s.ClientType = &v
	return s
}

func (s *GetDeviceOtaAutoStatusRequest) Validate() error {
	return dara.Validate(s)
}
