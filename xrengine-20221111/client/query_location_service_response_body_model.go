// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocationServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryLocationServiceResponseBody
	GetCode() *string
	SetData(v *QueryLocationServiceResponseBodyData) *QueryLocationServiceResponseBody
	GetData() *QueryLocationServiceResponseBodyData
	SetErrorName(v string) *QueryLocationServiceResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *QueryLocationServiceResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *QueryLocationServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryLocationServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryLocationServiceResponseBody
	GetSuccess() *bool
}

type QueryLocationServiceResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryLocationServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                               `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                                `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryLocationServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLocationServiceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLocationServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryLocationServiceResponseBody) GetData() *QueryLocationServiceResponseBodyData {
	return s.Data
}

func (s *QueryLocationServiceResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *QueryLocationServiceResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *QueryLocationServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryLocationServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLocationServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryLocationServiceResponseBody) SetCode(v string) *QueryLocationServiceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryLocationServiceResponseBody) SetData(v *QueryLocationServiceResponseBodyData) *QueryLocationServiceResponseBody {
	s.Data = v
	return s
}

func (s *QueryLocationServiceResponseBody) SetErrorName(v string) *QueryLocationServiceResponseBody {
	s.ErrorName = &v
	return s
}

func (s *QueryLocationServiceResponseBody) SetHttpCode(v int32) *QueryLocationServiceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryLocationServiceResponseBody) SetMessage(v string) *QueryLocationServiceResponseBody {
	s.Message = &v
	return s
}

func (s *QueryLocationServiceResponseBody) SetRequestId(v string) *QueryLocationServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLocationServiceResponseBody) SetSuccess(v bool) *QueryLocationServiceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryLocationServiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryLocationServiceResponseBodyData struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ExpireTime      *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId         *int64  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName       *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Note            *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Qps             *int64  `json:"Qps,omitempty" xml:"Qps,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SvcDeployStatus *string `json:"SvcDeployStatus,omitempty" xml:"SvcDeployStatus,omitempty"`
	SvcState        *string `json:"SvcState,omitempty" xml:"SvcState,omitempty"`
	TreeConfig      *string `json:"TreeConfig,omitempty" xml:"TreeConfig,omitempty"`
	UuidLogMap      *string `json:"UuidLogMap,omitempty" xml:"UuidLogMap,omitempty"`
}

func (s QueryLocationServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryLocationServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryLocationServiceResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *QueryLocationServiceResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *QueryLocationServiceResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryLocationServiceResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryLocationServiceResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QueryLocationServiceResponseBodyData) GetImageId() *int64 {
	return s.ImageId
}

func (s *QueryLocationServiceResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *QueryLocationServiceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryLocationServiceResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *QueryLocationServiceResponseBodyData) GetQps() *int64 {
	return s.Qps
}

func (s *QueryLocationServiceResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryLocationServiceResponseBodyData) GetSvcDeployStatus() *string {
	return s.SvcDeployStatus
}

func (s *QueryLocationServiceResponseBodyData) GetSvcState() *string {
	return s.SvcState
}

func (s *QueryLocationServiceResponseBodyData) GetTreeConfig() *string {
	return s.TreeConfig
}

func (s *QueryLocationServiceResponseBodyData) GetUuidLogMap() *string {
	return s.UuidLogMap
}

func (s *QueryLocationServiceResponseBodyData) SetAppId(v string) *QueryLocationServiceResponseBodyData {
	s.AppId = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetExpireTime(v string) *QueryLocationServiceResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetGmtCreate(v string) *QueryLocationServiceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetGmtModified(v string) *QueryLocationServiceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetId(v int64) *QueryLocationServiceResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetImageId(v int64) *QueryLocationServiceResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetImageName(v string) *QueryLocationServiceResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetName(v string) *QueryLocationServiceResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetNote(v string) *QueryLocationServiceResponseBodyData {
	s.Note = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetQps(v int64) *QueryLocationServiceResponseBodyData {
	s.Qps = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetStartTime(v string) *QueryLocationServiceResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetSvcDeployStatus(v string) *QueryLocationServiceResponseBodyData {
	s.SvcDeployStatus = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetSvcState(v string) *QueryLocationServiceResponseBodyData {
	s.SvcState = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetTreeConfig(v string) *QueryLocationServiceResponseBodyData {
	s.TreeConfig = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) SetUuidLogMap(v string) *QueryLocationServiceResponseBodyData {
	s.UuidLogMap = &v
	return s
}

func (s *QueryLocationServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
