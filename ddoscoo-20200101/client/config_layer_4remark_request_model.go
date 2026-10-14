// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v string) *ConfigLayer4RemarkRequest
	GetListeners() *string
}

type ConfigLayer4RemarkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{\\"InstanceId\\":\\"ddosDip-sg-4hr2b3l****\\",\\"FrontendPort\\":2020,\\"Protocol\\":\\"udp\\",\\"Remark\\":\\"test\\"}]
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s ConfigLayer4RemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RemarkRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RemarkRequest) GetListeners() *string {
	return s.Listeners
}

func (s *ConfigLayer4RemarkRequest) SetListeners(v string) *ConfigLayer4RemarkRequest {
	s.Listeners = &v
	return s
}

func (s *ConfigLayer4RemarkRequest) Validate() error {
	return dara.Validate(s)
}
