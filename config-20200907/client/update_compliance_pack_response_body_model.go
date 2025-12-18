// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCompliancePackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *UpdateCompliancePackResponseBody
	GetCompliancePackId() *string
	SetRequestId(v string) *UpdateCompliancePackResponseBody
	GetRequestId() *string
}

type UpdateCompliancePackResponseBody struct {
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-a8a8626622af0082****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCompliancePackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackResponseBody) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *UpdateCompliancePackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCompliancePackResponseBody) SetCompliancePackId(v string) *UpdateCompliancePackResponseBody {
	s.CompliancePackId = &v
	return s
}

func (s *UpdateCompliancePackResponseBody) SetRequestId(v string) *UpdateCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCompliancePackResponseBody) Validate() error {
	return dara.Validate(s)
}
