// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpLog interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *PdpLog
	GetContent() *string
	SetIp(v string) *PdpLog
	GetIp() *string
}

type PdpLog struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Ip      *string `json:"ip,omitempty" xml:"ip,omitempty"`
}

func (s PdpLog) String() string {
	return dara.Prettify(s)
}

func (s PdpLog) GoString() string {
	return s.String()
}

func (s *PdpLog) GetContent() *string {
	return s.Content
}

func (s *PdpLog) GetIp() *string {
	return s.Ip
}

func (s *PdpLog) SetContent(v string) *PdpLog {
	s.Content = &v
	return s
}

func (s *PdpLog) SetIp(v string) *PdpLog {
	s.Ip = &v
	return s
}

func (s *PdpLog) Validate() error {
	return dara.Validate(s)
}
