// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertIds(v string) *GetAlertRulesRequest
	GetAlertIds() *string
	SetAlertNames(v string) *GetAlertRulesRequest
	GetAlertNames() *string
	SetAlertStatus(v string) *GetAlertRulesRequest
	GetAlertStatus() *string
	SetAlertType(v string) *GetAlertRulesRequest
	GetAlertType() *string
	SetClusterId(v string) *GetAlertRulesRequest
	GetClusterId() *string
	SetPage(v int64) *GetAlertRulesRequest
	GetPage() *int64
	SetProductCode(v string) *GetAlertRulesRequest
	GetProductCode() *string
	SetRegionId(v string) *GetAlertRulesRequest
	GetRegionId() *string
	SetSize(v int64) *GetAlertRulesRequest
	GetSize() *int64
	SetTags(v []*GetAlertRulesRequestTags) *GetAlertRulesRequest
	GetTags() []*GetAlertRulesRequestTags
}

type GetAlertRulesRequest struct {
	// The unique IDs of alert rules.
	//
	// 	- If you do not specify this parameter, the API operation does not filter alert rules based on their IDs.
	//
	// 	- If you specify this parameter, the API operation returns only the information of the specified alert rules. Other filter conditions also take effect.
	//
	// > When you call the GetAlertRules operation, you can specify other request parameters to obtain the AlertIds parameter from the response. Then, you can specify the AlertIds parameter to query the specified alert rules.
	//
	// example:
	//
	// ["12345"]
	AlertIds *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
	// The names of alert rules. When you create alert rules of the new version, you cannot specify duplicate names. However, existing alert rules may have duplicate names. Therefore, the **AlertName*	- parameter does not uniquely identify an alert rule.
	//
	// 	- If you do not specify this parameter, the API operation does not filter alert rules based on their names.
	//
	// 	- If you specify this parameter, the API operation returns only the information of the specified alert rules. Other filter conditions also take effect.
	//
	// example:
	//
	// ["test"]
	AlertNames *string `json:"AlertNames,omitempty" xml:"AlertNames,omitempty"`
	// The status of the alert rule. Valid values:
	//
	// 	- RUNNING
	//
	// 	- STOPPED
	//
	// 	- PAUSED
	//
	// >  The PAUSED state indicates an abnormal and paused alert rule. This may result from excessively large threshold values or deleted associated clusters.
	//
	// example:
	//
	// RUNNING
	AlertStatus *string `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	// The type of the alert rule. This parameter is required for the new version of Alert Management. Valid values:
	//
	// 	- APPLICATION_MONITORING_ALERT_RULE: alert rule for Application Monitoring
	//
	// 	- BROWSER_MONITORING_ALERT_RULE: alert rule for Browser Monitoring
	//
	// 	- PROMETHEUS_MONITORING_ALERT_RULE: alert rule for Managed Service for Prometheus
	//
	// example:
	//
	// APPLICATION_MONITORING_ALERT_RULE
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The ID of the monitored cluster.
	//
	// example:
	//
	// ceba9b9ea5b924dd0b6726d2de6******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// You do not need to configure this parameter.
	//
	// example:
	//
	// null
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of alert rules to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The list of tags.
	Tags []*GetAlertRulesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *GetAlertRulesRequest) GetAlertIds() *string {
	return s.AlertIds
}

func (s *GetAlertRulesRequest) GetAlertNames() *string {
	return s.AlertNames
}

func (s *GetAlertRulesRequest) GetAlertStatus() *string {
	return s.AlertStatus
}

func (s *GetAlertRulesRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *GetAlertRulesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetAlertRulesRequest) GetPage() *int64 {
	return s.Page
}

func (s *GetAlertRulesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetAlertRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAlertRulesRequest) GetSize() *int64 {
	return s.Size
}

func (s *GetAlertRulesRequest) GetTags() []*GetAlertRulesRequestTags {
	return s.Tags
}

func (s *GetAlertRulesRequest) SetAlertIds(v string) *GetAlertRulesRequest {
	s.AlertIds = &v
	return s
}

func (s *GetAlertRulesRequest) SetAlertNames(v string) *GetAlertRulesRequest {
	s.AlertNames = &v
	return s
}

func (s *GetAlertRulesRequest) SetAlertStatus(v string) *GetAlertRulesRequest {
	s.AlertStatus = &v
	return s
}

func (s *GetAlertRulesRequest) SetAlertType(v string) *GetAlertRulesRequest {
	s.AlertType = &v
	return s
}

func (s *GetAlertRulesRequest) SetClusterId(v string) *GetAlertRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *GetAlertRulesRequest) SetPage(v int64) *GetAlertRulesRequest {
	s.Page = &v
	return s
}

func (s *GetAlertRulesRequest) SetProductCode(v string) *GetAlertRulesRequest {
	s.ProductCode = &v
	return s
}

func (s *GetAlertRulesRequest) SetRegionId(v string) *GetAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *GetAlertRulesRequest) SetSize(v int64) *GetAlertRulesRequest {
	s.Size = &v
	return s
}

func (s *GetAlertRulesRequest) SetTags(v []*GetAlertRulesRequestTags) *GetAlertRulesRequest {
	s.Tags = v
	return s
}

func (s *GetAlertRulesRequest) Validate() error {
	return dara.Validate(s)
}

type GetAlertRulesRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// type
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAlertRulesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesRequestTags) GoString() string {
	return s.String()
}

func (s *GetAlertRulesRequestTags) GetKey() *string {
	return s.Key
}

func (s *GetAlertRulesRequestTags) GetValue() *string {
	return s.Value
}

func (s *GetAlertRulesRequestTags) SetKey(v string) *GetAlertRulesRequestTags {
	s.Key = &v
	return s
}

func (s *GetAlertRulesRequestTags) SetValue(v string) *GetAlertRulesRequestTags {
	s.Value = &v
	return s
}

func (s *GetAlertRulesRequestTags) Validate() error {
	return dara.Validate(s)
}
