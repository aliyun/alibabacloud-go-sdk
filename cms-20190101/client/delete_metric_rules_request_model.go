// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DeleteMetricRulesRequest
	GetDryRun() *bool
	SetId(v []*string) *DeleteMetricRulesRequest
	GetId() []*string
	SetRegionId(v string) *DeleteMetricRulesRequest
	GetRegionId() *string
}

type DeleteMetricRulesRequest struct {
	// Specifies whether to perform a dry run of the request. Valid values:
	//
	// - **true**: performs a dry run. No actual operation is performed. The system checks whether the required parameters are specified, the request format is valid, and the service limits are met. If the request fails the dry run, an error is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IDs of the alert rules to delete.
	//
	// This parameter is required.
	Id       []*string `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMetricRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetricRulesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteMetricRulesRequest) GetId() []*string {
	return s.Id
}

func (s *DeleteMetricRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMetricRulesRequest) SetDryRun(v bool) *DeleteMetricRulesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteMetricRulesRequest) SetId(v []*string) *DeleteMetricRulesRequest {
	s.Id = v
	return s
}

func (s *DeleteMetricRulesRequest) SetRegionId(v string) *DeleteMetricRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMetricRulesRequest) Validate() error {
	return dara.Validate(s)
}
