// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizTraceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBizTraceResponseBody
	GetRequestId() *string
}

type DeleteBizTraceResponseBody struct {
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteBizTraceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizTraceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBizTraceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBizTraceResponseBody) SetRequestId(v string) *DeleteBizTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBizTraceResponseBody) Validate() error {
	return dara.Validate(s)
}
