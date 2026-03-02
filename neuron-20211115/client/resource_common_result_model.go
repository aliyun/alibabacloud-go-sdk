// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceCommonResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResourceCommonResult
	GetRequestId() *string
	SetSuccess(v bool) *ResourceCommonResult
	GetSuccess() *bool
}

type ResourceCommonResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ResourceCommonResult) String() string {
	return dara.Prettify(s)
}

func (s ResourceCommonResult) GoString() string {
	return s.String()
}

func (s *ResourceCommonResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ResourceCommonResult) GetSuccess() *bool {
	return s.Success
}

func (s *ResourceCommonResult) SetRequestId(v string) *ResourceCommonResult {
	s.RequestId = &v
	return s
}

func (s *ResourceCommonResult) SetSuccess(v bool) *ResourceCommonResult {
	s.Success = &v
	return s
}

func (s *ResourceCommonResult) Validate() error {
	return dara.Validate(s)
}
