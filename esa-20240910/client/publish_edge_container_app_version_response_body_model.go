// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishEdgeContainerAppVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishEdgeContainerAppVersionResponseBody
	GetRequestId() *string
}

type PublishEdgeContainerAppVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishEdgeContainerAppVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishEdgeContainerAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishEdgeContainerAppVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishEdgeContainerAppVersionResponseBody) SetRequestId(v string) *PublishEdgeContainerAppVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishEdgeContainerAppVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
