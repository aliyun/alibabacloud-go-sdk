// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIncidentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIncident(v *GetIncidentResponseBodyIncident) *GetIncidentResponseBody
	GetIncident() *GetIncidentResponseBodyIncident
	SetRequestId(v string) *GetIncidentResponseBody
	GetRequestId() *string
}

type GetIncidentResponseBody struct {
	// The event information.
	Incident *GetIncidentResponseBodyIncident `json:"Incident,omitempty" xml:"Incident,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIncidentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIncidentResponseBody) GoString() string {
	return s.String()
}

func (s *GetIncidentResponseBody) GetIncident() *GetIncidentResponseBodyIncident {
	return s.Incident
}

func (s *GetIncidentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIncidentResponseBody) SetIncident(v *GetIncidentResponseBodyIncident) *GetIncidentResponseBody {
	s.Incident = v
	return s
}

func (s *GetIncidentResponseBody) SetRequestId(v string) *GetIncidentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIncidentResponseBody) Validate() error {
	if s.Incident != nil {
		if err := s.Incident.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIncidentResponseBodyIncident struct {
	// The count of attack stages associated with the event alerts.
	//
	// example:
	//
	// {
	//
	// 	"AttckTactics": [
	//
	// 		{
	//
	// 			"tacticName": "Reconnaissance",
	//
	// 			"alertNum": 0,
	//
	// 			"tacticId": "TA0040"
	//
	// 		}
	//
	// 	]
	//
	// }
	AttckTactics interface{} `json:"AttckTactics,omitempty" xml:"AttckTactics,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1757386075000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the detection rule.
	//
	// example:
	//
	// dr-fy2zvgiykjifbiim****
	DetectionRuleId *string `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	// The event summaries type. Valid values:
	//
	// - none: no event is generated.
	//
	// - graph_compute: graph computing (supported by predefined rules).
	//
	// - expert: expert rule.
	//
	// - passthrough: alerting pass-through (one-to-one).
	//
	// - window: same-type aggregation (window).
	//
	// example:
	//
	// window
	IncidentAggregationType *string `json:"IncidentAggregationType,omitempty" xml:"IncidentAggregationType,omitempty"`
	// The description of the event.
	//
	// example:
	//
	// Forti incident desc
	IncidentDescription *string `json:"IncidentDescription,omitempty" xml:"IncidentDescription,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// Forti
	IncidentName *string `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	// The remarks of the event.
	//
	// example:
	//
	// Remark
	IncidentRemark *string `json:"IncidentRemark,omitempty" xml:"IncidentRemark,omitempty"`
	// The status of the event. Valid values:
	//
	// - 0: unhandled.
	//
	// - 1: handling.
	//
	// - 5: handling failed.
	//
	// - 10: handled.
	//
	// example:
	//
	// 0
	IncidentStatus *int32 `json:"IncidentStatus,omitempty" xml:"IncidentStatus,omitempty"`
	// The tags of the event.
	//
	// example:
	//
	// ["sys:data_source:waf"]
	IncidentTags *string `json:"IncidentTags,omitempty" xml:"IncidentTags,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The owner of the event.
	//
	// example:
	//
	// 1234567890xxxxxx
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The number of alerts associated with the event.
	//
	// example:
	//
	// 23
	RelateAlertCount *int32 `json:"RelateAlertCount,omitempty" xml:"RelateAlertCount,omitempty"`
	// The number of assets associated with the event.
	//
	// example:
	//
	// 2
	RelateAssetCount *int32 `json:"RelateAssetCount,omitempty" xml:"RelateAssetCount,omitempty"`
	// The list of associated data sources.
	//
	// example:
	//
	// ["siem"]
	RelateDataSourceIds interface{} `json:"RelateDataSourceIds,omitempty" xml:"RelateDataSourceIds,omitempty"`
	// The list of user IDs associated with the event.
	//
	// example:
	//
	// ["176618589410****","1130916744888****"]
	RelateUserIds interface{} `json:"RelateUserIds,omitempty" xml:"RelateUserIds,omitempty"`
	// The response time. Unit: milliseconds (ms).
	//
	// example:
	//
	// 1757386075000
	ResponseTime *int64 `json:"ResponseTime,omitempty" xml:"ResponseTime,omitempty"`
	// The threat level. Valid values:
	//
	// - 5: critical.
	//
	// - 4: high.
	//
	// - 3: medium.
	//
	// - 2: low.
	//
	// - 1: informational.
	//
	// example:
	//
	// 2
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	// The threat score of the event. Valid values: 0 to 100. A higher score indicates a higher risk level.
	//
	// example:
	//
	// 90
	ThreatScore *string `json:"ThreatScore,omitempty" xml:"ThreatScore,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1757386075000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetIncidentResponseBodyIncident) String() string {
	return dara.Prettify(s)
}

func (s GetIncidentResponseBodyIncident) GoString() string {
	return s.String()
}

func (s *GetIncidentResponseBodyIncident) GetAttckTactics() interface{} {
	return s.AttckTactics
}

func (s *GetIncidentResponseBodyIncident) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetIncidentResponseBodyIncident) GetDetectionRuleId() *string {
	return s.DetectionRuleId
}

func (s *GetIncidentResponseBodyIncident) GetIncidentAggregationType() *string {
	return s.IncidentAggregationType
}

func (s *GetIncidentResponseBodyIncident) GetIncidentDescription() *string {
	return s.IncidentDescription
}

func (s *GetIncidentResponseBodyIncident) GetIncidentName() *string {
	return s.IncidentName
}

func (s *GetIncidentResponseBodyIncident) GetIncidentRemark() *string {
	return s.IncidentRemark
}

func (s *GetIncidentResponseBodyIncident) GetIncidentStatus() *int32 {
	return s.IncidentStatus
}

func (s *GetIncidentResponseBodyIncident) GetIncidentTags() *string {
	return s.IncidentTags
}

func (s *GetIncidentResponseBodyIncident) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *GetIncidentResponseBodyIncident) GetOwner() *string {
	return s.Owner
}

func (s *GetIncidentResponseBodyIncident) GetRelateAlertCount() *int32 {
	return s.RelateAlertCount
}

func (s *GetIncidentResponseBodyIncident) GetRelateAssetCount() *int32 {
	return s.RelateAssetCount
}

func (s *GetIncidentResponseBodyIncident) GetRelateDataSourceIds() interface{} {
	return s.RelateDataSourceIds
}

func (s *GetIncidentResponseBodyIncident) GetRelateUserIds() interface{} {
	return s.RelateUserIds
}

func (s *GetIncidentResponseBodyIncident) GetResponseTime() *int64 {
	return s.ResponseTime
}

func (s *GetIncidentResponseBodyIncident) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *GetIncidentResponseBodyIncident) GetThreatScore() *string {
	return s.ThreatScore
}

func (s *GetIncidentResponseBodyIncident) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetIncidentResponseBodyIncident) SetAttckTactics(v interface{}) *GetIncidentResponseBodyIncident {
	s.AttckTactics = v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetCreateTime(v int64) *GetIncidentResponseBodyIncident {
	s.CreateTime = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetDetectionRuleId(v string) *GetIncidentResponseBodyIncident {
	s.DetectionRuleId = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetIncidentAggregationType(v string) *GetIncidentResponseBodyIncident {
	s.IncidentAggregationType = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetIncidentDescription(v string) *GetIncidentResponseBodyIncident {
	s.IncidentDescription = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetIncidentName(v string) *GetIncidentResponseBodyIncident {
	s.IncidentName = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetIncidentRemark(v string) *GetIncidentResponseBodyIncident {
	s.IncidentRemark = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetIncidentStatus(v int32) *GetIncidentResponseBodyIncident {
	s.IncidentStatus = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetIncidentTags(v string) *GetIncidentResponseBodyIncident {
	s.IncidentTags = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetIncidentUuid(v string) *GetIncidentResponseBodyIncident {
	s.IncidentUuid = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetOwner(v string) *GetIncidentResponseBodyIncident {
	s.Owner = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetRelateAlertCount(v int32) *GetIncidentResponseBodyIncident {
	s.RelateAlertCount = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetRelateAssetCount(v int32) *GetIncidentResponseBodyIncident {
	s.RelateAssetCount = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetRelateDataSourceIds(v interface{}) *GetIncidentResponseBodyIncident {
	s.RelateDataSourceIds = v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetRelateUserIds(v interface{}) *GetIncidentResponseBodyIncident {
	s.RelateUserIds = v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetResponseTime(v int64) *GetIncidentResponseBodyIncident {
	s.ResponseTime = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetThreatLevel(v string) *GetIncidentResponseBodyIncident {
	s.ThreatLevel = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetThreatScore(v string) *GetIncidentResponseBodyIncident {
	s.ThreatScore = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) SetUpdateTime(v int64) *GetIncidentResponseBodyIncident {
	s.UpdateTime = &v
	return s
}

func (s *GetIncidentResponseBodyIncident) Validate() error {
	return dara.Validate(s)
}
