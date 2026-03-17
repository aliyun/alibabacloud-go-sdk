// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmartAccessGatewayClientUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *ModifySmartAccessGatewayClientUserResponseBody
	GetBandwidth() *int32
	SetClientIp(v string) *ModifySmartAccessGatewayClientUserResponseBody
	GetClientIp() *string
	SetRequestId(v string) *ModifySmartAccessGatewayClientUserResponseBody
	GetRequestId() *string
	SetUserMail(v string) *ModifySmartAccessGatewayClientUserResponseBody
	GetUserMail() *string
	SetUserName(v string) *ModifySmartAccessGatewayClientUserResponseBody
	GetUserName() *string
}

type ModifySmartAccessGatewayClientUserResponseBody struct {
	// The maximum bandwidth allocated to the client account. Unit: Kbit/s.
	//
	// example:
	//
	// 1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 10.10.10.1
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5F0078B5-8AAD-4B53-8351-4C91B8EA528A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The email address of the client account.
	//
	// example:
	//
	// username@example.com
	UserMail *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	// The username of the client account.
	//
	// example:
	//
	// username
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ModifySmartAccessGatewayClientUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySmartAccessGatewayClientUserResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySmartAccessGatewayClientUserResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifySmartAccessGatewayClientUserResponseBody) GetClientIp() *string {
	return s.ClientIp
}

func (s *ModifySmartAccessGatewayClientUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySmartAccessGatewayClientUserResponseBody) GetUserMail() *string {
	return s.UserMail
}

func (s *ModifySmartAccessGatewayClientUserResponseBody) GetUserName() *string {
	return s.UserName
}

func (s *ModifySmartAccessGatewayClientUserResponseBody) SetBandwidth(v int32) *ModifySmartAccessGatewayClientUserResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserResponseBody) SetClientIp(v string) *ModifySmartAccessGatewayClientUserResponseBody {
	s.ClientIp = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserResponseBody) SetRequestId(v string) *ModifySmartAccessGatewayClientUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserResponseBody) SetUserMail(v string) *ModifySmartAccessGatewayClientUserResponseBody {
	s.UserMail = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserResponseBody) SetUserName(v string) *ModifySmartAccessGatewayClientUserResponseBody {
	s.UserName = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserResponseBody) Validate() error {
	return dara.Validate(s)
}
