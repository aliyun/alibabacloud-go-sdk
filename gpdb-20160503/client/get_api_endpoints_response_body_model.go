// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*GetApiEndpointsResponseBodyItems) *GetApiEndpointsResponseBody
	GetItems() []*GetApiEndpointsResponseBodyItems
	SetMaxResults(v int32) *GetApiEndpointsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *GetApiEndpointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetApiEndpointsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *GetApiEndpointsResponseBody
	GetTotalRecordCount() *int32
}

type GetApiEndpointsResponseBody struct {
	// The list of parameters.
	Items []*GetApiEndpointsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of records to return in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query during paging. Use this token to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s GetApiEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiEndpointsResponseBody) GetItems() []*GetApiEndpointsResponseBodyItems {
	return s.Items
}

func (s *GetApiEndpointsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetApiEndpointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetApiEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiEndpointsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *GetApiEndpointsResponseBody) SetItems(v []*GetApiEndpointsResponseBodyItems) *GetApiEndpointsResponseBody {
	s.Items = v
	return s
}

func (s *GetApiEndpointsResponseBody) SetMaxResults(v int32) *GetApiEndpointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetApiEndpointsResponseBody) SetNextToken(v string) *GetApiEndpointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetApiEndpointsResponseBody) SetRequestId(v string) *GetApiEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApiEndpointsResponseBody) SetTotalRecordCount(v int32) *GetApiEndpointsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *GetApiEndpointsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApiEndpointsResponseBodyItems struct {
	// The endpoint.
	//
	// example:
	//
	// https://api-longmemory-cn-beijing.opentrust.net/
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service type. Valid values:
	//
	// - **memory**
	//
	// - **drama**
	//
	// example:
	//
	// memory
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetApiEndpointsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetApiEndpointsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetApiEndpointsResponseBodyItems) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetApiEndpointsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *GetApiEndpointsResponseBodyItems) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetApiEndpointsResponseBodyItems) SetEndpoint(v string) *GetApiEndpointsResponseBodyItems {
	s.Endpoint = &v
	return s
}

func (s *GetApiEndpointsResponseBodyItems) SetRegionId(v string) *GetApiEndpointsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *GetApiEndpointsResponseBodyItems) SetServiceType(v string) *GetApiEndpointsResponseBodyItems {
	s.ServiceType = &v
	return s
}

func (s *GetApiEndpointsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
