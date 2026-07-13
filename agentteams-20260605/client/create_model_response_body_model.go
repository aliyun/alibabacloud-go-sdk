// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateModelResponseBody
	GetCode() *string
	SetData(v *CreateModelResponseBodyData) *CreateModelResponseBody
	GetData() *CreateModelResponseBodyData
	SetMessage(v string) *CreateModelResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateModelResponseBody
	GetSuccess() *bool
}

type CreateModelResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateModelResponseBody) GetData() *CreateModelResponseBodyData {
	return s.Data
}

func (s *CreateModelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateModelResponseBody) SetCode(v string) *CreateModelResponseBody {
	s.Code = &v
	return s
}

func (s *CreateModelResponseBody) SetData(v *CreateModelResponseBodyData) *CreateModelResponseBody {
	s.Data = v
	return s
}

func (s *CreateModelResponseBody) SetMessage(v string) *CreateModelResponseBody {
	s.Message = &v
	return s
}

func (s *CreateModelResponseBody) SetRequestId(v string) *CreateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelResponseBody) SetSuccess(v bool) *CreateModelResponseBody {
	s.Success = &v
	return s
}

func (s *CreateModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateModelResponseBodyData struct {
	CreateTime   *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Id           *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocols    []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	Provider     *string   `json:"Provider,omitempty" xml:"Provider,omitempty"`
	ProviderName *string   `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	UpdateTime   *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateModelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateModelResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateModelResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CreateModelResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateModelResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateModelResponseBodyData) GetProtocols() []*string {
	return s.Protocols
}

func (s *CreateModelResponseBodyData) GetProvider() *string {
	return s.Provider
}

func (s *CreateModelResponseBodyData) GetProviderName() *string {
	return s.ProviderName
}

func (s *CreateModelResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *CreateModelResponseBodyData) SetCreateTime(v int64) *CreateModelResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateModelResponseBodyData) SetDescription(v string) *CreateModelResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateModelResponseBodyData) SetId(v string) *CreateModelResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateModelResponseBodyData) SetInstanceId(v string) *CreateModelResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateModelResponseBodyData) SetName(v string) *CreateModelResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateModelResponseBodyData) SetProtocols(v []*string) *CreateModelResponseBodyData {
	s.Protocols = v
	return s
}

func (s *CreateModelResponseBodyData) SetProvider(v string) *CreateModelResponseBodyData {
	s.Provider = &v
	return s
}

func (s *CreateModelResponseBodyData) SetProviderName(v string) *CreateModelResponseBodyData {
	s.ProviderName = &v
	return s
}

func (s *CreateModelResponseBodyData) SetUpdateTime(v int64) *CreateModelResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *CreateModelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
