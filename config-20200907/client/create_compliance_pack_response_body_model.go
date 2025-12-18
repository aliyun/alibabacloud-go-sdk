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
	// The compliance package ID.
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
