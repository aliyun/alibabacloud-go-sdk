// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOperationAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelOperationAuditResponseBody
	GetRequestId() *string
}

type CancelOperationAuditResponseBody struct {
	// example:
	//
	// 9KFCF6F8-243C-40EC-8035-4B12KKFD7D90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelOperationAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelOperationAuditResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOperationAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelOperationAuditResponseBody) SetRequestId(v string) *CancelOperationAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelOperationAuditResponseBody) Validate() error {
	return dara.Validate(s)
}
