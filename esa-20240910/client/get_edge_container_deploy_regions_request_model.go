// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerDeployRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetEdgeContainerDeployRegionsRequest
	GetAppId() *string
}

type GetEdgeContainerDeployRegionsRequest struct {
	// The application ID. You can call the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation to obtain the application ID.	Notice: This parameter is required. If this parameter is not specified, the service returns InvalidParameter.appid (400). The valid format is app-{18-digit number}. You can call the ListEdgeContainerApps operation to obtain the application ID. Example: app-880****75783794688. If you have not activated the Edge Container service, activate it first and then call the CreateEdgeContainerApp operation to create an application and obtain the AppId.</notice>.
	//
	// example:
	//
	// GetEdgeContainerDeployRegions
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetEdgeContainerDeployRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerDeployRegionsRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerDeployRegionsRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerDeployRegionsRequest) SetAppId(v string) *GetEdgeContainerDeployRegionsRequest {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerDeployRegionsRequest) Validate() error {
	return dara.Validate(s)
}
