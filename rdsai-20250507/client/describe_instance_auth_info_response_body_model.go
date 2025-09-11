// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAuthInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeys(v *DescribeInstanceAuthInfoResponseBodyApiKeys) *DescribeInstanceAuthInfoResponseBody
	GetApiKeys() *DescribeInstanceAuthInfoResponseBodyApiKeys
	SetConfigList(v []*DescribeInstanceAuthInfoResponseBodyConfigList) *DescribeInstanceAuthInfoResponseBody
	GetConfigList() []*DescribeInstanceAuthInfoResponseBodyConfigList
	SetInstanceName(v string) *DescribeInstanceAuthInfoResponseBody
	GetInstanceName() *string
	SetJwtSecret(v string) *DescribeInstanceAuthInfoResponseBody
	GetJwtSecret() *string
	SetRequestId(v string) *DescribeInstanceAuthInfoResponseBody
	GetRequestId() *string
}

type DescribeInstanceAuthInfoResponseBody struct {
	// API Keys
	ApiKeys      *DescribeInstanceAuthInfoResponseBodyApiKeys      `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Struct"`
	ConfigList   []*DescribeInstanceAuthInfoResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	InstanceName *string                                           `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// i5o1XAp4sR*****oyOb3O
	JwtSecret *string `json:"JwtSecret,omitempty" xml:"JwtSecret,omitempty"`
	// example:
	//
	// 87249A6F-xxx-804C-E1E0AD1FAD90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceAuthInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAuthInfoResponseBody) GetApiKeys() *DescribeInstanceAuthInfoResponseBodyApiKeys {
	return s.ApiKeys
}

func (s *DescribeInstanceAuthInfoResponseBody) GetConfigList() []*DescribeInstanceAuthInfoResponseBodyConfigList {
	return s.ConfigList
}

func (s *DescribeInstanceAuthInfoResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceAuthInfoResponseBody) GetJwtSecret() *string {
	return s.JwtSecret
}

func (s *DescribeInstanceAuthInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAuthInfoResponseBody) SetApiKeys(v *DescribeInstanceAuthInfoResponseBodyApiKeys) *DescribeInstanceAuthInfoResponseBody {
	s.ApiKeys = v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBody) SetConfigList(v []*DescribeInstanceAuthInfoResponseBodyConfigList) *DescribeInstanceAuthInfoResponseBody {
	s.ConfigList = v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBody) SetInstanceName(v string) *DescribeInstanceAuthInfoResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBody) SetJwtSecret(v string) *DescribeInstanceAuthInfoResponseBody {
	s.JwtSecret = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBody) SetRequestId(v string) *DescribeInstanceAuthInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAuthInfoResponseBodyApiKeys struct {
	// Supabase ANON_KEY
	//
	// example:
	//
	// eyxxxJ9.ey****
	AnonKey *string `json:"AnonKey,omitempty" xml:"AnonKey,omitempty"`
	// Supabase SERVICE_ROLE_KEY
	//
	// example:
	//
	// eyxxxJ9.ey****KfQ.DaYxxxt4Q
	ServiceKey *string `json:"ServiceKey,omitempty" xml:"ServiceKey,omitempty"`
}

func (s DescribeInstanceAuthInfoResponseBodyApiKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAuthInfoResponseBodyApiKeys) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAuthInfoResponseBodyApiKeys) GetAnonKey() *string {
	return s.AnonKey
}

func (s *DescribeInstanceAuthInfoResponseBodyApiKeys) GetServiceKey() *string {
	return s.ServiceKey
}

func (s *DescribeInstanceAuthInfoResponseBodyApiKeys) SetAnonKey(v string) *DescribeInstanceAuthInfoResponseBodyApiKeys {
	s.AnonKey = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBodyApiKeys) SetServiceKey(v string) *DescribeInstanceAuthInfoResponseBodyApiKeys {
	s.ServiceKey = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBodyApiKeys) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAuthInfoResponseBodyConfigList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceAuthInfoResponseBodyConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAuthInfoResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAuthInfoResponseBodyConfigList) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceAuthInfoResponseBodyConfigList) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceAuthInfoResponseBodyConfigList) SetName(v string) *DescribeInstanceAuthInfoResponseBodyConfigList {
	s.Name = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBodyConfigList) SetValue(v string) *DescribeInstanceAuthInfoResponseBodyConfigList {
	s.Value = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBodyConfigList) Validate() error {
	return dara.Validate(s)
}
