// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmInstanceNameRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *UpdateCloudGtmInstanceNameRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateCloudGtmInstanceNameRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateCloudGtmInstanceNameRequest
	GetInstanceName() *string
}

type UpdateCloudGtmInstanceNameRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can specify a custom value for this parameter, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Global Traffic Manager (GTM) instance.
	//
	// example:
	//
	// gtm-cn-jmp3qnw**03
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance. You cannot leave this parameter empty.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s UpdateCloudGtmInstanceNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceNameRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmInstanceNameRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmInstanceNameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloudGtmInstanceNameRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateCloudGtmInstanceNameRequest) SetAcceptLanguage(v string) *UpdateCloudGtmInstanceNameRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmInstanceNameRequest) SetClientToken(v string) *UpdateCloudGtmInstanceNameRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmInstanceNameRequest) SetInstanceId(v string) *UpdateCloudGtmInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloudGtmInstanceNameRequest) SetInstanceName(v string) *UpdateCloudGtmInstanceNameRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateCloudGtmInstanceNameRequest) Validate() error {
	return dara.Validate(s)
}
