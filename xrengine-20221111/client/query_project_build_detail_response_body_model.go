// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProjectBuildDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryProjectBuildDetailResponseBody
	GetCode() *string
	SetData(v *QueryProjectBuildDetailResponseBodyData) *QueryProjectBuildDetailResponseBody
	GetData() *QueryProjectBuildDetailResponseBodyData
	SetMessage(v string) *QueryProjectBuildDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryProjectBuildDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryProjectBuildDetailResponseBody
	GetSuccess() *bool
}

type QueryProjectBuildDetailResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryProjectBuildDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryProjectBuildDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectBuildDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProjectBuildDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryProjectBuildDetailResponseBody) GetData() *QueryProjectBuildDetailResponseBodyData {
	return s.Data
}

func (s *QueryProjectBuildDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryProjectBuildDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryProjectBuildDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryProjectBuildDetailResponseBody) SetCode(v string) *QueryProjectBuildDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBody) SetData(v *QueryProjectBuildDetailResponseBodyData) *QueryProjectBuildDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryProjectBuildDetailResponseBody) SetMessage(v string) *QueryProjectBuildDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBody) SetRequestId(v string) *QueryProjectBuildDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBody) SetSuccess(v bool) *QueryProjectBuildDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryProjectBuildDetailResponseBodyData struct {
	GmtCreate                  *string                                                              `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified                *string                                                              `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                         *int64                                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceDetailResponseList []*QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList `json:"InstanceDetailResponseList,omitempty" xml:"InstanceDetailResponseList,omitempty" type:"Repeated"`
	Status                     *string                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Title                      *string                                                              `json:"Title,omitempty" xml:"Title,omitempty"`
	UserId                     *int64                                                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryProjectBuildDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectBuildDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProjectBuildDetailResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryProjectBuildDetailResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryProjectBuildDetailResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QueryProjectBuildDetailResponseBodyData) GetInstanceDetailResponseList() []*QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList {
	return s.InstanceDetailResponseList
}

func (s *QueryProjectBuildDetailResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryProjectBuildDetailResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *QueryProjectBuildDetailResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryProjectBuildDetailResponseBodyData) SetGmtCreate(v string) *QueryProjectBuildDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyData) SetGmtModified(v string) *QueryProjectBuildDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyData) SetId(v int64) *QueryProjectBuildDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyData) SetInstanceDetailResponseList(v []*QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) *QueryProjectBuildDetailResponseBodyData {
	s.InstanceDetailResponseList = v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyData) SetStatus(v string) *QueryProjectBuildDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyData) SetTitle(v string) *QueryProjectBuildDetailResponseBodyData {
	s.Title = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyData) SetUserId(v int64) *QueryProjectBuildDetailResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyData) Validate() error {
	if s.InstanceDetailResponseList != nil {
		for _, item := range s.InstanceDetailResponseList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList struct {
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Params      *string `json:"Params,omitempty" xml:"Params,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmitTime  *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	TemplateId  *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) GoString() string {
	return s.String()
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) GetId() *int64 {
	return s.Id
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) GetParams() *string {
	return s.Params
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) GetStatus() *string {
	return s.Status
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) SetGmtCreate(v string) *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList {
	s.GmtCreate = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) SetGmtModified(v string) *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList {
	s.GmtModified = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) SetId(v int64) *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList {
	s.Id = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) SetParams(v string) *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList {
	s.Params = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) SetStatus(v string) *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList {
	s.Status = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) SetSubmitTime(v string) *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList {
	s.SubmitTime = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) SetTemplateId(v int64) *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList {
	s.TemplateId = &v
	return s
}

func (s *QueryProjectBuildDetailResponseBodyDataInstanceDetailResponseList) Validate() error {
	return dara.Validate(s)
}
