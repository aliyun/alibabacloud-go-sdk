// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTequestId(v string) *UntagResourcesResponseBody
	GetTequestId() *string
}

type UntagResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1-A-0-B
	TequestId *string `json:"tequestId,omitempty" xml:"tequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) GetTequestId() *string {
	return s.TequestId
}

func (s *UntagResourcesResponseBody) SetTequestId(v string) *UntagResourcesResponseBody {
	s.TequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
