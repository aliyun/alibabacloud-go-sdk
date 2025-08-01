// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteServiceSourceRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *DeleteServiceSourceRequest
	GetGatewayUniqueId() *string
	SetSourceId(v int64) *DeleteServiceSourceRequest
	GetSourceId() *int64
}

type DeleteServiceSourceRequest struct {
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
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-492af9b04bb4474cae9d645be850****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// 17
	SourceId *int64 `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s DeleteServiceSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceSourceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteServiceSourceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteServiceSourceRequest) GetSourceId() *int64 {
	return s.SourceId
}

func (s *DeleteServiceSourceRequest) SetAcceptLanguage(v string) *DeleteServiceSourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteServiceSourceRequest) SetGatewayUniqueId(v string) *DeleteServiceSourceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteServiceSourceRequest) SetSourceId(v int64) *DeleteServiceSourceRequest {
	s.SourceId = &v
	return s
}

func (s *DeleteServiceSourceRequest) Validate() error {
	return dara.Validate(s)
}
