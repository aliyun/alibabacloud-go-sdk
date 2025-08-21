// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RollbackApplicationRequest
	GetAppId() *string
	SetFromAppVersion(v string) *RollbackApplicationRequest
	GetFromAppVersion() *string
	SetTimeout(v int32) *RollbackApplicationRequest
	GetTimeout() *int32
	SetToAppVersion(v string) *RollbackApplicationRequest
	GetToAppVersion() *string
}

type RollbackApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 474bdef0-d149-4695-abfb-52912d91****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The current version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// v2-1
	FromAppVersion *string `json:"FromAppVersion,omitempty" xml:"FromAppVersion,omitempty"`
	// The timeout period of the asynchronous rollback operation. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 1800
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The target version number. By default, the system automatically rolls back the container version to the previous version.
	//
	// example:
	//
	// v2
	ToAppVersion *string `json:"ToAppVersion,omitempty" xml:"ToAppVersion,omitempty"`
}

func (s RollbackApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackApplicationRequest) GoString() string {
	return s.String()
}

func (s *RollbackApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *RollbackApplicationRequest) GetFromAppVersion() *string {
	return s.FromAppVersion
}

func (s *RollbackApplicationRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *RollbackApplicationRequest) GetToAppVersion() *string {
	return s.ToAppVersion
}

func (s *RollbackApplicationRequest) SetAppId(v string) *RollbackApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RollbackApplicationRequest) SetFromAppVersion(v string) *RollbackApplicationRequest {
	s.FromAppVersion = &v
	return s
}

func (s *RollbackApplicationRequest) SetTimeout(v int32) *RollbackApplicationRequest {
	s.Timeout = &v
	return s
}

func (s *RollbackApplicationRequest) SetToAppVersion(v string) *RollbackApplicationRequest {
	s.ToAppVersion = &v
	return s
}

func (s *RollbackApplicationRequest) Validate() error {
	return dara.Validate(s)
}
