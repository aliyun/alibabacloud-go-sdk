// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagVodResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TagVodResourcesResponseBody
	GetRequestId() *string
}

type TagVodResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagVodResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagVodResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagVodResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagVodResourcesResponseBody) SetRequestId(v string) *TagVodResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagVodResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
