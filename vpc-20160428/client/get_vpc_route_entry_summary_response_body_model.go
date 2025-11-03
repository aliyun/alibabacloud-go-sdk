// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcRouteEntrySummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetVpcRouteEntrySummaryResponseBody
	GetRequestId() *string
	SetRouteEntrySummarys(v []*GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys) *GetVpcRouteEntrySummaryResponseBody
	GetRouteEntrySummarys() []*GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys
}

type GetVpcRouteEntrySummaryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DC668356-BCB4-42FD-9BC3-FA2B2E04B634
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the routes in the route tables.
	RouteEntrySummarys []*GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys `json:"RouteEntrySummarys,omitempty" xml:"RouteEntrySummarys,omitempty" type:"Repeated"`
}

func (s GetVpcRouteEntrySummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpcRouteEntrySummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcRouteEntrySummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVpcRouteEntrySummaryResponseBody) GetRouteEntrySummarys() []*GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys {
	return s.RouteEntrySummarys
}

func (s *GetVpcRouteEntrySummaryResponseBody) SetRequestId(v string) *GetVpcRouteEntrySummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpcRouteEntrySummaryResponseBody) SetRouteEntrySummarys(v []*GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys) *GetVpcRouteEntrySummaryResponseBody {
	s.RouteEntrySummarys = v
	return s
}

func (s *GetVpcRouteEntrySummaryResponseBody) Validate() error {
	if s.RouteEntrySummarys != nil {
		for _, item := range s.RouteEntrySummarys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys struct {
	// The information about the routes of different types in one route table.
	EntrySummarys []*GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys `json:"EntrySummarys,omitempty" xml:"EntrySummarys,omitempty" type:"Repeated"`
	// The ID of the route table.
	//
	// example:
	//
	// vtb-bp145q7glnuzdvzu2****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys) String() string {
	return dara.Prettify(s)
}

func (s GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys) GoString() string {
	return s.String()
}

func (s *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys) GetEntrySummarys() []*GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys {
	return s.EntrySummarys
}

func (s *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys) SetEntrySummarys(v []*GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys) *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys {
	s.EntrySummarys = v
	return s
}

func (s *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys) SetRouteTableId(v string) *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys {
	s.RouteTableId = &v
	return s
}

func (s *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarys) Validate() error {
	if s.EntrySummarys != nil {
		for _, item := range s.EntrySummarys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys struct {
	// The number of entries returned.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The type of the route. Valid values:
	//
	// 	- **All**: all route types
	//
	// 	- **Custom**: a custom route
	//
	// 	- **System**: a system route
	//
	// 	- **BGP**: a BGP route
	//
	// 	- **CEN**: a CEN route
	//
	// example:
	//
	// Custom
	RouteEntryType *string `json:"RouteEntryType,omitempty" xml:"RouteEntryType,omitempty"`
}

func (s GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys) String() string {
	return dara.Prettify(s)
}

func (s GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys) GoString() string {
	return s.String()
}

func (s *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys) GetCount() *int32 {
	return s.Count
}

func (s *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys) GetRouteEntryType() *string {
	return s.RouteEntryType
}

func (s *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys) SetCount(v int32) *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys {
	s.Count = &v
	return s
}

func (s *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys) SetRouteEntryType(v string) *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys {
	s.RouteEntryType = &v
	return s
}

func (s *GetVpcRouteEntrySummaryResponseBodyRouteEntrySummarysEntrySummarys) Validate() error {
	return dara.Validate(s)
}
