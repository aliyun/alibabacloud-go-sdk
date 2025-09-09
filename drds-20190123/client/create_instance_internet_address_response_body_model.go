// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceInternetAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateInstanceInternetAddressResponseBody
	GetCode() *int32
	SetData(v bool) *CreateInstanceInternetAddressResponseBody
	GetData() *bool
	SetRequestId(v string) *CreateInstanceInternetAddressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceInternetAddressResponseBody
	GetSuccess() *bool
}

type CreateInstanceInternetAddressResponseBody struct {
	// The error code returned when the activity fails.
	//
	// >  This parameter appears only when an error occurs during the API call.
	//
	// example:
	//
	// 404
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the public IP address was created.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1DF6052F-15E2-4E69-9628-D6BCC3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceInternetAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInternetAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceInternetAddressResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateInstanceInternetAddressResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateInstanceInternetAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceInternetAddressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceInternetAddressResponseBody) SetCode(v int32) *CreateInstanceInternetAddressResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceInternetAddressResponseBody) SetData(v bool) *CreateInstanceInternetAddressResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceInternetAddressResponseBody) SetRequestId(v string) *CreateInstanceInternetAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceInternetAddressResponseBody) SetSuccess(v bool) *CreateInstanceInternetAddressResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceInternetAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
