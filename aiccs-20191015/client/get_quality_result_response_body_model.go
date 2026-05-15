// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelTypeName(v string) *GetQualityResultResponseBody
	GetChannelTypeName() *string
	SetCode(v string) *GetQualityResultResponseBody
	GetCode() *string
	SetData(v *GetQualityResultResponseBodyData) *GetQualityResultResponseBody
	GetData() *GetQualityResultResponseBodyData
	SetMessage(v string) *GetQualityResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityResultResponseBody
	GetSuccess() *bool
}

type GetQualityResultResponseBody struct {
	ChannelTypeName *string                           `json:"ChannelTypeName,omitempty" xml:"ChannelTypeName,omitempty"`
	Code            *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            *GetQualityResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message         *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityResultResponseBody) GetChannelTypeName() *string {
	return s.ChannelTypeName
}

func (s *GetQualityResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityResultResponseBody) GetData() *GetQualityResultResponseBodyData {
	return s.Data
}

func (s *GetQualityResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityResultResponseBody) SetChannelTypeName(v string) *GetQualityResultResponseBody {
	s.ChannelTypeName = &v
	return s
}

func (s *GetQualityResultResponseBody) SetCode(v string) *GetQualityResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityResultResponseBody) SetData(v *GetQualityResultResponseBodyData) *GetQualityResultResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityResultResponseBody) SetMessage(v string) *GetQualityResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityResultResponseBody) SetRequestId(v string) *GetQualityResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityResultResponseBody) SetSuccess(v bool) *GetQualityResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityResultResponseBodyData struct {
	PageNo                    *int32                                                       `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize                  *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QualityResultResponseList []*GetQualityResultResponseBodyDataQualityResultResponseList `json:"QualityResultResponseList,omitempty" xml:"QualityResultResponseList,omitempty" type:"Repeated"`
	TotalNum                  *int32                                                       `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetQualityResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityResultResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQualityResultResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQualityResultResponseBodyData) GetQualityResultResponseList() []*GetQualityResultResponseBodyDataQualityResultResponseList {
	return s.QualityResultResponseList
}

func (s *GetQualityResultResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetQualityResultResponseBodyData) SetPageNo(v int32) *GetQualityResultResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQualityResultResponseBodyData) SetPageSize(v int32) *GetQualityResultResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQualityResultResponseBodyData) SetQualityResultResponseList(v []*GetQualityResultResponseBodyDataQualityResultResponseList) *GetQualityResultResponseBodyData {
	s.QualityResultResponseList = v
	return s
}

func (s *GetQualityResultResponseBodyData) SetTotalNum(v int32) *GetQualityResultResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetQualityResultResponseBodyData) Validate() error {
	if s.QualityResultResponseList != nil {
		for _, item := range s.QualityResultResponseList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQualityResultResponseBodyDataQualityResultResponseList struct {
	ChannelType     *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	ChannelTypeName *string `json:"ChannelTypeName,omitempty" xml:"ChannelTypeName,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HitDetail       *string `json:"HitDetail,omitempty" xml:"HitDetail,omitempty"`
	HitStatus       *bool   `json:"HitStatus,omitempty" xml:"HitStatus,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MemberName      *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RuleId          *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	ServicerId      *string `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	ServicerName    *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	TouchId         *string `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
	TouchStartTime  *string `json:"TouchStartTime,omitempty" xml:"TouchStartTime,omitempty"`
}

func (s GetQualityResultResponseBodyDataQualityResultResponseList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityResultResponseBodyDataQualityResultResponseList) GoString() string {
	return s.String()
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetChannelType() *string {
	return s.ChannelType
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetChannelTypeName() *string {
	return s.ChannelTypeName
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetGroupId() *string {
	return s.GroupId
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetGroupName() *string {
	return s.GroupName
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetHitDetail() *string {
	return s.HitDetail
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetHitStatus() *bool {
	return s.HitStatus
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetMemberName() *string {
	return s.MemberName
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetRuleId() *string {
	return s.RuleId
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetRuleName() *string {
	return s.RuleName
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetServicerId() *string {
	return s.ServicerId
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetServicerName() *string {
	return s.ServicerName
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetTouchId() *string {
	return s.TouchId
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) GetTouchStartTime() *string {
	return s.TouchStartTime
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetChannelType(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.ChannelType = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetChannelTypeName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.ChannelTypeName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetGroupId(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.GroupId = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetGroupName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.GroupName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetHitDetail(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.HitDetail = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetHitStatus(v bool) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.HitStatus = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetInstanceName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.InstanceName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetMemberName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.MemberName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetProjectId(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.ProjectId = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetProjectName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.ProjectName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetRuleId(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.RuleId = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetRuleName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.RuleName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetServicerId(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.ServicerId = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetServicerName(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.ServicerName = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetTouchId(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.TouchId = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) SetTouchStartTime(v string) *GetQualityResultResponseBodyDataQualityResultResponseList {
	s.TouchStartTime = &v
	return s
}

func (s *GetQualityResultResponseBodyDataQualityResultResponseList) Validate() error {
	return dara.Validate(s)
}
