// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteAuthResourceRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *DeleteAuthResourceRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *DeleteAuthResourceRequest
	GetId() *int64
}

type DeleteAuthResourceRequest struct {
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
	// gw-3f97e2989c344f35ab3fd62b19f1****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The destination ID.
	//
	// example:
	//
	// 36
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteAuthResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteAuthResourceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteAuthResourceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteAuthResourceRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteAuthResourceRequest) SetAcceptLanguage(v string) *DeleteAuthResourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteAuthResourceRequest) SetGatewayUniqueId(v string) *DeleteAuthResourceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteAuthResourceRequest) SetId(v int64) *DeleteAuthResourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteAuthResourceRequest) Validate() error {
	return dara.Validate(s)
}
