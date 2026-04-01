// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateReadWriteSplittingConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AllocateReadWriteSplittingConnectionResponseBody
	GetRequestId() *string
}

type AllocateReadWriteSplittingConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateReadWriteSplittingConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateReadWriteSplittingConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateReadWriteSplittingConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateReadWriteSplittingConnectionResponseBody) SetRequestId(v string) *AllocateReadWriteSplittingConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
