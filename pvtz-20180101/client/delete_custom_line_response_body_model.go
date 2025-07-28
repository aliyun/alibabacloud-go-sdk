// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLineId(v string) *DeleteCustomLineResponseBody
	GetLineId() *string
	SetRequestId(v string) *DeleteCustomLineResponseBody
	GetRequestId() *string
}

type DeleteCustomLineResponseBody struct {
	// The unique ID of the custom line.
	//
	// example:
	//
	// 520002
	LineId *string `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A73F3BD0-B1A8-42A9-A9B6-689BBABC4891
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomLineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomLineResponseBody) GetLineId() *string {
	return s.LineId
}

func (s *DeleteCustomLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomLineResponseBody) SetLineId(v string) *DeleteCustomLineResponseBody {
	s.LineId = &v
	return s
}

func (s *DeleteCustomLineResponseBody) SetRequestId(v string) *DeleteCustomLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomLineResponseBody) Validate() error {
	return dara.Validate(s)
}
