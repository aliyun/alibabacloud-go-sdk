// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountDeliveryAddressInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *QueryAccountDeliveryAddressInfoRequest
	GetAppName() *string
	SetPk(v string) *QueryAccountDeliveryAddressInfoRequest
	GetPk() *string
	SetRequestId(v string) *QueryAccountDeliveryAddressInfoRequest
	GetRequestId() *string
}

type QueryAccountDeliveryAddressInfoRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAccountDeliveryAddressInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountDeliveryAddressInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountDeliveryAddressInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *QueryAccountDeliveryAddressInfoRequest) GetPk() *string {
	return s.Pk
}

func (s *QueryAccountDeliveryAddressInfoRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountDeliveryAddressInfoRequest) SetAppName(v string) *QueryAccountDeliveryAddressInfoRequest {
	s.AppName = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoRequest) SetPk(v string) *QueryAccountDeliveryAddressInfoRequest {
	s.Pk = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoRequest) SetRequestId(v string) *QueryAccountDeliveryAddressInfoRequest {
	s.RequestId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoRequest) Validate() error {
	return dara.Validate(s)
}
