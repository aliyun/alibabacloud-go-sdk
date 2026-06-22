// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWhatsappCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *WhatsappCallRequest
	GetBusinessNumber() *string
	SetCallAction(v string) *WhatsappCallRequest
	GetCallAction() *string
	SetCallId(v string) *WhatsappCallRequest
	GetCallId() *string
	SetCustSpaceId(v string) *WhatsappCallRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *WhatsappCallRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *WhatsappCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *WhatsappCallRequest
	GetResourceOwnerId() *int64
	SetSession(v *WhatsappCallRequestSession) *WhatsappCallRequest
	GetSession() *WhatsappCallRequestSession
	SetUserNumber(v string) *WhatsappCallRequest
	GetUserNumber() *string
}

type WhatsappCallRequest struct {
	// The business number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86138***
	BusinessNumber *string `json:"BusinessNumber,omitempty" xml:"BusinessNumber,omitempty"`
	// The action to perform on the call.
	//
	// This parameter is required.
	//
	// example:
	//
	// connect
	CallAction *string `json:"CallAction,omitempty" xml:"CallAction,omitempty"`
	// The call ID.
	//
	// example:
	//
	// wamid-xx**
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The Space ID of the ISV sub-customer, or the instance ID of the direct customer. View the Space ID on the
	//
	// <props="china">[Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[Channel Management](https://chatapp.console.alibabacloud.com/CustomerList) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// cams-xx**
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The call negotiation information.
	Session *WhatsappCallRequestSession `json:"Session,omitempty" xml:"Session,omitempty" type:"Struct"`
	// The user\\"s number.
	//
	// example:
	//
	// 86131***
	UserNumber *string `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
}

func (s WhatsappCallRequest) String() string {
	return dara.Prettify(s)
}

func (s WhatsappCallRequest) GoString() string {
	return s.String()
}

func (s *WhatsappCallRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *WhatsappCallRequest) GetCallAction() *string {
	return s.CallAction
}

func (s *WhatsappCallRequest) GetCallId() *string {
	return s.CallId
}

func (s *WhatsappCallRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *WhatsappCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *WhatsappCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *WhatsappCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *WhatsappCallRequest) GetSession() *WhatsappCallRequestSession {
	return s.Session
}

func (s *WhatsappCallRequest) GetUserNumber() *string {
	return s.UserNumber
}

func (s *WhatsappCallRequest) SetBusinessNumber(v string) *WhatsappCallRequest {
	s.BusinessNumber = &v
	return s
}

func (s *WhatsappCallRequest) SetCallAction(v string) *WhatsappCallRequest {
	s.CallAction = &v
	return s
}

func (s *WhatsappCallRequest) SetCallId(v string) *WhatsappCallRequest {
	s.CallId = &v
	return s
}

func (s *WhatsappCallRequest) SetCustSpaceId(v string) *WhatsappCallRequest {
	s.CustSpaceId = &v
	return s
}

func (s *WhatsappCallRequest) SetOwnerId(v int64) *WhatsappCallRequest {
	s.OwnerId = &v
	return s
}

func (s *WhatsappCallRequest) SetResourceOwnerAccount(v string) *WhatsappCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *WhatsappCallRequest) SetResourceOwnerId(v int64) *WhatsappCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *WhatsappCallRequest) SetSession(v *WhatsappCallRequestSession) *WhatsappCallRequest {
	s.Session = v
	return s
}

func (s *WhatsappCallRequest) SetUserNumber(v string) *WhatsappCallRequest {
	s.UserNumber = &v
	return s
}

func (s *WhatsappCallRequest) Validate() error {
	if s.Session != nil {
		if err := s.Session.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WhatsappCallRequestSession struct {
	// The Session Description Protocol (SDP) payload used for call negotiation.
	//
	// example:
	//
	// v=0\\r\\no=- 3978582128 3978582128 IN IP4 0.0.0.0\\r\\ns=Kurento Media Server\\r\\nc=IN IP4 0.0.0.0\\r\\nt=0 0\\r\\na=group:BUNDLE audio0\\r\\nm=audio 61903 RTP/SAVPF 96 0 8\\r\\na=setup:actpass\\r\\na=extmap:3 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time\\r\\na=rtpmap:96 opus/48000/2\\r\\na=rtcp:9 IN IP4 0.0.0.0\\r\\na=rtcp-mux\\r\\na=sendrecv\\r\\na=mid:audio0\\r\\na=ssrc:1935711036 cname:user3617531270@host-126a650\\r\\na=ice-ufrag:/TW8\\r\\na=ice-pwd:GRKrVHYTLGT69PYlqqWQHO\\r\\na=fingerprint:sha-256 41:5C:01:6E:C0:1B:21:F1:0D:28:95:2E:4E:42:F6:64:F3:DB:DD:E3:18:54:32:28:3D:2F:9C:80:38:FA:E0:2Ca=candidate:candidate:1 1 UDP 2015363327 47.243.79.43 46008 typ host\\r\\na=candidate:candidate:1 1 UDP 2015363327 47.243.79.43 46008 typ host\\r\\na=candidate:candidate:2 1 TCP 1015021823 47.243.79.43 9 typ host tcptype active\\r\\na=candidate:candidate:3 1 TCP 1010827519 47.243.79.43 32450 typ host tcptype passive\\r\\na=candidate:candidate:3 1 TCP 1010827519 47.243.79.43 32450 typ host tcptype passive\\r\\na=candidate:candidate:2 1 TCP 1015021823 47.243.79.43 9 typ host tcptype active\\r\\na=candidate:candidate:4 1 UDP 2015363583 fe80::216:3eff:fe22:644f 12231 typ host\\r\\na=candidate:candidate:4 1 UDP 2015363583 fe80::216:3eff:fe22:644f 12231 typ host\\r\\na=candidate:candidate:5 1 TCP 1015022079 fe80::216:3eff:fe22:644f 9 typ host tcptype active\\r\\na=candidate:candidate:5 1 TCP 1015022079 fe80::216:3eff:fe22:644f 9 typ host tcptype active\\r\\na=candidate:candidate:6 1 TCP 1010827775 fe80::216:3eff:fe22:644f 6988 typ host tcptype passive\\r\\na=candidate:candidate:6 1 TCP 1010827775 fe80::216:3eff:fe22:644f 6988 typ host tcptype passive\\r\\na=candidate:candidate:1 2 UDP 2015363326 47.243.79.43 40782 typ host\\r\\na=candidate:candidate:1 2 UDP 2015363326 47.243.79.43 40782 typ host\\r\\na=candidate:candidate:2 2 TCP 1015021822 47.243.79.43 9 typ host tcptype active\\r\\na=candidate:candidate:2 2 TCP 1015021822 47.243.79.43 9 typ host tcptype active\\r\\na=candidate:candidate:3 2 TCP 1010827518 47.243.79.43 64559 typ host tcptype passive\\r\\na=candidate:candidate:3 2 TCP 1010827518 47.243.79.43 64559 typ host tcptype passive\\r\\na=candidate:candidate:4 2 UDP 2015363582 fe80::216:3eff:fe22:644f 39166 typ host\\r\\na=candidate:candidate:5 2 TCP 1015022078 fe80::216:3eff:fe22:644f 9 typ host tcptype active\\r\\na=candidate:candidate:4 2 UDP 2015363582 fe80::216:3eff:fe22:644f 39166 typ host\\r\\na=candidate:candidate:5 2 TCP 1015022078 fe80::216:3eff:fe22:644f 9 typ host tcptype active\\r\\na=candidate:candidate:6 2 TCP 1010827774 fe80::216:3eff:fe22:644f 16286 typ host tcptype passive\\r\\na=candidate:candidate:6 2 TCP 1010827774 fe80::216:3eff:fe22:644f 16286 typ host tcptype passive\\r\\na=candidate:candidate:7 1 UDP 1679819007 47.243.79.43 46008 typ srflx raddr 172.22.181.115 rport 46008\\r\\na=candidate:candidate:7 1 UDP 1679819007 47.243.79.43 46008 typ srflx raddr 172.22.181.115 rport 46008\\r\\na=candidate:candidate:8 1 TCP 847249663 47.243.79.43 9 typ srflx raddr 172.22.181.115 rport 9 tcptype active\\r\\na=candidate:candidate:8 1 TCP 847249663 47.243.79.43 9 typ srflx raddr 172.22.181.115 rport 9 tcptype active\\r\\na=candidate:candidate:9 1 TCP 843055359 47.243.79.43 32450 typ srflx raddr 172.22.181.115 rport 32450 tcptype passive\\r\\na=candidate:candidate:9 1 TCP 843055359 47.243.79.43 32450 typ srflx raddr 172.22.181.115 rport 32450 tcptype passive\\r\\na=candidate:candidate:10 1 UDP 505413887 47.243.79.43 61903 typ relay raddr 172.22.181.115 rport 46008\\r\\na=candidate:candidate:10 1 UDP 505413887 47.243.79.43 61903 typ relay raddr 172.22.181.115 rport 46008\\r\\na=candidate:candidate:7 2 UDP 1679819006 47.243.79.43 40782 typ srflx raddr 172.22.181.115 rport 40782\\r\\na=candidate:candidate:7 2 UDP 1679819006 47.243.79.43 40782 typ srflx raddr 172.22.181.115 rport 40782\\r\\na=candidate:candidate:8 2 TCP 847249662 47.243.79.43 9 typ srflx raddr 172.22.181.115 rport 9 tcptype active\\r\\na=candidate:candidate:8 2 TCP 847249662 47.243.79.43 9 typ srflx raddr 172.22.181.115 rport 9 tcptype active\\r\\na=candidate:candidate:9 2 TCP 843055358 47.243.79.43 64559 typ srflx raddr 172.22.181.115 rport 64559 tcptype passive\\r\\na=candidate:candidate:9 2 TCP 843055358 47.243.79.43 64559 typ srflx raddr 172.22.181.115 rport 64559 tcptype passive\\r\\na=candidate:candidate:10 2 UDP 505413886 47.243.79.43 63580 typ relay raddr 172.22.181.115 rport 40782\\r\\na=candidate:candidate:10 2 UDP 505413886 47.243.79.43 63580 typ relay raddr 172.22.181.115 rport 40782\\r\\n
	Sdp *string `json:"Sdp,omitempty" xml:"Sdp,omitempty"`
	// The type of the SDP payload.
	//
	// example:
	//
	// offer
	SdpType *string `json:"SdpType,omitempty" xml:"SdpType,omitempty"`
}

func (s WhatsappCallRequestSession) String() string {
	return dara.Prettify(s)
}

func (s WhatsappCallRequestSession) GoString() string {
	return s.String()
}

func (s *WhatsappCallRequestSession) GetSdp() *string {
	return s.Sdp
}

func (s *WhatsappCallRequestSession) GetSdpType() *string {
	return s.SdpType
}

func (s *WhatsappCallRequestSession) SetSdp(v string) *WhatsappCallRequestSession {
	s.Sdp = &v
	return s
}

func (s *WhatsappCallRequestSession) SetSdpType(v string) *WhatsappCallRequestSession {
	s.SdpType = &v
	return s
}

func (s *WhatsappCallRequestSession) Validate() error {
	return dara.Validate(s)
}
