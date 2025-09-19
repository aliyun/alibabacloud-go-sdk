// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHoneypotId(v string) *DeleteHoneypotRequest
	GetHoneypotId() *string
	SetLang(v string) *DeleteHoneypotRequest
	GetLang() *string
}

type DeleteHoneypotRequest struct {
	// The ID of the honeypot.
	//
	// > You can call the [ListHoneypot](~~ListHoneypot~~) operation to query the IDs of honeypots.
	//
	// This parameter is required.
	//
	// example:
	//
	// 558b5fa40948ebe2171a74757c54dc7e58f761870fa7ee6a105e70947ec82aa9
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

func (s DeleteHoneypotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotRequest) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotRequest) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *DeleteHoneypotRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteHoneypotRequest) SetHoneypotId(v string) *DeleteHoneypotRequest {
	s.HoneypotId = &v
	return s
}

func (s *DeleteHoneypotRequest) SetLang(v string) *DeleteHoneypotRequest {
	s.Lang = &v
	return s
}

func (s *DeleteHoneypotRequest) Validate() error {
	return dara.Validate(s)
}
