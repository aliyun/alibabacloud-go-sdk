// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerStagingDeployStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetEdgeContainerStagingDeployStatusRequest
	GetAppId() *string
}

type GetEdgeContainerStagingDeployStatusRequest struct {
	// The application ID. You can call [ListEdgeContainerApps](~~ListEdgeContainerApps~~) to obtain the application ID.
	//
	// 	Notice: AppId is required. If AppId is not specified, the API returns InvalidParameter.appid (400). You can call ListEdgeContainerApps to obtain the application ID.
	//
	// The AppId format is the prefix app- followed by 18 or more digits (at least 20 characters in total). You can obtain the actual value from the AppId field in the ListEdgeContainerApps response.
	//
	// Complete call chain example: CreateEdgeContainerApp → Call ListEdgeContainerApps to obtain AppId → GetEdgeContainerStagingDeployStatus</notice>.
	//
	// example:
	//
	// GetEdgeContainerStagingDeployStatus
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetEdgeContainerStagingDeployStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerStagingDeployStatusRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerStagingDeployStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerStagingDeployStatusRequest) SetAppId(v string) *GetEdgeContainerStagingDeployStatusRequest {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusRequest) Validate() error {
	return dara.Validate(s)
}
