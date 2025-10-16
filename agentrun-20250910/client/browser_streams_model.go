// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowserStreams interface {
	dara.Model
	String() string
	GoString() string
	SetAutomationStream(v *BrowserAutomationStream) *BrowserStreams
	GetAutomationStream() *BrowserAutomationStream
	SetLiveViewStream(v *BrowserLiveViewStream) *BrowserStreams
	GetLiveViewStream() *BrowserLiveViewStream
}

type BrowserStreams struct {
	AutomationStream *BrowserAutomationStream `json:"automationStream,omitempty" xml:"automationStream,omitempty"`
	LiveViewStream   *BrowserLiveViewStream   `json:"liveViewStream,omitempty" xml:"liveViewStream,omitempty"`
}

func (s BrowserStreams) String() string {
	return dara.Prettify(s)
}

func (s BrowserStreams) GoString() string {
	return s.String()
}

func (s *BrowserStreams) GetAutomationStream() *BrowserAutomationStream {
	return s.AutomationStream
}

func (s *BrowserStreams) GetLiveViewStream() *BrowserLiveViewStream {
	return s.LiveViewStream
}

func (s *BrowserStreams) SetAutomationStream(v *BrowserAutomationStream) *BrowserStreams {
	s.AutomationStream = v
	return s
}

func (s *BrowserStreams) SetLiveViewStream(v *BrowserLiveViewStream) *BrowserStreams {
	s.LiveViewStream = v
	return s
}

func (s *BrowserStreams) Validate() error {
	if s.AutomationStream != nil {
		if err := s.AutomationStream.Validate(); err != nil {
			return err
		}
	}
	if s.LiveViewStream != nil {
		if err := s.LiveViewStream.Validate(); err != nil {
			return err
		}
	}
	return nil
}
