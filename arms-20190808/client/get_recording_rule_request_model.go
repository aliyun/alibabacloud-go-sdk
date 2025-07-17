// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetRecordingRuleRequest
	GetClusterId() *string
	SetRegionId(v string) *GetRecordingRuleRequest
	GetRegionId() *string
}

type GetRecordingRuleRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetRecordingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecordingRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRecordingRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetRecordingRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRecordingRuleRequest) SetClusterId(v string) *GetRecordingRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *GetRecordingRuleRequest) SetRegionId(v string) *GetRecordingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *GetRecordingRuleRequest) Validate() error {
	return dara.Validate(s)
}
