// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWebTerminalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsShared(v bool) *GetWebTerminalRequest
	GetIsShared() *bool
	SetPodUid(v string) *GetWebTerminalRequest
	GetPodUid() *string
}

type GetWebTerminalRequest struct {
	// Specifies whether to create a shareable link to access the container. Valid values:
	//
	// 	- true: returns a shareable link to access the container. The link will expire after 30 seconds and can only be used once. After you access the container by using the link, other requests that use this link to access the container become invalid.
	//
	// 	- false: returns a common shareable link to access the container. If you use a common shareable link to access a container, Alibaba Cloud identity authentication is required. The link will expire after 30 seconds.
	//
	// example:
	//
	// true
	IsShared *bool `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The pod UID.
	//
	// example:
	//
	// 94a7cc7c-0033-48b5-85bd-71c63592c268
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
}

func (s GetWebTerminalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWebTerminalRequest) GoString() string {
	return s.String()
}

func (s *GetWebTerminalRequest) GetIsShared() *bool {
	return s.IsShared
}

func (s *GetWebTerminalRequest) GetPodUid() *string {
	return s.PodUid
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
