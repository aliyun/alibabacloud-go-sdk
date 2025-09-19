// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopHoneypotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHoneypotId(v string) *StopHoneypotRequest
	GetHoneypotId() *string
	SetLang(v string) *StopHoneypotRequest
	GetLang() *string
}

type StopHoneypotRequest struct {
	// The honeypot ID.
	//
	// >  You can call the [ListHoneypot](~~ListHoneypot~~) operation to obtain IDs of honeypots.
	//
	// This parameter is required.
	//
	// example:
	//
	// 444c699ac151b183b04b562b1dc02639d504c9d097246a322de75c963fe*****
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s StopHoneypotRequest) String() string {
	return dara.Prettify(s)
}

func (s StopHoneypotRequest) GoString() string {
	return s.String()
}

func (s *StopHoneypotRequest) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *StopHoneypotRequest) GetLang() *string {
	return s.Lang
}

func (s *StopHoneypotRequest) SetHoneypotId(v string) *StopHoneypotRequest {
	s.HoneypotId = &v
	return s
}

func (s *StopHoneypotRequest) SetLang(v string) *StopHoneypotRequest {
	s.Lang = &v
	return s
}

func (s *StopHoneypotRequest) Validate() error {
	return dara.Validate(s)
}
