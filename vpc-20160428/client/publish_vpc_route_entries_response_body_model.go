// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishVpcRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishVpcRouteEntriesResponseBody
	GetRequestId() *string
}

type PublishVpcRouteEntriesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 980960B0-2969-40BF-8542-EBB34FD358AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishVpcRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishVpcRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *PublishVpcRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishVpcRouteEntriesResponseBody) SetRequestId(v string) *PublishVpcRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishVpcRouteEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
