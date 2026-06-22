// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageRegistryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ModifyImageRegistryRequest
	GetDomainName() *string
	SetId(v int64) *ModifyImageRegistryRequest
	GetId() *int64
	SetPassword(v string) *ModifyImageRegistryRequest
	GetPassword() *string
	SetPort(v int32) *ModifyImageRegistryRequest
	GetPort() *int32
	SetRegistryHostIp(v string) *ModifyImageRegistryRequest
	GetRegistryHostIp() *string
	SetTransPerHour(v int32) *ModifyImageRegistryRequest
	GetTransPerHour() *int32
	SetUserName(v string) *ModifyImageRegistryRequest
	GetUserName() *string
}

type ModifyImageRegistryRequest struct {
	// The domain name.
	//
	// example:
	//
	// 114.55.xxx.xxx
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The image registry ID. Call the [ListImageRegistry](https://help.aliyun.com/document_detail/471986.html) operation to obtain this ID.
	//
	// example:
	//
	// 390103286
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The password.
	//
	// example:
	//
	// ********************
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port number.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The registry IP address.
	//
	// example:
	//
	// 192.168.0.1
	RegistryHostIp *string `json:"RegistryHostIp,omitempty" xml:"RegistryHostIp,omitempty"`
	// The number of images to scan per hour.
	//
	// example:
	//
	// 10
	TransPerHour *int32 `json:"TransPerHour,omitempty" xml:"TransPerHour,omitempty"`
	// The username.
	//
	// example:
	//
	// xxxx
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ModifyImageRegistryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageRegistryRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageRegistryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyImageRegistryRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyImageRegistryRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyImageRegistryRequest) GetPort() *int32 {
	return s.Port
}

func (s *ModifyImageRegistryRequest) GetRegistryHostIp() *string {
	return s.RegistryHostIp
}

func (s *ModifyImageRegistryRequest) GetTransPerHour() *int32 {
	return s.TransPerHour
}

func (s *ModifyImageRegistryRequest) GetUserName() *string {
	return s.UserName
}

func (s *ModifyImageRegistryRequest) SetDomainName(v string) *ModifyImageRegistryRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyImageRegistryRequest) SetId(v int64) *ModifyImageRegistryRequest {
	s.Id = &v
	return s
}

func (s *ModifyImageRegistryRequest) SetPassword(v string) *ModifyImageRegistryRequest {
	s.Password = &v
	return s
}

func (s *ModifyImageRegistryRequest) SetPort(v int32) *ModifyImageRegistryRequest {
	s.Port = &v
	return s
}

func (s *ModifyImageRegistryRequest) SetRegistryHostIp(v string) *ModifyImageRegistryRequest {
	s.RegistryHostIp = &v
	return s
}

func (s *ModifyImageRegistryRequest) SetTransPerHour(v int32) *ModifyImageRegistryRequest {
	s.TransPerHour = &v
	return s
}

func (s *ModifyImageRegistryRequest) SetUserName(v string) *ModifyImageRegistryRequest {
	s.UserName = &v
	return s
}

func (s *ModifyImageRegistryRequest) Validate() error {
	return dara.Validate(s)
}
