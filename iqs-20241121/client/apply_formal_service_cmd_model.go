// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyFormalServiceCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ApplyFormalServiceCmd
	GetAccountId() *string
	SetAccountName(v string) *ApplyFormalServiceCmd
	GetAccountName() *string
	SetAddress(v string) *ApplyFormalServiceCmd
	GetAddress() *string
	SetApplyServiceInfos(v []map[string]interface{}) *ApplyFormalServiceCmd
	GetApplyServiceInfos() []map[string]interface{}
	SetApplyType(v string) *ApplyFormalServiceCmd
	GetApplyType() *string
	SetContactName(v string) *ApplyFormalServiceCmd
	GetContactName() *string
	SetInstanceId(v string) *ApplyFormalServiceCmd
	GetInstanceId() *string
	SetPhone(v string) *ApplyFormalServiceCmd
	GetPhone() *string
	SetProductName(v string) *ApplyFormalServiceCmd
	GetProductName() *string
	SetQps(v int32) *ApplyFormalServiceCmd
	GetQps() *int32
	SetSceneDesc(v string) *ApplyFormalServiceCmd
	GetSceneDesc() *string
	SetServiceType(v string) *ApplyFormalServiceCmd
	GetServiceType() *string
}

type ApplyFormalServiceCmd struct {
	AccountId         *string                  `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AccountName       *string                  `json:"accountName,omitempty" xml:"accountName,omitempty"`
	Address           *string                  `json:"address,omitempty" xml:"address,omitempty"`
	ApplyServiceInfos []map[string]interface{} `json:"applyServiceInfos,omitempty" xml:"applyServiceInfos,omitempty" type:"Repeated"`
	ApplyType         *string                  `json:"applyType,omitempty" xml:"applyType,omitempty"`
	ContactName       *string                  `json:"contactName,omitempty" xml:"contactName,omitempty"`
	InstanceId        *string                  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Phone             *string                  `json:"phone,omitempty" xml:"phone,omitempty"`
	ProductName       *string                  `json:"productName,omitempty" xml:"productName,omitempty"`
	Qps               *int32                   `json:"qps,omitempty" xml:"qps,omitempty"`
	SceneDesc         *string                  `json:"sceneDesc,omitempty" xml:"sceneDesc,omitempty"`
	ServiceType       *string                  `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ApplyFormalServiceCmd) String() string {
	return dara.Prettify(s)
}

func (s ApplyFormalServiceCmd) GoString() string {
	return s.String()
}

func (s *ApplyFormalServiceCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *ApplyFormalServiceCmd) GetAccountName() *string {
	return s.AccountName
}

func (s *ApplyFormalServiceCmd) GetAddress() *string {
	return s.Address
}

func (s *ApplyFormalServiceCmd) GetApplyServiceInfos() []map[string]interface{} {
	return s.ApplyServiceInfos
}

func (s *ApplyFormalServiceCmd) GetApplyType() *string {
	return s.ApplyType
}

func (s *ApplyFormalServiceCmd) GetContactName() *string {
	return s.ContactName
}

func (s *ApplyFormalServiceCmd) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ApplyFormalServiceCmd) GetPhone() *string {
	return s.Phone
}

func (s *ApplyFormalServiceCmd) GetProductName() *string {
	return s.ProductName
}

func (s *ApplyFormalServiceCmd) GetQps() *int32 {
	return s.Qps
}

func (s *ApplyFormalServiceCmd) GetSceneDesc() *string {
	return s.SceneDesc
}

func (s *ApplyFormalServiceCmd) GetServiceType() *string {
	return s.ServiceType
}

func (s *ApplyFormalServiceCmd) SetAccountId(v string) *ApplyFormalServiceCmd {
	s.AccountId = &v
	return s
}

func (s *ApplyFormalServiceCmd) SetAccountName(v string) *ApplyFormalServiceCmd {
	s.AccountName = &v
	return s
}

func (s *ApplyFormalServiceCmd) SetAddress(v string) *ApplyFormalServiceCmd {
	s.Address = &v
	return s
}

func (s *ApplyFormalServiceCmd) SetApplyServiceInfos(v []map[string]interface{}) *ApplyFormalServiceCmd {
	s.ApplyServiceInfos = v
	return s
}

func (s *ApplyFormalServiceCmd) SetApplyType(v string) *ApplyFormalServiceCmd {
	s.ApplyType = &v
	return s
}

func (s *ApplyFormalServiceCmd) SetContactName(v string) *ApplyFormalServiceCmd {
	s.ContactName = &v
	return s
}

func (s *ApplyFormalServiceCmd) SetInstanceId(v string) *ApplyFormalServiceCmd {
	s.InstanceId = &v
	return s
}

func (s *ApplyFormalServiceCmd) SetPhone(v string) *ApplyFormalServiceCmd {
	s.Phone = &v
	return s
}

func (s *ApplyFormalServiceCmd) SetProductName(v string) *ApplyFormalServiceCmd {
	s.ProductName = &v
	return s
}

func (s *ApplyFormalServiceCmd) SetQps(v int32) *ApplyFormalServiceCmd {
	s.Qps = &v
	return s
}

func (s *ApplyFormalServiceCmd) SetSceneDesc(v string) *ApplyFormalServiceCmd {
	s.SceneDesc = &v
	return s
}

func (s *ApplyFormalServiceCmd) SetServiceType(v string) *ApplyFormalServiceCmd {
	s.ServiceType = &v
	return s
}

func (s *ApplyFormalServiceCmd) Validate() error {
	return dara.Validate(s)
}
