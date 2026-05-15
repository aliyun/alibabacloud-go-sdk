// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *GetQualityResultRequest
	GetChannelType() *string
	SetGroupIds(v []*int64) *GetQualityResultRequest
	GetGroupIds() []*int64
	SetHitStatus(v int32) *GetQualityResultRequest
	GetHitStatus() *int32
	SetInstanceId(v string) *GetQualityResultRequest
	GetInstanceId() *string
	SetPageNo(v int32) *GetQualityResultRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetQualityResultRequest
	GetPageSize() *int32
	SetProjectIds(v []*int64) *GetQualityResultRequest
	GetProjectIds() []*int64
	SetQualityRuleIds(v []*int64) *GetQualityResultRequest
	GetQualityRuleIds() []*int64
	SetTouchEndTime(v string) *GetQualityResultRequest
	GetTouchEndTime() *string
	SetTouchStartTime(v string) *GetQualityResultRequest
	GetTouchStartTime() *string
}

type GetQualityResultRequest struct {
	ChannelType *string  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	GroupIds    []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	HitStatus   *int32   `json:"HitStatus,omitempty" xml:"HitStatus,omitempty"`
	// This parameter is required.
	InstanceId     *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo         *int32   `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectIds     []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	QualityRuleIds []*int64 `json:"QualityRuleIds,omitempty" xml:"QualityRuleIds,omitempty" type:"Repeated"`
	// This parameter is required.
	TouchEndTime *string `json:"TouchEndTime,omitempty" xml:"TouchEndTime,omitempty"`
	// This parameter is required.
	TouchStartTime *string `json:"TouchStartTime,omitempty" xml:"TouchStartTime,omitempty"`
}

func (s GetQualityResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityResultRequest) GoString() string {
	return s.String()
}

func (s *GetQualityResultRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *GetQualityResultRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetQualityResultRequest) GetHitStatus() *int32 {
	return s.HitStatus
}

func (s *GetQualityResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQualityResultRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQualityResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQualityResultRequest) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *GetQualityResultRequest) GetQualityRuleIds() []*int64 {
	return s.QualityRuleIds
}

func (s *GetQualityResultRequest) GetTouchEndTime() *string {
	return s.TouchEndTime
}

func (s *GetQualityResultRequest) GetTouchStartTime() *string {
	return s.TouchStartTime
}

func (s *GetQualityResultRequest) SetChannelType(v string) *GetQualityResultRequest {
	s.ChannelType = &v
	return s
}

func (s *GetQualityResultRequest) SetGroupIds(v []*int64) *GetQualityResultRequest {
	s.GroupIds = v
	return s
}

func (s *GetQualityResultRequest) SetHitStatus(v int32) *GetQualityResultRequest {
	s.HitStatus = &v
	return s
}

func (s *GetQualityResultRequest) SetInstanceId(v string) *GetQualityResultRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityResultRequest) SetPageNo(v int32) *GetQualityResultRequest {
	s.PageNo = &v
	return s
}

func (s *GetQualityResultRequest) SetPageSize(v int32) *GetQualityResultRequest {
	s.PageSize = &v
	return s
}

func (s *GetQualityResultRequest) SetProjectIds(v []*int64) *GetQualityResultRequest {
	s.ProjectIds = v
	return s
}

func (s *GetQualityResultRequest) SetQualityRuleIds(v []*int64) *GetQualityResultRequest {
	s.QualityRuleIds = v
	return s
}

func (s *GetQualityResultRequest) SetTouchEndTime(v string) *GetQualityResultRequest {
	s.TouchEndTime = &v
	return s
}

func (s *GetQualityResultRequest) SetTouchStartTime(v string) *GetQualityResultRequest {
	s.TouchStartTime = &v
	return s
}

func (s *GetQualityResultRequest) Validate() error {
	return dara.Validate(s)
}
