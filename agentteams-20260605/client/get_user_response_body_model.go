// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserResponseBody
	GetCode() *string
	SetData(v *GetUserResponseBodyData) *GetUserResponseBody
	GetData() *GetUserResponseBodyData
	SetHttpStatusCode(v int32) *GetUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserResponseBody
	GetSuccess() *bool
}

type GetUserResponseBody struct {
	Code           *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserResponseBody) GetData() *GetUserResponseBodyData {
	return s.Data
}

func (s *GetUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserResponseBody) SetCode(v string) *GetUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResponseBody) SetData(v *GetUserResponseBodyData) *GetUserResponseBody {
	s.Data = v
	return s
}

func (s *GetUserResponseBody) SetHttpStatusCode(v int32) *GetUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetSuccess(v bool) *GetUserResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserResponseBodyData struct {
	AuthMethod      *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedAt       *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Note            *string `json:"Note,omitempty" xml:"Note,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyData) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *GetUserResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUserResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetUserResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *GetUserResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetUserResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *GetUserResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUserResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetUserResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetUserResponseBodyData) SetAuthMethod(v string) *GetUserResponseBodyData {
	s.AuthMethod = &v
	return s
}

func (s *GetUserResponseBodyData) SetCreateTime(v string) *GetUserResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetUserResponseBodyData) SetCreatedAt(v string) *GetUserResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetUserResponseBodyData) SetDisplayName(v string) *GetUserResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBodyData) SetEmail(v string) *GetUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyData) SetInstanceId(v string) *GetUserResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetUserResponseBodyData) SetName(v string) *GetUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetUserResponseBodyData) SetNote(v string) *GetUserResponseBodyData {
	s.Note = &v
	return s
}

func (s *GetUserResponseBodyData) SetRegionId(v string) *GetUserResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetUserResponseBodyData) SetResourceGroupId(v string) *GetUserResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetUserResponseBodyData) SetStatus(v string) *GetUserResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
