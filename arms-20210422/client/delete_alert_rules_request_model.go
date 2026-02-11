// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertIds(v string) *DeleteAlertRulesRequest
	GetAlertIds() *string
	SetRegionId(v string) *DeleteAlertRulesRequest
	GetRegionId() *string
}

type DeleteAlertRulesRequest struct {
	// This parameter is required.
	AlertIds *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertRulesRequest) GetAlertIds() *string {
	return s.AlertIds
}

func (s *DeleteAlertRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAlertRulesRequest) SetAlertIds(v string) *DeleteAlertRulesRequest {
	s.AlertIds = &v
	return s
}

func (s *DeleteAlertRulesRequest) SetRegionId(v string) *DeleteAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAlertRulesRequest) Validate() error {
	return dara.Validate(s)
}
