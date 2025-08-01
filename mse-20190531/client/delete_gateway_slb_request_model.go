// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewaySlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteGatewaySlbRequest
	GetAcceptLanguage() *string
	SetDeleteSlb(v bool) *DeleteGatewaySlbRequest
	GetDeleteSlb() *bool
	SetGatewayUniqueId(v string) *DeleteGatewaySlbRequest
	GetGatewayUniqueId() *string
	SetId(v string) *DeleteGatewaySlbRequest
	GetId() *string
	SetSlbId(v string) *DeleteGatewaySlbRequest
	GetSlbId() *string
}

type DeleteGatewaySlbRequest struct {
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
	// Specifies whether to delete the SLB instance purchased for the gateway when you delete the gateway.
	//
	// example:
	//
	// true
	DeleteSlb *bool `json:"DeleteSlb,omitempty" xml:"DeleteSlb,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-c9bc5afd61014165bd58f621b491****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the association record.
	//
	// example:
	//
	// 395
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the SLB instance that needs to be deleted.
	//
	// example:
	//
	// lb-uf6duug6s13x4abc8****
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
}

func (s DeleteGatewaySlbRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewaySlbRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySlbRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteGatewaySlbRequest) GetDeleteSlb() *bool {
	return s.DeleteSlb
}

func (s *DeleteGatewaySlbRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewaySlbRequest) GetId() *string {
	return s.Id
}

func (s *DeleteGatewaySlbRequest) GetSlbId() *string {
	return s.SlbId
}

func (s *DeleteGatewaySlbRequest) SetAcceptLanguage(v string) *DeleteGatewaySlbRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteGatewaySlbRequest) SetDeleteSlb(v bool) *DeleteGatewaySlbRequest {
	s.DeleteSlb = &v
	return s
}

func (s *DeleteGatewaySlbRequest) SetGatewayUniqueId(v string) *DeleteGatewaySlbRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewaySlbRequest) SetId(v string) *DeleteGatewaySlbRequest {
	s.Id = &v
	return s
}

func (s *DeleteGatewaySlbRequest) SetSlbId(v string) *DeleteGatewaySlbRequest {
	s.SlbId = &v
	return s
}

func (s *DeleteGatewaySlbRequest) Validate() error {
	return dara.Validate(s)
}
