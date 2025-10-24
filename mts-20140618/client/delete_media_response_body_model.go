// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMediaResponseBody
	GetRequestId() *string
}

type DeleteMediaResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 05F8B913-E9F3-4A6F-9922-48CADA0FFAAD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaResponseBody) SetRequestId(v string) *DeleteMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaResponseBody) Validate() error {
	return dara.Validate(s)
}
