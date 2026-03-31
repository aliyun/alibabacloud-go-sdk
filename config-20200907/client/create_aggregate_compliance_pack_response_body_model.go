// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateCompliancePackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *CreateAggregateCompliancePackResponseBody
	GetCompliancePackId() *string
	SetRequestId(v string) *CreateAggregateCompliancePackResponseBody
	GetRequestId() *string
}

type CreateAggregateCompliancePackResponseBody struct {
	// The compliance package ID.
	//
	// example:
	//
	// cp-fc56626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC0CE5EB-E51E-48EB-B4AB-9A9E131ECC0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAggregateCompliancePackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackResponseBody) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *CreateAggregateCompliancePackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAggregateCompliancePackResponseBody) SetCompliancePackId(v string) *CreateAggregateCompliancePackResponseBody {
	s.CompliancePackId = &v
	return s
}

func (s *CreateAggregateCompliancePackResponseBody) SetRequestId(v string) *CreateAggregateCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAggregateCompliancePackResponseBody) Validate() error {
	return dara.Validate(s)
}
