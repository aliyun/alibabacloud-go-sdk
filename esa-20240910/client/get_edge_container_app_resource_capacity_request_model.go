// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppResourceCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetEdgeContainerAppResourceCapacityRequest
	GetAppId() *string
}

type GetEdgeContainerAppResourceCapacityRequest struct {
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetEdgeContainerAppResourceCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceCapacityRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceCapacityRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerAppResourceCapacityRequest) SetAppId(v string) *GetEdgeContainerAppResourceCapacityRequest {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerAppResourceCapacityRequest) Validate() error {
	return dara.Validate(s)
}
