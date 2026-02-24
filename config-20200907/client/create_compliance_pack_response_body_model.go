// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompliancePackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *CreateCompliancePackResponseBody
	GetCompliancePackId() *string
	SetRequestId(v string) *CreateCompliancePackResponseBody
	GetRequestId() *string
}

type CreateCompliancePackResponseBody struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCompliancePackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackResponseBody) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *CreateCompliancePackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCompliancePackResponseBody) SetCompliancePackId(v string) *CreateCompliancePackResponseBody {
	s.CompliancePackId = &v
	return s
}

func (s *CreateCompliancePackResponseBody) SetRequestId(v string) *CreateCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCompliancePackResponseBody) Validate() error {
	return dara.Validate(s)
}
