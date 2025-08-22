// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartDcdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchStartDcdnDomainResponseBody
	GetRequestId() *string
}

type BatchStartDcdnDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStartDcdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStartDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartDcdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStartDcdnDomainResponseBody) SetRequestId(v string) *BatchStartDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStartDcdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
