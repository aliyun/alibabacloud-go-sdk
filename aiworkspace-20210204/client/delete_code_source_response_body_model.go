// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCodeSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeSourceId(v string) *DeleteCodeSourceResponseBody
	GetCodeSourceId() *string
	SetRequestId(v string) *DeleteCodeSourceResponseBody
	GetRequestId() *string
}

type DeleteCodeSourceResponseBody struct {
	// The ID of the deleted code source.
	//
	// example:
	//
	// code-20210111103721-85qz78ia96lu
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCodeSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCodeSourceResponseBody) GetCodeSourceId() *string {
	return s.CodeSourceId
}

func (s *DeleteCodeSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCodeSourceResponseBody) SetCodeSourceId(v string) *DeleteCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *DeleteCodeSourceResponseBody) SetRequestId(v string) *DeleteCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCodeSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
