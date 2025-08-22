// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDcdnServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillType(v string) *OpenDcdnServiceRequest
	GetBillType() *string
	SetOwnerId(v int64) *OpenDcdnServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *OpenDcdnServiceRequest
	GetSecurityToken() *string
	SetWebsocketBillType(v string) *OpenDcdnServiceRequest
	GetWebsocketBillType() *string
}

type OpenDcdnServiceRequest struct {
	// The metering method of DCDN. Valid values:
	//
	// 	- **PayByTraffic**: pay-by-traffic
	//
	// 	- **PayByBandwidth**: pay-by-bandwidth
	//
	// This parameter is required.
	//
	// example:
	//
	// PayByTraffic
	BillType      *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The metering method of WebSocket. Valid values:
	//
	// 	- **websockettraffic**: pay-by-data-transfer
	//
	// 	- **websocketbps**: pay-by-bandwidth
	//
	// This parameter is required.
	//
	// example:
	//
	// websockettraffic
	WebsocketBillType *string `json:"WebsocketBillType,omitempty" xml:"WebsocketBillType,omitempty"`
}

func (s OpenDcdnServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenDcdnServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenDcdnServiceRequest) GetBillType() *string {
	return s.BillType
}

func (s *OpenDcdnServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenDcdnServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *OpenDcdnServiceRequest) GetWebsocketBillType() *string {
	return s.WebsocketBillType
}

func (s *OpenDcdnServiceRequest) SetBillType(v string) *OpenDcdnServiceRequest {
	s.BillType = &v
	return s
}

func (s *OpenDcdnServiceRequest) SetOwnerId(v int64) *OpenDcdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenDcdnServiceRequest) SetSecurityToken(v string) *OpenDcdnServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *OpenDcdnServiceRequest) SetWebsocketBillType(v string) *OpenDcdnServiceRequest {
	s.WebsocketBillType = &v
	return s
}

func (s *OpenDcdnServiceRequest) Validate() error {
	return dara.Validate(s)
}
