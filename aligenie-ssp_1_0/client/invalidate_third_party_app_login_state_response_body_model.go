// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidateThirdPartyAppLoginStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InvalidateThirdPartyAppLoginStateResponseBody
	GetCode() *int32
	SetMessage(v string) *InvalidateThirdPartyAppLoginStateResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvalidateThirdPartyAppLoginStateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvalidateThirdPartyAppLoginStateResponseBody
	GetSuccess() *bool
}

type InvalidateThirdPartyAppLoginStateResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FAFCD152-4791-5F2F-B0BE-2DC06FD4F05B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InvalidateThirdPartyAppLoginStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvalidateThirdPartyAppLoginStateResponseBody) GoString() string {
	return s.String()
}

func (s *InvalidateThirdPartyAppLoginStateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InvalidateThirdPartyAppLoginStateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvalidateThirdPartyAppLoginStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvalidateThirdPartyAppLoginStateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvalidateThirdPartyAppLoginStateResponseBody) SetCode(v int32) *InvalidateThirdPartyAppLoginStateResponseBody {
	s.Code = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateResponseBody) SetMessage(v string) *InvalidateThirdPartyAppLoginStateResponseBody {
	s.Message = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateResponseBody) SetRequestId(v string) *InvalidateThirdPartyAppLoginStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateResponseBody) SetSuccess(v bool) *InvalidateThirdPartyAppLoginStateResponseBody {
	s.Success = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateResponseBody) Validate() error {
	return dara.Validate(s)
}
