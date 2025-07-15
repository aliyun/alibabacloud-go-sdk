// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishVpnRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishVpnRouteEntryResponseBody
	GetRequestId() *string
}

type PublishVpnRouteEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishVpnRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishVpnRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *PublishVpnRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishVpnRouteEntryResponseBody) SetRequestId(v string) *PublishVpnRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishVpnRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
