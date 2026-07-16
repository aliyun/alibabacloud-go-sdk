// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVersionResponseBody
	GetRequestId() *string
}

type DeleteVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVersionResponseBody) SetRequestId(v string) *DeleteVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
