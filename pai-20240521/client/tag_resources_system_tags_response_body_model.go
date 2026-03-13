// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesSystemTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TagResourcesSystemTagsResponseBody
	GetRequestId() *string
}

type TagResourcesSystemTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesSystemTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesSystemTagsResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesSystemTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagResourcesSystemTagsResponseBody) SetRequestId(v string) *TagResourcesSystemTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesSystemTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
