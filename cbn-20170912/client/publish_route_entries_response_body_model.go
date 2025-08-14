// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishRouteEntriesResponseBody
	GetRequestId() *string
}

type PublishRouteEntriesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FBDB18D8-E91E-4978-8D6C-6E2E3EE10133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *PublishRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishRouteEntriesResponseBody) SetRequestId(v string) *PublishRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishRouteEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
