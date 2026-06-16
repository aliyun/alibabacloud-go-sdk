// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CommonResult
	GetCode() *string
	SetData(v interface{}) *CommonResult
	GetData() interface{}
	SetRequestId(v string) *CommonResult
	GetRequestId() *string
}

type CommonResult struct {
	// Indicates the request status. A value of `SUCCESS` indicates success. On failure, an error code is returned, such as `ERR_BAD_REQUEST`, `ERR_VALIDATION_FAILED`, or `ERR_INTERNAL_SERVER_ERROR`.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned business data.
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// A unique request identifier for issue tracking.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CommonResult) String() string {
	return dara.Prettify(s)
}

func (s CommonResult) GoString() string {
	return s.String()
}

func (s *CommonResult) GetCode() *string {
	return s.Code
}

func (s *CommonResult) GetData() interface{} {
	return s.Data
}

func (s *CommonResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CommonResult) SetCode(v string) *CommonResult {
	s.Code = &v
	return s
}

func (s *CommonResult) SetData(v interface{}) *CommonResult {
	s.Data = v
	return s
}

func (s *CommonResult) SetRequestId(v string) *CommonResult {
	s.RequestId = &v
	return s
}

func (s *CommonResult) Validate() error {
	return dara.Validate(s)
}
