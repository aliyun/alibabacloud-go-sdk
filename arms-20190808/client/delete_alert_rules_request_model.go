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
	// The IDs of the alert rules that you want to delete. The value is a JSON array, for example, `[123, 234]`. You can call the SearchAlertRules operation and view the `Id` parameter in the response to obtain the alert rule ID. For more information, see [SearchAlertRules](https://help.aliyun.com/document_detail/175825.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// [123, 234]
	AlertIds *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
	// The region ID. Default value: `cn-hangzhou`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
