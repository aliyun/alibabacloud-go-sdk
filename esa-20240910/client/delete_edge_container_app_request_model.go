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
	// The application ID. You can call the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation to obtain the application ID.
	//
	// <notice>AppId is required. If this parameter is not specified, the API returns InvalidParameter.appid (400). You can obtain the value by calling ListEdgeContainerApps.
	//
	// The AppId value is in the format of the app- prefix followed by 18 or more digits (at least 20 characters in total). You can obtain the actual value from the AppId field in the ListEdgeContainerApps response.</notice>.
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
