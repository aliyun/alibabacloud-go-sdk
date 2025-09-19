// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InstallCloudMonitorResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *InstallCloudMonitorResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *InstallCloudMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallCloudMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InstallCloudMonitorResponseBody
	GetSuccess() *bool
}

type InstallCloudMonitorResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// IllegalParam
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// There was an error with your request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F92AFB96-FACC-57E7-928E-678D04B94CAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InstallCloudMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *InstallCloudMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *InstallCloudMonitorResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *InstallCloudMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallCloudMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallCloudMonitorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InstallCloudMonitorResponseBody) SetCode(v string) *InstallCloudMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *InstallCloudMonitorResponseBody) SetHttpStatusCode(v int32) *InstallCloudMonitorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InstallCloudMonitorResponseBody) SetMessage(v string) *InstallCloudMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *InstallCloudMonitorResponseBody) SetRequestId(v string) *InstallCloudMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallCloudMonitorResponseBody) SetSuccess(v bool) *InstallCloudMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *InstallCloudMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
