// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppResourceReserveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetEdgeContainerAppResourceReserveRequest
	GetAppId() *string
}

type GetEdgeContainerAppResourceReserveRequest struct {
	// The application ID. You can call the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation to obtain the application ID.
	//
	// 	Notice: This parameter is required. If this parameter is not specified, the service returns InvalidParameter.Appid (400).
	//
	// If no applications exist under your account, call CreateEdgeContainerApp to create an application first, and then call ListEdgeContainerApps to obtain the AppId.
	//
	// </notice>.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetEdgeContainerAppResourceReserveRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceReserveRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceReserveRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerAppResourceReserveRequest) SetAppId(v string) *GetEdgeContainerAppResourceReserveRequest {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveRequest) Validate() error {
	return dara.Validate(s)
}
