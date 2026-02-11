// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListPrometheusAlertRulesRequest
	GetClusterId() *string
	SetMatchExpressions(v string) *ListPrometheusAlertRulesRequest
	GetMatchExpressions() *string
	SetName(v string) *ListPrometheusAlertRulesRequest
	GetName() *string
	SetRegionId(v string) *ListPrometheusAlertRulesRequest
	GetRegionId() *string
	SetStatus(v int32) *ListPrometheusAlertRulesRequest
	GetStatus() *int32
	SetType(v string) *ListPrometheusAlertRulesRequest
	GetType() *string
}

type ListPrometheusAlertRulesRequest struct {
	// This parameter is required.
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MatchExpressions *string `json:"MatchExpressions,omitempty" xml:"MatchExpressions,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status   *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPrometheusAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPrometheusAlertRulesRequest) GetMatchExpressions() *string {
	return s.MatchExpressions
}

func (s *ListPrometheusAlertRulesRequest) GetName() *string {
	return s.Name
}

func (s *ListPrometheusAlertRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrometheusAlertRulesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListPrometheusAlertRulesRequest) GetType() *string {
	return s.Type
}

func (s *ListPrometheusAlertRulesRequest) SetClusterId(v string) *ListPrometheusAlertRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetMatchExpressions(v string) *ListPrometheusAlertRulesRequest {
	s.MatchExpressions = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetName(v string) *ListPrometheusAlertRulesRequest {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetRegionId(v string) *ListPrometheusAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetStatus(v int32) *ListPrometheusAlertRulesRequest {
	s.Status = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetType(v string) *ListPrometheusAlertRulesRequest {
	s.Type = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) Validate() error {
	return dara.Validate(s)
}
