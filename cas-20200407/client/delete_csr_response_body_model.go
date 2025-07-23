// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCsrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCsrResponseBody
	GetRequestId() *string
}

type DeleteCsrResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D3F1FA43-1C26-50A2-8F0F-7A03851DBB46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCsrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCsrResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCsrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCsrResponseBody) SetRequestId(v string) *DeleteCsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCsrResponseBody) Validate() error {
	return dara.Validate(s)
}
