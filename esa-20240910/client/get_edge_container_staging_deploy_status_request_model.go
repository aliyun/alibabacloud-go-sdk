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
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
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
