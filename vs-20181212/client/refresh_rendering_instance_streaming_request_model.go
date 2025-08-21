// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshRenderingInstanceStreamingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientInfo(v *RefreshRenderingInstanceStreamingRequestClientInfo) *RefreshRenderingInstanceStreamingRequest
	GetClientInfo() *RefreshRenderingInstanceStreamingRequestClientInfo
	SetRenderingInstanceId(v string) *RefreshRenderingInstanceStreamingRequest
	GetRenderingInstanceId() *string
}

type RefreshRenderingInstanceStreamingRequest struct {
	ClientInfo *RefreshRenderingInstanceStreamingRequestClientInfo `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s RefreshRenderingInstanceStreamingRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshRenderingInstanceStreamingRequest) GoString() string {
	return s.String()
}

func (s *RefreshRenderingInstanceStreamingRequest) GetClientInfo() *RefreshRenderingInstanceStreamingRequestClientInfo {
	return s.ClientInfo
}

func (s *RefreshRenderingInstanceStreamingRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *RefreshRenderingInstanceStreamingRequest) SetClientInfo(v *RefreshRenderingInstanceStreamingRequestClientInfo) *RefreshRenderingInstanceStreamingRequest {
	s.ClientInfo = v
	return s
}

func (s *RefreshRenderingInstanceStreamingRequest) SetRenderingInstanceId(v string) *RefreshRenderingInstanceStreamingRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *RefreshRenderingInstanceStreamingRequest) Validate() error {
	return dara.Validate(s)
}

type RefreshRenderingInstanceStreamingRequestClientInfo struct {
	// example:
	//
	// 172.21.128.110
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// true
	NewClient *bool `json:"NewClient,omitempty" xml:"NewClient,omitempty"`
}

func (s RefreshRenderingInstanceStreamingRequestClientInfo) String() string {
	return dara.Prettify(s)
}

func (s RefreshRenderingInstanceStreamingRequestClientInfo) GoString() string {
	return s.String()
}

func (s *RefreshRenderingInstanceStreamingRequestClientInfo) GetClientIp() *string {
	return s.ClientIp
}

func (s *RefreshRenderingInstanceStreamingRequestClientInfo) GetNewClient() *bool {
	return s.NewClient
}

func (s *RefreshRenderingInstanceStreamingRequestClientInfo) SetClientIp(v string) *RefreshRenderingInstanceStreamingRequestClientInfo {
	s.ClientIp = &v
	return s
}

func (s *RefreshRenderingInstanceStreamingRequestClientInfo) SetNewClient(v bool) *RefreshRenderingInstanceStreamingRequestClientInfo {
	s.NewClient = &v
	return s
}

func (s *RefreshRenderingInstanceStreamingRequestClientInfo) Validate() error {
	return dara.Validate(s)
}
