// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeprecateFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeprecateFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeprecateFlowResponseBody
	GetCode() *string
	SetMessage(v string) *DeprecateFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeprecateFlowResponseBody
	GetRequestId() *string
}

type DeprecateFlowResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The result returns OK as normal.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error description information.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeprecateFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeprecateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeprecateFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeprecateFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeprecateFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeprecateFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeprecateFlowResponseBody) SetAccessDeniedDetail(v string) *DeprecateFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeprecateFlowResponseBody) SetCode(v string) *DeprecateFlowResponseBody {
	s.Code = &v
	return s
}

func (s *DeprecateFlowResponseBody) SetMessage(v string) *DeprecateFlowResponseBody {
	s.Message = &v
	return s
}

func (s *DeprecateFlowResponseBody) SetRequestId(v string) *DeprecateFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeprecateFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
