// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBgpGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBgpGroupResponseBody
	GetRequestId() *string
}

type DeleteBgpGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBgpGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBgpGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBgpGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBgpGroupResponseBody) SetRequestId(v string) *DeleteBgpGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBgpGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
