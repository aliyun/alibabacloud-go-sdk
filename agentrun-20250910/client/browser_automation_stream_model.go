// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowserAutomationStream interface {
	dara.Model
	String() string
	GoString() string
	SetStreamEndpoint(v string) *BrowserAutomationStream
	GetStreamEndpoint() *string
	SetStreamStatus(v string) *BrowserAutomationStream
	GetStreamStatus() *string
}

type BrowserAutomationStream struct {
	StreamEndpoint *string `json:"streamEndpoint,omitempty" xml:"streamEndpoint,omitempty"`
	StreamStatus   *string `json:"streamStatus,omitempty" xml:"streamStatus,omitempty"`
}

func (s BrowserAutomationStream) String() string {
	return dara.Prettify(s)
}

func (s BrowserAutomationStream) GoString() string {
	return s.String()
}

func (s *BrowserAutomationStream) GetStreamEndpoint() *string {
	return s.StreamEndpoint
}

func (s *BrowserAutomationStream) GetStreamStatus() *string {
	return s.StreamStatus
}

func (s *BrowserAutomationStream) SetStreamEndpoint(v string) *BrowserAutomationStream {
	s.StreamEndpoint = &v
	return s
}

func (s *BrowserAutomationStream) SetStreamStatus(v string) *BrowserAutomationStream {
	s.StreamStatus = &v
	return s
}

func (s *BrowserAutomationStream) Validate() error {
	return dara.Validate(s)
}
