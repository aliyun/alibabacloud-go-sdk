// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowserLiveViewStream interface {
	dara.Model
	String() string
	GoString() string
	SetStreamEndpoint(v string) *BrowserLiveViewStream
	GetStreamEndpoint() *string
}

type BrowserLiveViewStream struct {
	// example:
	//
	// wss://browser-liveview.cn-hangzhou.agentrun.aliyuncs.com/stream/bs-1234567890abcdef
	StreamEndpoint *string `json:"streamEndpoint,omitempty" xml:"streamEndpoint,omitempty"`
}

func (s BrowserLiveViewStream) String() string {
	return dara.Prettify(s)
}

func (s BrowserLiveViewStream) GoString() string {
	return s.String()
}

func (s *BrowserLiveViewStream) GetStreamEndpoint() *string {
	return s.StreamEndpoint
}

func (s *BrowserLiveViewStream) SetStreamEndpoint(v string) *BrowserLiveViewStream {
	s.StreamEndpoint = &v
	return s
}

func (s *BrowserLiveViewStream) Validate() error {
	return dara.Validate(s)
}
