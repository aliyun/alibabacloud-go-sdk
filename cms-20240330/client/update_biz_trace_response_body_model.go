// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizTraceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizTraceId(v string) *UpdateBizTraceResponseBody
	GetBizTraceId() *string
	SetRequestId(v string) *UpdateBizTraceResponseBody
	GetRequestId() *string
}

type UpdateBizTraceResponseBody struct {
	// example:
	//
	// e339260ed64c95d
	BizTraceId *string `json:"bizTraceId,omitempty" xml:"bizTraceId,omitempty"`
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateBizTraceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizTraceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBizTraceResponseBody) GetBizTraceId() *string {
	return s.BizTraceId
}

func (s *UpdateBizTraceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBizTraceResponseBody) SetBizTraceId(v string) *UpdateBizTraceResponseBody {
	s.BizTraceId = &v
	return s
}

func (s *UpdateBizTraceResponseBody) SetRequestId(v string) *UpdateBizTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBizTraceResponseBody) Validate() error {
	return dara.Validate(s)
}
