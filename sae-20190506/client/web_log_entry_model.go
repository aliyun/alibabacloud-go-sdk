// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebLogEntry interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *WebLogEntry
	GetMessage() *string
}

type WebLogEntry struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s WebLogEntry) String() string {
	return dara.Prettify(s)
}

func (s WebLogEntry) GoString() string {
	return s.String()
}

func (s *WebLogEntry) GetMessage() *string {
	return s.Message
}

func (s *WebLogEntry) SetMessage(v string) *WebLogEntry {
	s.Message = &v
	return s
}

func (s *WebLogEntry) Validate() error {
	return dara.Validate(s)
}
