// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeContainerAppVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteEdgeContainerAppVersionRequest
	GetAppId() *string
	SetVersionId(v string) *DeleteEdgeContainerAppVersionRequest
	GetVersionId() *string
}

type DeleteEdgeContainerAppVersionRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeleteEdgeContainerAppVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeContainerAppVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteEdgeContainerAppVersionRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteEdgeContainerAppVersionRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *DeleteEdgeContainerAppVersionRequest) SetAppId(v string) *DeleteEdgeContainerAppVersionRequest {
	s.AppId = &v
	return s
}

func (s *DeleteEdgeContainerAppVersionRequest) SetVersionId(v string) *DeleteEdgeContainerAppVersionRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteEdgeContainerAppVersionRequest) Validate() error {
	return dara.Validate(s)
}
