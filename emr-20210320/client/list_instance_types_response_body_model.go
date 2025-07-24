// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTypes(v []*InstanceType) *ListInstanceTypesResponseBody
	GetInstanceTypes() []*InstanceType
	SetMaxResults(v int32) *ListInstanceTypesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstanceTypesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListInstanceTypesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListInstanceTypesResponseBody
	GetTotalCount() *int32
}

type ListInstanceTypesResponseBody struct {
	// The instance types.
	InstanceTypes []*InstanceType `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The maximum number of records returned in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Returns the position of the read data.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total amount of data under the conditions of this request.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceTypesResponseBody) GetInstanceTypes() []*InstanceType {
	return s.InstanceTypes
}

func (s *ListInstanceTypesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceTypesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceTypesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstanceTypesResponseBody) SetInstanceTypes(v []*InstanceType) *ListInstanceTypesResponseBody {
	s.InstanceTypes = v
	return s
}

func (s *ListInstanceTypesResponseBody) SetMaxResults(v int32) *ListInstanceTypesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceTypesResponseBody) SetNextToken(v string) *ListInstanceTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstanceTypesResponseBody) SetRequestId(v string) *ListInstanceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceTypesResponseBody) SetTotalCount(v int32) *ListInstanceTypesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceTypesResponseBody) Validate() error {
	return dara.Validate(s)
}
