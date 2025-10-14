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
	Incident *GetIncidentResponseBodyIncident `json:"Incident,omitempty" xml:"Incident,omitempty" type:"Struct"`
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
	// example:
	//
	// 1757386075000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// dr-fy2zvgiykjifbiim****
	DetectionRuleId *string `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	// example:
	//
	// window
	IncidentAggregationType *string `json:"IncidentAggregationType,omitempty" xml:"IncidentAggregationType,omitempty"`
	// example:
	//
	// Forti incident desc
	IncidentDescription *string `json:"IncidentDescription,omitempty" xml:"IncidentDescription,omitempty"`
	// example:
	//
	// Forti
	IncidentName *string `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	// example:
	//
	// Remark
	IncidentRemark *string `json:"IncidentRemark,omitempty" xml:"IncidentRemark,omitempty"`
	// example:
	//
	// 0
	IncidentStatus *int32 `json:"IncidentStatus,omitempty" xml:"IncidentStatus,omitempty"`
	// example:
	//
	// ["sys:data_source:waf"]
	IncidentTags *string `json:"IncidentTags,omitempty" xml:"IncidentTags,omitempty"`
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// example:
	//
	// 23
	RelateAlertCount *int32 `json:"RelateAlertCount,omitempty" xml:"RelateAlertCount,omitempty"`
	// example:
	//
	// 2
	RelateAssetCount *int32 `json:"RelateAssetCount,omitempty" xml:"RelateAssetCount,omitempty"`
	// example:
	//
	// ["siem"]
	RelateDataSourceIds interface{} `json:"RelateDataSourceIds,omitempty" xml:"RelateDataSourceIds,omitempty"`
	// example:
	//
	// ["176618589410****","1130916744888****"]
	RelateUserIds interface{} `json:"RelateUserIds,omitempty" xml:"RelateUserIds,omitempty"`
	// example:
	//
	// 2
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	// example:
	//
	// 90
	ThreatScore *string `json:"ThreatScore,omitempty" xml:"ThreatScore,omitempty"`
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
