// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteMetricRulesRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteMetricRulesRequest
	GetDryRun() *bool
	SetId(v []*string) *DeleteMetricRulesRequest
	GetId() []*string
	SetRegionId(v string) *DeleteMetricRulesRequest
	GetRegionId() *string
}

type DeleteMetricRulesRequest struct {
	// Ensures the idempotence of the request. Generate a parameter value on your client to ensure that the value is unique across different requests. ClientToken supports only ASCII characters and can be up to 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// 6868E9A1-56E6-5FD7-9684-E2963BBAFB5F
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run only. Valid values:
	//
	// - **true**: Sends a check request and does not bind the security group to the instance. The check items include whether required parameters are specified, the request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends a normal request. If the check succeeds, an HTTP 2xx status code is returned and the operation is performed.
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

func (s *DeleteMetricRulesRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *DeleteMetricRulesRequest) SetClientToken(v string) *DeleteMetricRulesRequest {
	s.ClientToken = &v
	return s
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
