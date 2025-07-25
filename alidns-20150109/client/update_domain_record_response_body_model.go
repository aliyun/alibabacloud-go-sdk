// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v string) *UpdateDomainRecordResponseBody
	GetRecordId() *string
	SetRequestId(v string) *UpdateDomainRecordResponseBody
	GetRequestId() *string
}

type UpdateDomainRecordResponseBody struct {
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

func (s UpdateDomainRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainRecordResponseBody) GetRecordId() *string {
	return s.RecordId
}

func (s *UpdateDomainRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDomainRecordResponseBody) SetRecordId(v string) *UpdateDomainRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *UpdateDomainRecordResponseBody) SetRequestId(v string) *UpdateDomainRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
