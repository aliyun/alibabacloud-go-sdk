// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWebTerminalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerName(v string) *GetWebTerminalRequest
	GetContainerName() *string
	SetIsShared(v bool) *GetWebTerminalRequest
	GetIsShared() *bool
	SetPodUid(v string) *GetWebTerminalRequest
	GetPodUid() *string
}

type GetWebTerminalRequest struct {
	// The name of the container. If this parameter is not specified, the WebTerminal of the primary container is accessed.
	//
	// example:
	//
	// sidecar-test
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// Specifies whether to create a shareable link for accessing the container. Valid values:
	//
	// - true: Returns a shareable link for accessing the container. The link expires after 30 seconds and can be used only once. After the container is accessed through this link, subsequent requests to access the container by using the same link become invalid.
	//
	// - false: Returns a regular link for accessing the container. Alibaba Cloud identity verification is required when you use the link. The link expires after 30 seconds.
	//
	// example:
	//
	// true
	IsShared *bool `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The Pod UID.
	//
	// example:
	//
	// 94a7cc7c-0033-48b5-85bd-71c63592****
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
}

func (s GetWebTerminalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWebTerminalRequest) GoString() string {
	return s.String()
}

func (s *GetWebTerminalRequest) GetContainerName() *string {
	return s.ContainerName
}

func (s *GetWebTerminalRequest) GetIsShared() *bool {
	return s.IsShared
}

func (s *GetWebTerminalRequest) GetPodUid() *string {
	return s.PodUid
}

func (s *GetWebTerminalRequest) SetContainerName(v string) *GetWebTerminalRequest {
	s.ContainerName = &v
	return s
}

func (s *GetWebTerminalRequest) SetIsShared(v bool) *GetWebTerminalRequest {
	s.IsShared = &v
	return s
}

func (s *GetWebTerminalRequest) SetPodUid(v string) *GetWebTerminalRequest {
	s.PodUid = &v
	return s
}

func (s *GetWebTerminalRequest) Validate() error {
	return dara.Validate(s)
}
