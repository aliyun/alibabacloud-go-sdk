// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJvmConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetJvmConfigurationRequest
	GetAppId() *string
	SetGroupId(v string) *GetJvmConfigurationRequest
	GetGroupId() *string
}

type GetJvmConfigurationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-4***************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the instance group.
	//
	// 	- If an ID is specified, this operation queries the JVM configuration information of the instance group.
	//
	// 	- If an ID is not specified, this operation queries the JVM configuration information of the application.
	//
	// example:
	//
	// 8123db90-880f-48**************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GetJvmConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJvmConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetJvmConfigurationRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetJvmConfigurationRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetJvmConfigurationRequest) SetAppId(v string) *GetJvmConfigurationRequest {
	s.AppId = &v
	return s
}

func (s *GetJvmConfigurationRequest) SetGroupId(v string) *GetJvmConfigurationRequest {
	s.GroupId = &v
	return s
}

func (s *GetJvmConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
