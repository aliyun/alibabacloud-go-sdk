// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCredentialResponseBody
	GetCode() *string
	SetData(v *CreateCredentialResponseBodyData) *CreateCredentialResponseBody
	GetData() *CreateCredentialResponseBodyData
	SetHttpStatusCode(v int32) *CreateCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCredentialResponseBody
	GetSuccess() *bool
}

type CreateCredentialResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCredentialResponseBody) GetData() *CreateCredentialResponseBodyData {
	return s.Data
}

func (s *CreateCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCredentialResponseBody) SetCode(v string) *CreateCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCredentialResponseBody) SetData(v *CreateCredentialResponseBodyData) *CreateCredentialResponseBody {
	s.Data = v
	return s
}

func (s *CreateCredentialResponseBody) SetHttpStatusCode(v int32) *CreateCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCredentialResponseBody) SetMessage(v string) *CreateCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCredentialResponseBody) SetRequestId(v string) *CreateCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCredentialResponseBody) SetSuccess(v bool) *CreateCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCredentialResponseBodyData struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCredentialResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateCredentialResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateCredentialResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCredentialResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateCredentialResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateCredentialResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateCredentialResponseBodyData) SetCreateTime(v string) *CreateCredentialResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetDescription(v string) *CreateCredentialResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetInstanceId(v string) *CreateCredentialResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetName(v string) *CreateCredentialResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetStatus(v string) *CreateCredentialResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetUpdateTime(v string) *CreateCredentialResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *CreateCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
