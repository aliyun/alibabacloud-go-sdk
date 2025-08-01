// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreserveHeaderFormatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *PreserveHeaderFormatRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *PreserveHeaderFormatRequest
	GetGatewayUniqueId() *string
	SetPreserveHeaderFormat(v bool) *PreserveHeaderFormatRequest
	GetPreserveHeaderFormat() *bool
}

type PreserveHeaderFormatRequest struct {
	// The language in which you want to display the results. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-3f97e2989c344f35ab3fd62b19f1****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Specifies whether the request header is case-sensitive. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	PreserveHeaderFormat *bool `json:"PreserveHeaderFormat,omitempty" xml:"PreserveHeaderFormat,omitempty"`
}

func (s PreserveHeaderFormatRequest) String() string {
	return dara.Prettify(s)
}

func (s PreserveHeaderFormatRequest) GoString() string {
	return s.String()
}

func (s *PreserveHeaderFormatRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *PreserveHeaderFormatRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *PreserveHeaderFormatRequest) GetPreserveHeaderFormat() *bool {
	return s.PreserveHeaderFormat
}

func (s *PreserveHeaderFormatRequest) SetAcceptLanguage(v string) *PreserveHeaderFormatRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *PreserveHeaderFormatRequest) SetGatewayUniqueId(v string) *PreserveHeaderFormatRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *PreserveHeaderFormatRequest) SetPreserveHeaderFormat(v bool) *PreserveHeaderFormatRequest {
	s.PreserveHeaderFormat = &v
	return s
}

func (s *PreserveHeaderFormatRequest) Validate() error {
	return dara.Validate(s)
}
