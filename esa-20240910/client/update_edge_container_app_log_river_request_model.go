// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeContainerAppLogRiverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateEdgeContainerAppLogRiverRequest
	GetAppId() *string
	SetPath(v string) *UpdateEdgeContainerAppLogRiverRequest
	GetPath() *string
	SetStdout(v bool) *UpdateEdgeContainerAppLogRiverRequest
	GetStdout() *bool
}

type UpdateEdgeContainerAppLogRiverRequest struct {
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](https://help.aliyun.com/document_detail/2852396.html) operation.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The log path of the container.
	//
	// example:
	//
	// /root/hello.log
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Specifies whether to collect the standard output of the container.
	//
	// example:
	//
	// true
	Stdout *bool `json:"Stdout,omitempty" xml:"Stdout,omitempty"`
}

func (s UpdateEdgeContainerAppLogRiverRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppLogRiverRequest) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppLogRiverRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateEdgeContainerAppLogRiverRequest) GetPath() *string {
	return s.Path
}

func (s *UpdateEdgeContainerAppLogRiverRequest) GetStdout() *bool {
	return s.Stdout
}

func (s *UpdateEdgeContainerAppLogRiverRequest) SetAppId(v string) *UpdateEdgeContainerAppLogRiverRequest {
	s.AppId = &v
	return s
}

func (s *UpdateEdgeContainerAppLogRiverRequest) SetPath(v string) *UpdateEdgeContainerAppLogRiverRequest {
	s.Path = &v
	return s
}

func (s *UpdateEdgeContainerAppLogRiverRequest) SetStdout(v bool) *UpdateEdgeContainerAppLogRiverRequest {
	s.Stdout = &v
	return s
}

func (s *UpdateEdgeContainerAppLogRiverRequest) Validate() error {
	return dara.Validate(s)
}
