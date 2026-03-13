// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesSystemTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UntagResourcesSystemTagsResponseBody
	GetRequestId() *string
}

type UntagResourcesSystemTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesSystemTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesSystemTagsResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesSystemTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagResourcesSystemTagsResponseBody) SetRequestId(v string) *UntagResourcesSystemTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesSystemTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
