// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetImageRequest
	GetAcceptLanguage() *string
	SetVersionCode(v string) *GetImageRequest
	GetVersionCode() *string
}

type GetImageRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The version number of the current instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZooKeeper_3_5_5
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s GetImageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetImageRequest) GetVersionCode() *string {
	return s.VersionCode
}

func (s *GetImageRequest) SetAcceptLanguage(v string) *GetImageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetImageRequest) SetVersionCode(v string) *GetImageRequest {
	s.VersionCode = &v
	return s
}

func (s *GetImageRequest) Validate() error {
	return dara.Validate(s)
}
