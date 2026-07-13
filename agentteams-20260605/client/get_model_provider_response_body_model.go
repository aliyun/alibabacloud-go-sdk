// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetModelProviderResponseBody
	GetCode() *string
	SetData(v *GetModelProviderResponseBodyData) *GetModelProviderResponseBody
	GetData() *GetModelProviderResponseBodyData
	SetMessage(v string) *GetModelProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetModelProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetModelProviderResponseBody
	GetSuccess() *bool
}

type GetModelProviderResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetModelProviderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetModelProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetModelProviderResponseBody) GetData() *GetModelProviderResponseBodyData {
	return s.Data
}

func (s *GetModelProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetModelProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetModelProviderResponseBody) SetCode(v string) *GetModelProviderResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelProviderResponseBody) SetData(v *GetModelProviderResponseBodyData) *GetModelProviderResponseBody {
	s.Data = v
	return s
}

func (s *GetModelProviderResponseBody) SetMessage(v string) *GetModelProviderResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelProviderResponseBody) SetRequestId(v string) *GetModelProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelProviderResponseBody) SetSuccess(v bool) *GetModelProviderResponseBody {
	s.Success = &v
	return s
}

func (s *GetModelProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelProviderResponseBodyData struct {
	Address      *string   `json:"Address,omitempty" xml:"Address,omitempty"`
	ApiKeys      []*string `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	CreateTime   *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeployStatus *string   `json:"DeployStatus,omitempty" xml:"DeployStatus,omitempty"`
	Description  *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Id           *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocols    []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	Provider     *string   `json:"Provider,omitempty" xml:"Provider,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetModelProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetModelProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModelProviderResponseBodyData) GetAddress() *string {
	return s.Address
}

func (s *GetModelProviderResponseBodyData) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *GetModelProviderResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetModelProviderResponseBodyData) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *GetModelProviderResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetModelProviderResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetModelProviderResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetModelProviderResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetModelProviderResponseBodyData) GetProtocols() []*string {
	return s.Protocols
}

func (s *GetModelProviderResponseBodyData) GetProvider() *string {
	return s.Provider
}

func (s *GetModelProviderResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetModelProviderResponseBodyData) SetAddress(v string) *GetModelProviderResponseBodyData {
	s.Address = &v
	return s
}

func (s *GetModelProviderResponseBodyData) SetApiKeys(v []*string) *GetModelProviderResponseBodyData {
	s.ApiKeys = v
	return s
}

func (s *GetModelProviderResponseBodyData) SetCreateTime(v string) *GetModelProviderResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetModelProviderResponseBodyData) SetDeployStatus(v string) *GetModelProviderResponseBodyData {
	s.DeployStatus = &v
	return s
}

func (s *GetModelProviderResponseBodyData) SetDescription(v string) *GetModelProviderResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetModelProviderResponseBodyData) SetId(v string) *GetModelProviderResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetModelProviderResponseBodyData) SetInstanceId(v string) *GetModelProviderResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetModelProviderResponseBodyData) SetName(v string) *GetModelProviderResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetModelProviderResponseBodyData) SetProtocols(v []*string) *GetModelProviderResponseBodyData {
	s.Protocols = v
	return s
}

func (s *GetModelProviderResponseBodyData) SetProvider(v string) *GetModelProviderResponseBodyData {
	s.Provider = &v
	return s
}

func (s *GetModelProviderResponseBodyData) SetRegionId(v string) *GetModelProviderResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetModelProviderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
