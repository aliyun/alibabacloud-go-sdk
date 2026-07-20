// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetEdgeContainerAppStatusRequest
	GetAppId() *string
	SetPublishEnv(v string) *GetEdgeContainerAppStatusRequest
	GetPublishEnv() *string
}

type GetEdgeContainerAppStatusRequest struct {
	// The application ID. You can call [ListEdgeContainerApps](~~ListEdgeContainerApps~~) to obtain the application ID. Before calling this operation, you must first activate the edge container service by calling OpenEdgeContainer, and then confirm that an available application exists by calling ListEdgeContainerApps or create an application by calling CreateEdgeContainerApp.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The publishing environment. Valid values: prod and staging.
	//
	// example:
	//
	// staging
	PublishEnv *string `json:"PublishEnv,omitempty" xml:"PublishEnv,omitempty"`
}

func (s GetEdgeContainerAppStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppStatusRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerAppStatusRequest) GetPublishEnv() *string {
	return s.PublishEnv
}

func (s *GetEdgeContainerAppStatusRequest) SetAppId(v string) *GetEdgeContainerAppStatusRequest {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerAppStatusRequest) SetPublishEnv(v string) *GetEdgeContainerAppStatusRequest {
	s.PublishEnv = &v
	return s
}

func (s *GetEdgeContainerAppStatusRequest) Validate() error {
	return dara.Validate(s)
}
