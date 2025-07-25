// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainRecordRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDomainRecordRemarkResponseBody
	GetRequestId() *string
}

type UpdateDomainRecordRemarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDomainRecordRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainRecordRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainRecordRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDomainRecordRemarkResponseBody) SetRequestId(v string) *UpdateDomainRecordRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainRecordRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
