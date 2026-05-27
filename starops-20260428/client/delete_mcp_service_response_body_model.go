// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMcpServiceResponseBody
	GetRequestId() *string
}

type DeleteMcpServiceResponseBody struct {
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMcpServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcpServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcpServiceResponseBody) SetRequestId(v string) *DeleteMcpServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcpServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
