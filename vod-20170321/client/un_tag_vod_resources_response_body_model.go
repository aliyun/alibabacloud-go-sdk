// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagVodResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnTagVodResourcesResponseBody
	GetRequestId() *string
}

type UnTagVodResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnTagVodResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnTagVodResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagVodResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnTagVodResourcesResponseBody) SetRequestId(v string) *UnTagVodResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnTagVodResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
