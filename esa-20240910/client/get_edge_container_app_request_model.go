// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetEdgeContainerAppRequest
	GetAppId() *string
}

type GetEdgeContainerAppRequest struct {
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetEdgeContainerAppRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerAppRequest) SetAppId(v string) *GetEdgeContainerAppRequest {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerAppRequest) Validate() error {
	return dara.Validate(s)
}
