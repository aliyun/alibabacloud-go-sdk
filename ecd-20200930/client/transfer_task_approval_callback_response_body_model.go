// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTaskApprovalCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransferTaskApprovalCallbackResponseBody
	GetRequestId() *string
}

type TransferTaskApprovalCallbackResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AE7B699F-625C-587E-BC5F-1395CA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferTaskApprovalCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferTaskApprovalCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *TransferTaskApprovalCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferTaskApprovalCallbackResponseBody) SetRequestId(v string) *TransferTaskApprovalCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferTaskApprovalCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
