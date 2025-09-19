// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartHoneypotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHoneypotId(v string) *StartHoneypotRequest
	GetHoneypotId() *string
	SetLang(v string) *StartHoneypotRequest
	GetLang() *string
}

type StartHoneypotRequest struct {
	// The ID of the honeypot.
	//
	// >  You can call the [ListHoneypot](~~ListHoneypot~~) operation to query the IDs of honeypots.
	//
	// This parameter is required.
	//
	// example:
	//
	// dba7d44775be8e0e5888ee3b1a62554a93d2512247cabc38ddeac17a3b3f****
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
	// The language of the content in the request and response messages. Default value: **zh**. Valid values:
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

func (s StartHoneypotRequest) String() string {
	return dara.Prettify(s)
}

func (s StartHoneypotRequest) GoString() string {
	return s.String()
}

func (s *StartHoneypotRequest) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *StartHoneypotRequest) GetLang() *string {
	return s.Lang
}

func (s *StartHoneypotRequest) SetHoneypotId(v string) *StartHoneypotRequest {
	s.HoneypotId = &v
	return s
}

func (s *StartHoneypotRequest) SetLang(v string) *StartHoneypotRequest {
	s.Lang = &v
	return s
}

func (s *StartHoneypotRequest) Validate() error {
	return dara.Validate(s)
}
