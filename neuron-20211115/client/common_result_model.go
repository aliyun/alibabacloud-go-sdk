// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CommonResult
	GetRequestId() *string
	SetSuccess(v bool) *CommonResult
	GetSuccess() *bool
}

type CommonResult struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CommonResult) String() string {
	return dara.Prettify(s)
}

func (s CommonResult) GoString() string {
	return s.String()
}

func (s *CommonResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CommonResult) GetSuccess() *bool {
	return s.Success
}

func (s *CommonResult) SetRequestId(v string) *CommonResult {
	s.RequestId = &v
	return s
}

func (s *CommonResult) SetSuccess(v bool) *CommonResult {
	s.Success = &v
	return s
}

func (s *CommonResult) Validate() error {
	return dara.Validate(s)
}
