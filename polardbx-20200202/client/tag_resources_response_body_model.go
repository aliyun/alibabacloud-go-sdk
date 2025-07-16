// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TagResourcesResponseBody
	GetRequestId() *string
}

type TagResourcesResponseBody struct {
	// requestId
	//
	// example:
	//
	// xxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
