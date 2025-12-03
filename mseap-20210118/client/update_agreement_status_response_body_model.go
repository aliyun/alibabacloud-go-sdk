// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgreementStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAgreementStatusResponseBody
	GetRequestId() *string
}

type UpdateAgreementStatusResponseBody struct {
	// example:
	//
	// 6BDB1964-A6E9-5946-89A4-523D35645BB6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAgreementStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgreementStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgreementStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgreementStatusResponseBody) SetRequestId(v string) *UpdateAgreementStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgreementStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
