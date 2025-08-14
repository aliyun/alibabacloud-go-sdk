// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdInsertionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAdInsertionResponseBody
	GetRequestId() *string
}

type DeleteAdInsertionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAdInsertionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdInsertionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAdInsertionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAdInsertionResponseBody) SetRequestId(v string) *DeleteAdInsertionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAdInsertionResponseBody) Validate() error {
	return dara.Validate(s)
}
