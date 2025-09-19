// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneyPotUploadPolicyInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetHoneyPotUploadPolicyInfoRequest
	GetLang() *string
}

type GetHoneyPotUploadPolicyInfoRequest struct {
	// The language of the content within the request and the response. Valid values:
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

func (s GetHoneyPotUploadPolicyInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHoneyPotUploadPolicyInfoRequest) GoString() string {
	return s.String()
}

func (s *GetHoneyPotUploadPolicyInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *GetHoneyPotUploadPolicyInfoRequest) SetLang(v string) *GetHoneyPotUploadPolicyInfoRequest {
	s.Lang = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoRequest) Validate() error {
	return dara.Validate(s)
}
