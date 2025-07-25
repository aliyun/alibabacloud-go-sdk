// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainRecordStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v string) *SetDomainRecordStatusResponseBody
	GetRecordId() *string
	SetRequestId(v string) *SetDomainRecordStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *SetDomainRecordStatusResponseBody
	GetStatus() *string
}

type SetDomainRecordStatusResponseBody struct {
	// The ID of the DNS record.
	//
	// example:
	//
	// 9999985
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the DNS record.
	//
	// example:
	//
	// Disable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDomainRecordStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDomainRecordStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainRecordStatusResponseBody) GetRecordId() *string {
	return s.RecordId
}

func (s *SetDomainRecordStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDomainRecordStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SetDomainRecordStatusResponseBody) SetRecordId(v string) *SetDomainRecordStatusResponseBody {
	s.RecordId = &v
	return s
}

func (s *SetDomainRecordStatusResponseBody) SetRequestId(v string) *SetDomainRecordStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDomainRecordStatusResponseBody) SetStatus(v string) *SetDomainRecordStatusResponseBody {
	s.Status = &v
	return s
}

func (s *SetDomainRecordStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
