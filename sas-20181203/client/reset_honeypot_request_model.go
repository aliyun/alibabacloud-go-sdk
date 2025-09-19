// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetHoneypotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHoneypotId(v string) *ResetHoneypotRequest
	GetHoneypotId() *string
	SetLang(v string) *ResetHoneypotRequest
	GetLang() *string
}

type ResetHoneypotRequest struct {
	// The ID of the honeypot.
	//
	// >  You can call the [ListHoneypot](~~ListHoneypot~~) operation to query the IDs of honeypots.
	//
	// This parameter is required.
	//
	// example:
	//
	// 945607c2ae2a1a737c04599d6608065688bfc6048d9b9d306ce8dc8191c*****
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

func (s ResetHoneypotRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetHoneypotRequest) GoString() string {
	return s.String()
}

func (s *ResetHoneypotRequest) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *ResetHoneypotRequest) GetLang() *string {
	return s.Lang
}

func (s *ResetHoneypotRequest) SetHoneypotId(v string) *ResetHoneypotRequest {
	s.HoneypotId = &v
	return s
}

func (s *ResetHoneypotRequest) SetLang(v string) *ResetHoneypotRequest {
	s.Lang = &v
	return s
}

func (s *ResetHoneypotRequest) Validate() error {
	return dara.Validate(s)
}
