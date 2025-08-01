// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersByIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListListenersByIpRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *ListListenersByIpRequest
	GetInstanceId() *string
	SetIp(v string) *ListListenersByIpRequest
	GetIp() *string
	SetNamespaceId(v string) *ListListenersByIpRequest
	GetNamespaceId() *string
	SetRequestPars(v string) *ListListenersByIpRequest
	GetRequestPars() *string
}

type ListListenersByIpRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-i7m2cecji09
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.22.0.240
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// ea02a522-9482-4255-bb65-dc0636d783f2
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ListListenersByIpRequest) String() string {
	return dara.Prettify(s)
}

func (s ListListenersByIpRequest) GoString() string {
	return s.String()
}

func (s *ListListenersByIpRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListListenersByIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListListenersByIpRequest) GetIp() *string {
	return s.Ip
}

func (s *ListListenersByIpRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListListenersByIpRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListListenersByIpRequest) SetAcceptLanguage(v string) *ListListenersByIpRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListListenersByIpRequest) SetInstanceId(v string) *ListListenersByIpRequest {
	s.InstanceId = &v
	return s
}

func (s *ListListenersByIpRequest) SetIp(v string) *ListListenersByIpRequest {
	s.Ip = &v
	return s
}

func (s *ListListenersByIpRequest) SetNamespaceId(v string) *ListListenersByIpRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListListenersByIpRequest) SetRequestPars(v string) *ListListenersByIpRequest {
	s.RequestPars = &v
	return s
}

func (s *ListListenersByIpRequest) Validate() error {
	return dara.Validate(s)
}
