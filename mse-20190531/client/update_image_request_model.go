// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateImageRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *UpdateImageRequest
	GetClusterId() *string
	SetVersionCode(v string) *UpdateImageRequest
	GetVersionCode() *string
}

type UpdateImageRequest struct {
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
	// The ID of the destination cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-8e8e9060
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The destination version number.
	//
	// > You must call the GetImage operation to obtain the maximum destination version number that corresponds to MaxVersionCode.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZooKeeper_3_5_5
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s UpdateImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateImageRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateImageRequest) GetVersionCode() *string {
	return s.VersionCode
}

func (s *UpdateImageRequest) SetAcceptLanguage(v string) *UpdateImageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateImageRequest) SetClusterId(v string) *UpdateImageRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateImageRequest) SetVersionCode(v string) *UpdateImageRequest {
	s.VersionCode = &v
	return s
}

func (s *UpdateImageRequest) Validate() error {
	return dara.Validate(s)
}
