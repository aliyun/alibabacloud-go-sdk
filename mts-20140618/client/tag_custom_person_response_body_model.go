// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagCustomPersonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TagCustomPersonResponseBody
	GetRequestId() *string
}

type TagCustomPersonResponseBody struct {
	// example:
	//
	// FD8B5B8C-0C3D-5776-B3B1-EE6AD11F905A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagCustomPersonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagCustomPersonResponseBody) GoString() string {
	return s.String()
}

func (s *TagCustomPersonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagCustomPersonResponseBody) SetRequestId(v string) *TagCustomPersonResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagCustomPersonResponseBody) Validate() error {
	return dara.Validate(s)
}
