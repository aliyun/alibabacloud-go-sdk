// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipFullJobTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SkipFullJobTableResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SkipFullJobTableResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SkipFullJobTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SkipFullJobTableResponseBody
	GetSuccess() *bool
}

type SkipFullJobTableResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4D0ADAD5-DD97-41B6-B78F-D1961AB1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SkipFullJobTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SkipFullJobTableResponseBody) GoString() string {
	return s.String()
}

func (s *SkipFullJobTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *SkipFullJobTableResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SkipFullJobTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SkipFullJobTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SkipFullJobTableResponseBody) SetCode(v string) *SkipFullJobTableResponseBody {
	s.Code = &v
	return s
}

func (s *SkipFullJobTableResponseBody) SetHttpStatusCode(v int32) *SkipFullJobTableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SkipFullJobTableResponseBody) SetRequestId(v string) *SkipFullJobTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *SkipFullJobTableResponseBody) SetSuccess(v bool) *SkipFullJobTableResponseBody {
	s.Success = &v
	return s
}

func (s *SkipFullJobTableResponseBody) Validate() error {
	return dara.Validate(s)
}
