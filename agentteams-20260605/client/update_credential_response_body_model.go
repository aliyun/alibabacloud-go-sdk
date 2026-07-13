// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCredentialResponseBody
	GetCode() *string
	SetData(v *UpdateCredentialResponseBodyData) *UpdateCredentialResponseBody
	GetData() *UpdateCredentialResponseBodyData
	SetHttpStatusCode(v int32) *UpdateCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCredentialResponseBody
	GetSuccess() *bool
}

type UpdateCredentialResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCredentialResponseBody) GetData() *UpdateCredentialResponseBodyData {
	return s.Data
}

func (s *UpdateCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCredentialResponseBody) SetCode(v string) *UpdateCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCredentialResponseBody) SetData(v *UpdateCredentialResponseBodyData) *UpdateCredentialResponseBody {
	s.Data = v
	return s
}

func (s *UpdateCredentialResponseBody) SetHttpStatusCode(v int32) *UpdateCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCredentialResponseBody) SetMessage(v string) *UpdateCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCredentialResponseBody) SetRequestId(v string) *UpdateCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCredentialResponseBody) SetSuccess(v bool) *UpdateCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCredentialResponseBodyData struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateCredentialResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateCredentialResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateCredentialResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCredentialResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateCredentialResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateCredentialResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateCredentialResponseBodyData) SetCreateTime(v string) *UpdateCredentialResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetDescription(v string) *UpdateCredentialResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetInstanceId(v string) *UpdateCredentialResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetName(v string) *UpdateCredentialResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetStatus(v string) *UpdateCredentialResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetUpdateTime(v string) *UpdateCredentialResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
