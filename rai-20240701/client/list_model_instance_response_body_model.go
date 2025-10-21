// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListModelInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListModelInstanceResponseBody
	GetMessage() *string
	SetModelInstanceInfoList(v []*ListModelInstanceResponseBodyModelInstanceInfoList) *ListModelInstanceResponseBody
	GetModelInstanceInfoList() []*ListModelInstanceResponseBodyModelInstanceInfoList
	SetPageNumber(v int32) *ListModelInstanceResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelInstanceResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListModelInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListModelInstanceResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListModelInstanceResponseBody
	GetTotalCount() *int32
}

type ListModelInstanceResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message               *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	ModelInstanceInfoList []*ListModelInstanceResponseBodyModelInstanceInfoList `json:"ModelInstanceInfoList,omitempty" xml:"ModelInstanceInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListModelInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListModelInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListModelInstanceResponseBody) GetModelInstanceInfoList() []*ListModelInstanceResponseBodyModelInstanceInfoList {
	return s.ModelInstanceInfoList
}

func (s *ListModelInstanceResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelInstanceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListModelInstanceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListModelInstanceResponseBody) SetCode(v string) *ListModelInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelInstanceResponseBody) SetHttpStatusCode(v int32) *ListModelInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListModelInstanceResponseBody) SetMessage(v string) *ListModelInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ListModelInstanceResponseBody) SetModelInstanceInfoList(v []*ListModelInstanceResponseBodyModelInstanceInfoList) *ListModelInstanceResponseBody {
	s.ModelInstanceInfoList = v
	return s
}

func (s *ListModelInstanceResponseBody) SetPageNumber(v int32) *ListModelInstanceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModelInstanceResponseBody) SetPageSize(v int32) *ListModelInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListModelInstanceResponseBody) SetRequestId(v string) *ListModelInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelInstanceResponseBody) SetSuccess(v bool) *ListModelInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ListModelInstanceResponseBody) SetTotalCount(v int32) *ListModelInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelInstanceResponseBody) Validate() error {
	if s.ModelInstanceInfoList != nil {
		for _, item := range s.ModelInstanceInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelInstanceResponseBodyModelInstanceInfoList struct {
	// example:
	//
	// torch_rank_v1
	EasServiceName *string `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	// example:
	//
	// 1749450490000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// True
	IsSupportContentSafe *int32 `json:"IsSupportContentSafe,omitempty" xml:"IsSupportContentSafe,omitempty"`
	// example:
	//
	// False
	IsSupportImage *bool `json:"IsSupportImage,omitempty" xml:"IsSupportImage,omitempty"`
	// example:
	//
	// False
	IsSupportPromptAttack *int32 `json:"IsSupportPromptAttack,omitempty" xml:"IsSupportPromptAttack,omitempty"`
	// example:
	//
	// True
	IsSupportSensitiveTopic *int32 `json:"IsSupportSensitiveTopic,omitempty" xml:"IsSupportSensitiveTopic,omitempty"`
	// example:
	//
	// True
	IsSupportText *bool `json:"IsSupportText,omitempty" xml:"IsSupportText,omitempty"`
	// example:
	//
	// 123
	ModelInstanceId *int64 `json:"ModelInstanceId,omitempty" xml:"ModelInstanceId,omitempty"`
	// example:
	//
	// 1
	ModelSource *int32 `json:"ModelSource,omitempty" xml:"ModelSource,omitempty"`
	// example:
	//
	// 608226
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListModelInstanceResponseBodyModelInstanceInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListModelInstanceResponseBodyModelInstanceInfoList) GoString() string {
	return s.String()
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) GetEasServiceName() *string {
	return s.EasServiceName
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) GetIsSupportContentSafe() *int32 {
	return s.IsSupportContentSafe
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) GetIsSupportImage() *bool {
	return s.IsSupportImage
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) GetIsSupportPromptAttack() *int32 {
	return s.IsSupportPromptAttack
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) GetIsSupportSensitiveTopic() *int32 {
	return s.IsSupportSensitiveTopic
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) GetIsSupportText() *bool {
	return s.IsSupportText
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) GetModelInstanceId() *int64 {
	return s.ModelInstanceId
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) GetModelSource() *int32 {
	return s.ModelSource
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) SetEasServiceName(v string) *ListModelInstanceResponseBodyModelInstanceInfoList {
	s.EasServiceName = &v
	return s
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) SetGmtModified(v int64) *ListModelInstanceResponseBodyModelInstanceInfoList {
	s.GmtModified = &v
	return s
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) SetIsSupportContentSafe(v int32) *ListModelInstanceResponseBodyModelInstanceInfoList {
	s.IsSupportContentSafe = &v
	return s
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) SetIsSupportImage(v bool) *ListModelInstanceResponseBodyModelInstanceInfoList {
	s.IsSupportImage = &v
	return s
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) SetIsSupportPromptAttack(v int32) *ListModelInstanceResponseBodyModelInstanceInfoList {
	s.IsSupportPromptAttack = &v
	return s
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) SetIsSupportSensitiveTopic(v int32) *ListModelInstanceResponseBodyModelInstanceInfoList {
	s.IsSupportSensitiveTopic = &v
	return s
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) SetIsSupportText(v bool) *ListModelInstanceResponseBodyModelInstanceInfoList {
	s.IsSupportText = &v
	return s
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) SetModelInstanceId(v int64) *ListModelInstanceResponseBodyModelInstanceInfoList {
	s.ModelInstanceId = &v
	return s
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) SetModelSource(v int32) *ListModelInstanceResponseBodyModelInstanceInfoList {
	s.ModelSource = &v
	return s
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) SetWorkspaceId(v int64) *ListModelInstanceResponseBodyModelInstanceInfoList {
	s.WorkspaceId = &v
	return s
}

func (s *ListModelInstanceResponseBodyModelInstanceInfoList) Validate() error {
	return dara.Validate(s)
}
