// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeContainerAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteEdgeContainerAppRequest
	GetAppId() *string
}

type DeleteEdgeContainerAppRequest struct {
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// example:
	//
	// app-1232321454***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteEdgeContainerAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeContainerAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteEdgeContainerAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteEdgeContainerAppRequest) SetAppId(v string) *DeleteEdgeContainerAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteEdgeContainerAppRequest) Validate() error {
	return dara.Validate(s)
}
