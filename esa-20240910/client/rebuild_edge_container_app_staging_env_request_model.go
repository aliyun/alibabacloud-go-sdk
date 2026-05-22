// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildEdgeContainerAppStagingEnvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RebuildEdgeContainerAppStagingEnvRequest
	GetAppId() *string
}

type RebuildEdgeContainerAppStagingEnvRequest struct {
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s RebuildEdgeContainerAppStagingEnvRequest) String() string {
	return dara.Prettify(s)
}

func (s RebuildEdgeContainerAppStagingEnvRequest) GoString() string {
	return s.String()
}

func (s *RebuildEdgeContainerAppStagingEnvRequest) GetAppId() *string {
	return s.AppId
}

func (s *RebuildEdgeContainerAppStagingEnvRequest) SetAppId(v string) *RebuildEdgeContainerAppStagingEnvRequest {
	s.AppId = &v
	return s
}

func (s *RebuildEdgeContainerAppStagingEnvRequest) Validate() error {
	return dara.Validate(s)
}
