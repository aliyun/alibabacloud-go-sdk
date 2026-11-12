// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomImagesResponseBody
	GetRequestId() *string
}

type DeleteCustomImagesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomImagesResponseBody) SetRequestId(v string) *DeleteCustomImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomImagesResponseBody) Validate() error {
	return dara.Validate(s)
}
