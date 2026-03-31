// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateCompliancePackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *UpdateAggregateCompliancePackResponseBody
	GetCompliancePackId() *string
	SetRequestId(v string) *UpdateAggregateCompliancePackResponseBody
	GetRequestId() *string
}

type UpdateAggregateCompliancePackResponseBody struct {
	// The compliance package ID.
	//
	// example:
	//
	// ca-f632626622af0079****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAggregateCompliancePackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackResponseBody) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *UpdateAggregateCompliancePackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAggregateCompliancePackResponseBody) SetCompliancePackId(v string) *UpdateAggregateCompliancePackResponseBody {
	s.CompliancePackId = &v
	return s
}

func (s *UpdateAggregateCompliancePackResponseBody) SetRequestId(v string) *UpdateAggregateCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAggregateCompliancePackResponseBody) Validate() error {
	return dara.Validate(s)
}
