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
