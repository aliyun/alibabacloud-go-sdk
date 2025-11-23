// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderBaseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetOrderBaseInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetOrderBaseInfoResponseBody
	GetErrorMessage() *string
	SetOrderBaseInfo(v *GetOrderBaseInfoResponseBodyOrderBaseInfo) *GetOrderBaseInfoResponseBody
	GetOrderBaseInfo() *GetOrderBaseInfoResponseBodyOrderBaseInfo
	SetRequestId(v string) *GetOrderBaseInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOrderBaseInfoResponseBody
	GetSuccess() *bool
}

type GetOrderBaseInfoResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The basic information about the ticket.
	OrderBaseInfo *GetOrderBaseInfoResponseBodyOrderBaseInfo `json:"OrderBaseInfo,omitempty" xml:"OrderBaseInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7133DF67-5B25-460F-8285-C4CC93472C2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrderBaseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrderBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrderBaseInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetOrderBaseInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetOrderBaseInfoResponseBody) GetOrderBaseInfo() *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	return s.OrderBaseInfo
}

func (s *GetOrderBaseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrderBaseInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOrderBaseInfoResponseBody) SetErrorCode(v string) *GetOrderBaseInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOrderBaseInfoResponseBody) SetErrorMessage(v string) *GetOrderBaseInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOrderBaseInfoResponseBody) SetOrderBaseInfo(v *GetOrderBaseInfoResponseBodyOrderBaseInfo) *GetOrderBaseInfoResponseBody {
	s.OrderBaseInfo = v
	return s
}

func (s *GetOrderBaseInfoResponseBody) SetRequestId(v string) *GetOrderBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrderBaseInfoResponseBody) SetSuccess(v bool) *GetOrderBaseInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetOrderBaseInfoResponseBody) Validate() error {
	if s.OrderBaseInfo != nil {
		if err := s.OrderBaseInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrderBaseInfoResponseBodyOrderBaseInfo struct {
	// The Key of the ticket attachment. This information is returned only when an attachment is uploaded when a ticket is created.
	//
	// example:
	//
	// upload_order_info_856887_f356366f-f0f8-42fc-ba57-4a509303e814_18072d8a9bce876e3073bc655c2865f.png
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The remarks of the ticket.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The applicant.
	//
	// example:
	//
	// xxx
	Committer *string `json:"Committer,omitempty" xml:"Committer,omitempty"`
	// The ID of the applicant. Note: The ID is different from the Alibaba Cloud account ID of the applicant.
	//
	// example:
	//
	// 1
	CommitterId *int64 `json:"CommitterId,omitempty" xml:"CommitterId,omitempty"`
	// The time when the ticket was created.
	//
	// example:
	//
	// 2019-10-10 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the ticket was last modified.
	//
	// example:
	//
	// 2019-10-10 00:00:00
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The ID of the ticket.
	//
	// example:
	//
	// 12345
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The original file name of the ticket attachment. This information is returned only when an attachment is uploaded when a ticket is created.
	//
	// example:
	//
	// 18072d8a9bce876e3073bc655c2865f.png
	OriginAttachmentName *string `json:"OriginAttachmentName,omitempty" xml:"OriginAttachmentName,omitempty"`
	// The type of the ticket. For more information about the value of this parameter, see the request parameters of the [CreateOrder](https://help.aliyun.com/document_detail/465865.html) operation.
	//
	// example:
	//
	// DC_COMMON
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// The IDs of the operators that are related to the ticket.
	RelatedUserList *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Struct"`
	// The nicknames of the operators that are related to the ticket.
	RelatedUserNickList *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList `json:"RelatedUserNickList,omitempty" xml:"RelatedUserNickList,omitempty" type:"Struct"`
	// The status code of the ticket. Valid values:
	//
	// 	- **new**: The ticket is created.
	//
	// 	- **toaudit**: The ticket is being reviewed.
	//
	// 	- **Approved**: The ticket is approved.
	//
	// 	- **reject**: The ticket is rejected.
	//
	// 	- **processing**: The ticket is being executed.
	//
	// 	- **success**: The ticket is executed.
	//
	// 	- **closed**: The ticket is closed.
	//
	// example:
	//
	// success
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The description of the status.
	//
	// example:
	//
	// success
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The ID of the approval process.
	//
	// example:
	//
	// 1
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	// The description of the approval process.
	//
	// example:
	//
	// approved
	WorkflowStatusDesc *string `json:"WorkflowStatusDesc,omitempty" xml:"WorkflowStatusDesc,omitempty"`
}

func (s GetOrderBaseInfoResponseBodyOrderBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s GetOrderBaseInfoResponseBodyOrderBaseInfo) GoString() string {
	return s.String()
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetComment() *string {
	return s.Comment
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetCommitter() *string {
	return s.Committer
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetCommitterId() *int64 {
	return s.CommitterId
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetLastModifyTime() *string {
	return s.LastModifyTime
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetOriginAttachmentName() *string {
	return s.OriginAttachmentName
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetPluginType() *string {
	return s.PluginType
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetRelatedUserList() *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList {
	return s.RelatedUserList
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetRelatedUserNickList() *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList {
	return s.RelatedUserNickList
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetStatusCode() *string {
	return s.StatusCode
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) GetWorkflowStatusDesc() *string {
	return s.WorkflowStatusDesc
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetAttachmentKey(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.AttachmentKey = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetComment(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.Comment = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetCommitter(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.Committer = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetCommitterId(v int64) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.CommitterId = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetCreateTime(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetLastModifyTime(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.LastModifyTime = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetOrderId(v int64) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.OrderId = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetOriginAttachmentName(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.OriginAttachmentName = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetPluginType(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.PluginType = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetRelatedUserList(v *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.RelatedUserList = v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetRelatedUserNickList(v *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.RelatedUserNickList = v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetStatusCode(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.StatusCode = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetStatusDesc(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.StatusDesc = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetWorkflowInstanceId(v int64) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.WorkflowInstanceId = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetWorkflowStatusDesc(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.WorkflowStatusDesc = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) Validate() error {
	if s.RelatedUserList != nil {
		if err := s.RelatedUserList.Validate(); err != nil {
			return err
		}
	}
	if s.RelatedUserNickList != nil {
		if err := s.RelatedUserNickList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList struct {
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList) String() string {
	return dara.Prettify(s)
}

func (s GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList) GoString() string {
	return s.String()
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList) GetUserIds() []*string {
	return s.UserIds
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList) SetUserIds(v []*string) *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList {
	s.UserIds = v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList) Validate() error {
	return dara.Validate(s)
}

type GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList struct {
	UserNicks []*string `json:"UserNicks,omitempty" xml:"UserNicks,omitempty" type:"Repeated"`
}

func (s GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList) String() string {
	return dara.Prettify(s)
}

func (s GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList) GoString() string {
	return s.String()
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList) GetUserNicks() []*string {
	return s.UserNicks
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList) SetUserNicks(v []*string) *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList {
	s.UserNicks = v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList) Validate() error {
	return dara.Validate(s)
}
