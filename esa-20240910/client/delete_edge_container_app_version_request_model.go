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
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-96253477062511****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the version that you want to delete. To obtain the version ID, call the [ListEdgeContainerAppVersions](~~ListEdgeContainerAppVersions~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ver-89884764010378****
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
