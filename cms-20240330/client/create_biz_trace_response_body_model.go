// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizTraceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizTraceId(v string) *CreateBizTraceResponseBody
	GetBizTraceId() *string
	SetRequestId(v string) *CreateBizTraceResponseBody
	GetRequestId() *string
}

type CreateBizTraceResponseBody struct {
	// example:
	//
	// e339260ed64c95d
	BizTraceId *string `json:"bizTraceId,omitempty" xml:"bizTraceId,omitempty"`
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateBizTraceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBizTraceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBizTraceResponseBody) GetBizTraceId() *string {
	return s.BizTraceId
}

func (s *CreateBizTraceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBizTraceResponseBody) SetBizTraceId(v string) *CreateBizTraceResponseBody {
	s.BizTraceId = &v
	return s
}

func (s *CreateBizTraceResponseBody) SetRequestId(v string) *CreateBizTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBizTraceResponseBody) Validate() error {
	return dara.Validate(s)
}
