// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopDcdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchStopDcdnDomainResponseBody
	GetRequestId() *string
}

type BatchStopDcdnDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStopDcdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStopDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopDcdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStopDcdnDomainResponseBody) SetRequestId(v string) *BatchStopDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStopDcdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
