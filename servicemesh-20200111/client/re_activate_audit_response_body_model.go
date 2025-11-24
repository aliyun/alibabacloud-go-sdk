// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReActivateAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ReActivateAuditResponseBody
	GetData() *string
	SetRequestId(v string) *ReActivateAuditResponseBody
	GetRequestId() *string
}

type ReActivateAuditResponseBody struct {
	// The name of the project that is used to store audit logs.
	//
	// example:
	//
	// k8s-log-c0703599f695f4b8fa1c6492a33af****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17163CE9-CE4B-1B87-9185-1A1AD7E7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReActivateAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReActivateAuditResponseBody) GoString() string {
	return s.String()
}

func (s *ReActivateAuditResponseBody) GetData() *string {
	return s.Data
}

func (s *ReActivateAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReActivateAuditResponseBody) SetData(v string) *ReActivateAuditResponseBody {
	s.Data = &v
	return s
}

func (s *ReActivateAuditResponseBody) SetRequestId(v string) *ReActivateAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReActivateAuditResponseBody) Validate() error {
	return dara.Validate(s)
}
