// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v string) *AddDomainRecordResponseBody
	GetRecordId() *string
	SetRequestId(v string) *AddDomainRecordResponseBody
	GetRequestId() *string
}

type AddDomainRecordResponseBody struct {
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
}

func (s AddDomainRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDomainRecordResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainRecordResponseBody) GetRecordId() *string {
	return s.RecordId
}

func (s *AddDomainRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDomainRecordResponseBody) SetRecordId(v string) *AddDomainRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *AddDomainRecordResponseBody) SetRequestId(v string) *AddDomainRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDomainRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
