// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v string) *DeleteDomainRecordResponseBody
	GetRecordId() *string
	SetRequestId(v string) *DeleteDomainRecordResponseBody
	GetRequestId() *string
}

type DeleteDomainRecordResponseBody struct {
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

func (s DeleteDomainRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainRecordResponseBody) GetRecordId() *string {
	return s.RecordId
}

func (s *DeleteDomainRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDomainRecordResponseBody) SetRecordId(v string) *DeleteDomainRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *DeleteDomainRecordResponseBody) SetRequestId(v string) *DeleteDomainRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
