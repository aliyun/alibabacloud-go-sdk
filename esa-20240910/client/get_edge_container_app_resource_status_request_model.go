// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppResourceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetEdgeContainerAppResourceStatusRequest
	GetAppId() *string
}

type GetEdgeContainerAppResourceStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app-96253477062511****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetEdgeContainerAppResourceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerAppResourceStatusRequest) SetAppId(v string) *GetEdgeContainerAppResourceStatusRequest {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerAppResourceStatusRequest) Validate() error {
	return dara.Validate(s)
}
