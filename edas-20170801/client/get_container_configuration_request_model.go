// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContainerConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetContainerConfigurationRequest
	GetAppId() *string
	SetGroupId(v string) *GetContainerConfigurationRequest
	GetGroupId() *string
}

type GetContainerConfigurationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-**************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the instance group.
	//
	// 	- If this parameter is specified, this operation queries the Tomcat configuration of the instance group.
	//
	// 	- If this parameter is not specified, this operation queries the Tomcat configuration of the application.
	//
	// example:
	//
	// 8123db90-880f-**************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GetContainerConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContainerConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetContainerConfigurationRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetContainerConfigurationRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetContainerConfigurationRequest) SetAppId(v string) *GetContainerConfigurationRequest {
	s.AppId = &v
	return s
}

func (s *GetContainerConfigurationRequest) SetGroupId(v string) *GetContainerConfigurationRequest {
	s.GroupId = &v
	return s
}

func (s *GetContainerConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
