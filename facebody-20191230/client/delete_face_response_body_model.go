// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFaceResponseBody
	GetRequestId() *string
}

type DeleteFaceResponseBody struct {
	// example:
	//
	// FAC90D32-2F04-5AD4-B94B-BAA163AB3724
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaceResponseBody) SetRequestId(v string) *DeleteFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceResponseBody) Validate() error {
	return dara.Validate(s)
}
