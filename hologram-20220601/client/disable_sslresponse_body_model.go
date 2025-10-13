// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DisableSSLResponseBody
	GetData() *bool
	SetErrorCode(v string) *DisableSSLResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DisableSSLResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *DisableSSLResponseBody
	GetHttpStatusCode() *string
	SetSuccess(v string) *DisableSSLResponseBody
	GetSuccess() *string
	SetRequestId(v string) *DisableSSLResponseBody
	GetRequestId() *string
}

type DisableSSLResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 819A7F0F-2951-540F-BD94-6A41ECF0281F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DisableSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSSLResponseBody) GetData() *bool {
	return s.Data
}

func (s *DisableSSLResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DisableSSLResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DisableSSLResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DisableSSLResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DisableSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSSLResponseBody) SetData(v bool) *DisableSSLResponseBody {
	s.Data = &v
	return s
}

func (s *DisableSSLResponseBody) SetErrorCode(v string) *DisableSSLResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DisableSSLResponseBody) SetErrorMessage(v string) *DisableSSLResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DisableSSLResponseBody) SetHttpStatusCode(v string) *DisableSSLResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableSSLResponseBody) SetSuccess(v string) *DisableSSLResponseBody {
	s.Success = &v
	return s
}

func (s *DisableSSLResponseBody) SetRequestId(v string) *DisableSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
