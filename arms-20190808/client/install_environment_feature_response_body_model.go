// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallEnvironmentFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InstallEnvironmentFeatureResponseBody
	GetCode() *int32
	SetData(v string) *InstallEnvironmentFeatureResponseBody
	GetData() *string
	SetMessage(v string) *InstallEnvironmentFeatureResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallEnvironmentFeatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InstallEnvironmentFeatureResponseBody
	GetSuccess() *bool
}

type InstallEnvironmentFeatureResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The release ID.
	//
	// example:
	//
	// 83FCC44C-A056-18AF-A902-7043E723F0D9
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4C518054-852F-4023-ABC1-4AF95FF7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InstallEnvironmentFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallEnvironmentFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *InstallEnvironmentFeatureResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InstallEnvironmentFeatureResponseBody) GetData() *string {
	return s.Data
}

func (s *InstallEnvironmentFeatureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallEnvironmentFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallEnvironmentFeatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InstallEnvironmentFeatureResponseBody) SetCode(v int32) *InstallEnvironmentFeatureResponseBody {
	s.Code = &v
	return s
}

func (s *InstallEnvironmentFeatureResponseBody) SetData(v string) *InstallEnvironmentFeatureResponseBody {
	s.Data = &v
	return s
}

func (s *InstallEnvironmentFeatureResponseBody) SetMessage(v string) *InstallEnvironmentFeatureResponseBody {
	s.Message = &v
	return s
}

func (s *InstallEnvironmentFeatureResponseBody) SetRequestId(v string) *InstallEnvironmentFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallEnvironmentFeatureResponseBody) SetSuccess(v bool) *InstallEnvironmentFeatureResponseBody {
	s.Success = &v
	return s
}

func (s *InstallEnvironmentFeatureResponseBody) Validate() error {
	return dara.Validate(s)
}
