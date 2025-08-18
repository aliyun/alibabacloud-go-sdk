// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerTerminalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetEdgeContainerTerminalRequest
	GetAppId() *string
}

type GetEdgeContainerTerminalRequest struct {
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetEdgeContainerTerminalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerTerminalRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerTerminalRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerTerminalRequest) SetAppId(v string) *GetEdgeContainerTerminalRequest {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerTerminalRequest) Validate() error {
	return dara.Validate(s)
}
