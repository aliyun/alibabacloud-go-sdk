// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddTagsResponseBody
	GetRequestId() *string
}

type AddTagsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTagsResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTagsResponseBody) SetRequestId(v string) *AddTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
