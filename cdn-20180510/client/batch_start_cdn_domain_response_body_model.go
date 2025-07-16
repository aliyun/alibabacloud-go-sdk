// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartCdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchStartCdnDomainResponseBody
	GetRequestId() *string
}

type BatchStartCdnDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStartCdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStartCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartCdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStartCdnDomainResponseBody) SetRequestId(v string) *BatchStartCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStartCdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
