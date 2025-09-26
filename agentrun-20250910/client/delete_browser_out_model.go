// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBrowserOut interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserId(v string) *DeleteBrowserOut
	GetBrowserId() *string
	SetBrowserName(v string) *DeleteBrowserOut
	GetBrowserName() *string
	SetStatus(v string) *DeleteBrowserOut
	GetStatus() *string
}

type DeleteBrowserOut struct {
	BrowserId   *string `json:"browserId,omitempty" xml:"browserId,omitempty"`
	BrowserName *string `json:"browserName,omitempty" xml:"browserName,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteBrowserOut) String() string {
	return dara.Prettify(s)
}

func (s DeleteBrowserOut) GoString() string {
	return s.String()
}

func (s *DeleteBrowserOut) GetBrowserId() *string {
	return s.BrowserId
}

func (s *DeleteBrowserOut) GetBrowserName() *string {
	return s.BrowserName
}

func (s *DeleteBrowserOut) GetStatus() *string {
	return s.Status
}

func (s *DeleteBrowserOut) SetBrowserId(v string) *DeleteBrowserOut {
	s.BrowserId = &v
	return s
}

func (s *DeleteBrowserOut) SetBrowserName(v string) *DeleteBrowserOut {
	s.BrowserName = &v
	return s
}

func (s *DeleteBrowserOut) SetStatus(v string) *DeleteBrowserOut {
	s.Status = &v
	return s
}

func (s *DeleteBrowserOut) Validate() error {
	return dara.Validate(s)
}
