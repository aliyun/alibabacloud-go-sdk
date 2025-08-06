// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAuditResponseBody
	GetRequestId() *string
}

type CreateAuditResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAuditResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAuditResponseBody) SetRequestId(v string) *CreateAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAuditResponseBody) Validate() error {
	return dara.Validate(s)
}
