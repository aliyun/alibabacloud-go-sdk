// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBResourcePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBResourcePoolResponseBody
	GetRequestId() *string
}

type DeleteDBResourcePoolResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBResourcePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBResourcePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBResourcePoolResponseBody) SetRequestId(v string) *DeleteDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBResourcePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
