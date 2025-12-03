// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveTagsResponseBody
	GetRequestId() *string
}

type RemoveTagsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveTagsResponseBody) SetRequestId(v string) *RemoveTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
