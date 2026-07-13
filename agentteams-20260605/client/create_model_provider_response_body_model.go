// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateModelProviderResponseBody
	GetCode() *string
	SetData(v *CreateModelProviderResponseBodyData) *CreateModelProviderResponseBody
	GetData() *CreateModelProviderResponseBodyData
	SetMessage(v string) *CreateModelProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateModelProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateModelProviderResponseBody
	GetSuccess() *bool
}

type CreateModelProviderResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateModelProviderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateModelProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateModelProviderResponseBody) GetData() *CreateModelProviderResponseBodyData {
	return s.Data
}

func (s *CreateModelProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateModelProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateModelProviderResponseBody) SetCode(v string) *CreateModelProviderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateModelProviderResponseBody) SetData(v *CreateModelProviderResponseBodyData) *CreateModelProviderResponseBody {
	s.Data = v
	return s
}

func (s *CreateModelProviderResponseBody) SetMessage(v string) *CreateModelProviderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateModelProviderResponseBody) SetRequestId(v string) *CreateModelProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelProviderResponseBody) SetSuccess(v bool) *CreateModelProviderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateModelProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateModelProviderResponseBodyData struct {
	Address     *string   `json:"Address,omitempty" xml:"Address,omitempty"`
	CreateTime  *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocols   []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	Provider    *string   `json:"Provider,omitempty" xml:"Provider,omitempty"`
}

func (s CreateModelProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateModelProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateModelProviderResponseBodyData) GetAddress() *string {
	return s.Address
}

func (s *CreateModelProviderResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateModelProviderResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateModelProviderResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CreateModelProviderResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateModelProviderResponseBodyData) GetProtocols() []*string {
	return s.Protocols
}

func (s *CreateModelProviderResponseBodyData) GetProvider() *string {
	return s.Provider
}

func (s *CreateModelProviderResponseBodyData) SetAddress(v string) *CreateModelProviderResponseBodyData {
	s.Address = &v
	return s
}

func (s *CreateModelProviderResponseBodyData) SetCreateTime(v string) *CreateModelProviderResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateModelProviderResponseBodyData) SetDescription(v string) *CreateModelProviderResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateModelProviderResponseBodyData) SetId(v string) *CreateModelProviderResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateModelProviderResponseBodyData) SetName(v string) *CreateModelProviderResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateModelProviderResponseBodyData) SetProtocols(v []*string) *CreateModelProviderResponseBodyData {
	s.Protocols = v
	return s
}

func (s *CreateModelProviderResponseBodyData) SetProvider(v string) *CreateModelProviderResponseBodyData {
	s.Provider = &v
	return s
}

func (s *CreateModelProviderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
