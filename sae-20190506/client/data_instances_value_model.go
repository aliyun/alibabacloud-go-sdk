// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataInstancesValue interface {
	dara.Model
	String() string
	GoString() string
	SetDnsName(v string) *DataInstancesValue
	GetDnsName() *string
	SetListeners(v map[string]*DataInstancesValueListenersValue) *DataInstancesValue
	GetListeners() map[string]*DataInstancesValueListenersValue
	SetCreatedBySae(v bool) *DataInstancesValue
	GetCreatedBySae() *bool
}

type DataInstancesValue struct {
	// The domain name.
	//
	// example:
	//
	// nlb-wb7r6dlwetvt5j****.cn-hangzhou.nlb.aliyuncs.com
	DnsName *string `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	// The listeners.
	Listeners map[string]*DataInstancesValueListenersValue `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	// Indicates whether the instance is created by SAE.
	//
	// 	- **true**: The instance is created by SAE.
	//
	// 	- **false**: The existing instance is reused.
	//
	// example:
	//
	// true
	CreatedBySae *bool `json:"CreatedBySae,omitempty" xml:"CreatedBySae,omitempty"`
}

func (s DataInstancesValue) String() string {
	return dara.Prettify(s)
}

func (s DataInstancesValue) GoString() string {
	return s.String()
}

func (s *DataInstancesValue) GetDnsName() *string {
	return s.DnsName
}

func (s *DataInstancesValue) GetListeners() map[string]*DataInstancesValueListenersValue {
	return s.Listeners
}

func (s *DataInstancesValue) GetCreatedBySae() *bool {
	return s.CreatedBySae
}

func (s *DataInstancesValue) SetDnsName(v string) *DataInstancesValue {
	s.DnsName = &v
	return s
}

func (s *DataInstancesValue) SetListeners(v map[string]*DataInstancesValueListenersValue) *DataInstancesValue {
	s.Listeners = v
	return s
}

func (s *DataInstancesValue) SetCreatedBySae(v bool) *DataInstancesValue {
	s.CreatedBySae = &v
	return s
}

func (s *DataInstancesValue) Validate() error {
	return dara.Validate(s)
}
