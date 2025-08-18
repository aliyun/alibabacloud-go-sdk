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
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The release environment. Valid values: prod and staging.
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
