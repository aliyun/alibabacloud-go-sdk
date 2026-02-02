// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataChannelCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *GetDataChannelCredentialRequest
	GetBusinessUnitId() *string
	SetDeviceId(v string) *GetDataChannelCredentialRequest
	GetDeviceId() *string
}

type GetDataChannelCredentialRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// device-3i5x4234f2j4w55e
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s GetDataChannelCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataChannelCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetDataChannelCredentialRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *GetDataChannelCredentialRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetDataChannelCredentialRequest) SetBusinessUnitId(v string) *GetDataChannelCredentialRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *GetDataChannelCredentialRequest) SetDeviceId(v string) *GetDataChannelCredentialRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDataChannelCredentialRequest) Validate() error {
	return dara.Validate(s)
}
