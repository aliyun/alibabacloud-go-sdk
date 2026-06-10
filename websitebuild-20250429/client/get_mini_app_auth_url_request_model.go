// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiniAppAuthUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetMiniAppAuthUrlRequest
	GetBizId() *string
	SetChannel(v string) *GetMiniAppAuthUrlRequest
	GetChannel() *string
	SetRedirectUri(v string) *GetMiniAppAuthUrlRequest
	GetRedirectUri() *string
}

type GetMiniAppAuthUrlRequest struct {
	// Site ID
	//
	// example:
	//
	// WS20250731233102000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Channel information
	//
	// example:
	//
	// WECHAT
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// Hyperlink URL
	//
	// example:
	//
	// null
	RedirectUri *string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty"`
}

func (s GetMiniAppAuthUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppAuthUrlRequest) GoString() string {
	return s.String()
}

func (s *GetMiniAppAuthUrlRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetMiniAppAuthUrlRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetMiniAppAuthUrlRequest) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *GetMiniAppAuthUrlRequest) SetBizId(v string) *GetMiniAppAuthUrlRequest {
	s.BizId = &v
	return s
}

func (s *GetMiniAppAuthUrlRequest) SetChannel(v string) *GetMiniAppAuthUrlRequest {
	s.Channel = &v
	return s
}

func (s *GetMiniAppAuthUrlRequest) SetRedirectUri(v string) *GetMiniAppAuthUrlRequest {
	s.RedirectUri = &v
	return s
}

func (s *GetMiniAppAuthUrlRequest) Validate() error {
	return dara.Validate(s)
}
