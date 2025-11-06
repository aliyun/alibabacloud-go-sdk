// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOperationAuditInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *SubmitOperationAuditInfoResponseBody
	GetId() *int64
	SetRequestId(v string) *SubmitOperationAuditInfoResponseBody
	GetRequestId() *string
}

type SubmitOperationAuditInfoResponseBody struct {
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 9DKCF6F8-243C-40EC-8035-4B12FEFD7C22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitOperationAuditInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitOperationAuditInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitOperationAuditInfoResponseBody) GetId() *int64 {
	return s.Id
}

func (s *SubmitOperationAuditInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitOperationAuditInfoResponseBody) SetId(v int64) *SubmitOperationAuditInfoResponseBody {
	s.Id = &v
	return s
}

func (s *SubmitOperationAuditInfoResponseBody) SetRequestId(v string) *SubmitOperationAuditInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitOperationAuditInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
