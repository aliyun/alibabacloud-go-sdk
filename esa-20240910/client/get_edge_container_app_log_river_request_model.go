// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppLogRiverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetEdgeContainerAppLogRiverRequest
	GetAppId() *string
}

type GetEdgeContainerAppLogRiverRequest struct {
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// example:
	//
	// app-880688675****88
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetEdgeContainerAppLogRiverRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppLogRiverRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppLogRiverRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerAppLogRiverRequest) SetAppId(v string) *GetEdgeContainerAppLogRiverRequest {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerAppLogRiverRequest) Validate() error {
	return dara.Validate(s)
}
