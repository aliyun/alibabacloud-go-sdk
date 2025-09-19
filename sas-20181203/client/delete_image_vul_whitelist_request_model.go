// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageVulWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *DeleteImageVulWhitelistRequest
	GetIds() *string
	SetLang(v string) *DeleteImageVulWhitelistRequest
	GetLang() *string
}

type DeleteImageVulWhitelistRequest struct {
	// The ID of the whitelist. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 123
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
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

func (s DeleteImageVulWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageVulWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageVulWhitelistRequest) GetIds() *string {
	return s.Ids
}

func (s *DeleteImageVulWhitelistRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteImageVulWhitelistRequest) SetIds(v string) *DeleteImageVulWhitelistRequest {
	s.Ids = &v
	return s
}

func (s *DeleteImageVulWhitelistRequest) SetLang(v string) *DeleteImageVulWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *DeleteImageVulWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
