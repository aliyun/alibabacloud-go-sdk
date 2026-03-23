// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTagsFromResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveTagsFromResourceResponseBody
	GetRequestId() *string
}

type RemoveTagsFromResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveTagsFromResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsFromResourceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTagsFromResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveTagsFromResourceResponseBody) SetRequestId(v string) *RemoveTagsFromResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTagsFromResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
