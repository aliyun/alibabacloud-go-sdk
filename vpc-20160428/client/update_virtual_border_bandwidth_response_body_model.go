// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirtualBorderBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateVirtualBorderBandwidthResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateVirtualBorderBandwidthResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateVirtualBorderBandwidthResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateVirtualBorderBandwidthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateVirtualBorderBandwidthResponseBody
	GetSuccess() *bool
}

type UpdateVirtualBorderBandwidthResponseBody struct {
	// The error code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// none
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AA4486A8-B6AE-469E-AB09-820EF8ECFA2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateVirtualBorderBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirtualBorderBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVirtualBorderBandwidthResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVirtualBorderBandwidthResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateVirtualBorderBandwidthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVirtualBorderBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVirtualBorderBandwidthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateVirtualBorderBandwidthResponseBody) SetCode(v string) *UpdateVirtualBorderBandwidthResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthResponseBody) SetHttpStatusCode(v int32) *UpdateVirtualBorderBandwidthResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthResponseBody) SetMessage(v string) *UpdateVirtualBorderBandwidthResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthResponseBody) SetRequestId(v string) *UpdateVirtualBorderBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthResponseBody) SetSuccess(v bool) *UpdateVirtualBorderBandwidthResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
