// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregatedDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregations(v *GetAggregatedDesktopsResponseBodyAggregations) *GetAggregatedDesktopsResponseBody
	GetAggregations() *GetAggregatedDesktopsResponseBodyAggregations
	SetRequestId(v string) *GetAggregatedDesktopsResponseBody
	GetRequestId() *string
}

type GetAggregatedDesktopsResponseBody struct {
	// The list of aggregation field information.
	//
	// 	Notice: When you use an aggregate query, only aggregation results are returned. The list of matched metadata is not returned.
	Aggregations *GetAggregatedDesktopsResponseBodyAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 3147E094-C1F7-5001-8F04-C8CEE75D6552
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregatedDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregatedDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregatedDesktopsResponseBody) GetAggregations() *GetAggregatedDesktopsResponseBodyAggregations {
	return s.Aggregations
}

func (s *GetAggregatedDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregatedDesktopsResponseBody) SetAggregations(v *GetAggregatedDesktopsResponseBodyAggregations) *GetAggregatedDesktopsResponseBody {
	s.Aggregations = v
	return s
}

func (s *GetAggregatedDesktopsResponseBody) SetRequestId(v string) *GetAggregatedDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregatedDesktopsResponseBody) Validate() error {
	if s.Aggregations != nil {
		if err := s.Aggregations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregatedDesktopsResponseBodyAggregations struct {
	// The aggregation results.
	DesktopAggregation []map[string]*string `json:"DesktopAggregation,omitempty" xml:"DesktopAggregation,omitempty" type:"Repeated"`
}

func (s GetAggregatedDesktopsResponseBodyAggregations) String() string {
	return dara.Prettify(s)
}

func (s GetAggregatedDesktopsResponseBodyAggregations) GoString() string {
	return s.String()
}

func (s *GetAggregatedDesktopsResponseBodyAggregations) GetDesktopAggregation() []map[string]*string {
	return s.DesktopAggregation
}

func (s *GetAggregatedDesktopsResponseBodyAggregations) SetDesktopAggregation(v []map[string]*string) *GetAggregatedDesktopsResponseBodyAggregations {
	s.DesktopAggregation = v
	return s
}

func (s *GetAggregatedDesktopsResponseBodyAggregations) Validate() error {
	return dara.Validate(s)
}
