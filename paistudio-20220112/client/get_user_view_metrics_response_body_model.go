// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserViewMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *GetUserViewMetricsResponseBody
	GetResourceGroupId() *string
	SetSummary(v *UserViewMetric) *GetUserViewMetricsResponseBody
	GetSummary() *UserViewMetric
	SetTotal(v int32) *GetUserViewMetricsResponseBody
	GetTotal() *int32
	SetUserMetrics(v []*UserViewMetric) *GetUserViewMetricsResponseBody
	GetUserMetrics() []*UserViewMetric
}

type GetUserViewMetricsResponseBody struct {
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupId *string         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Summary         *UserViewMetric `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 2
	Total       *int32            `json:"Total,omitempty" xml:"Total,omitempty"`
	UserMetrics []*UserViewMetric `json:"UserMetrics,omitempty" xml:"UserMetrics,omitempty" type:"Repeated"`
}

func (s GetUserViewMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserViewMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserViewMetricsResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetUserViewMetricsResponseBody) GetSummary() *UserViewMetric {
	return s.Summary
}

func (s *GetUserViewMetricsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetUserViewMetricsResponseBody) GetUserMetrics() []*UserViewMetric {
	return s.UserMetrics
}

func (s *GetUserViewMetricsResponseBody) SetResourceGroupId(v string) *GetUserViewMetricsResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetUserViewMetricsResponseBody) SetSummary(v *UserViewMetric) *GetUserViewMetricsResponseBody {
	s.Summary = v
	return s
}

func (s *GetUserViewMetricsResponseBody) SetTotal(v int32) *GetUserViewMetricsResponseBody {
	s.Total = &v
	return s
}

func (s *GetUserViewMetricsResponseBody) SetUserMetrics(v []*UserViewMetric) *GetUserViewMetricsResponseBody {
	s.UserMetrics = v
	return s
}

func (s *GetUserViewMetricsResponseBody) Validate() error {
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	if s.UserMetrics != nil {
		for _, item := range s.UserMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
