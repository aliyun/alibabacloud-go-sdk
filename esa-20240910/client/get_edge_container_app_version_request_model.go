// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetEdgeContainerAppVersionRequest
	GetAppId() *string
	SetVersionId(v string) *GetEdgeContainerAppVersionRequest
	GetVersionId() *string
}

type GetEdgeContainerAppVersionRequest struct {
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The version ID, which can be obtained by calling the [ListEdgeContainerAppVersions](~~ListEdgeContainerAppVersions~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ver-87962637161651****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetEdgeContainerAppVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppVersionRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppVersionRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerAppVersionRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetEdgeContainerAppVersionRequest) SetAppId(v string) *GetEdgeContainerAppVersionRequest {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerAppVersionRequest) SetVersionId(v string) *GetEdgeContainerAppVersionRequest {
	s.VersionId = &v
	return s
}

func (s *GetEdgeContainerAppVersionRequest) Validate() error {
	return dara.Validate(s)
}
