// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmartAccessGatewayClientUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *CreateSmartAccessGatewayClientUserResponseBody
	GetBandwidth() *int32
	SetClientIp(v string) *CreateSmartAccessGatewayClientUserResponseBody
	GetClientIp() *string
	SetRequestId(v string) *CreateSmartAccessGatewayClientUserResponseBody
	GetRequestId() *string
	SetUserMail(v string) *CreateSmartAccessGatewayClientUserResponseBody
	GetUserMail() *string
	SetUserName(v string) *CreateSmartAccessGatewayClientUserResponseBody
	GetUserName() *string
}

type CreateSmartAccessGatewayClientUserResponseBody struct {
	// The maximum bandwidth value. Unit: Kbit/s.
	//
	// example:
	//
	// 20
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The IP address of the client app.
	//
	// example:
	//
	// 10.0.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 72E82F5E-66E8-4C22-BF1F-5CEB7DC132E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The email address of the client account.
	//
	// example:
	//
	// username@example.com
	UserMail *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	// The username.
	//
	// example:
	//
	// doc
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateSmartAccessGatewayClientUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartAccessGatewayClientUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmartAccessGatewayClientUserResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateSmartAccessGatewayClientUserResponseBody) GetClientIp() *string {
	return s.ClientIp
}

func (s *CreateSmartAccessGatewayClientUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSmartAccessGatewayClientUserResponseBody) GetUserMail() *string {
	return s.UserMail
}

func (s *CreateSmartAccessGatewayClientUserResponseBody) GetUserName() *string {
	return s.UserName
}

func (s *CreateSmartAccessGatewayClientUserResponseBody) SetBandwidth(v int32) *CreateSmartAccessGatewayClientUserResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserResponseBody) SetClientIp(v string) *CreateSmartAccessGatewayClientUserResponseBody {
	s.ClientIp = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserResponseBody) SetRequestId(v string) *CreateSmartAccessGatewayClientUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserResponseBody) SetUserMail(v string) *CreateSmartAccessGatewayClientUserResponseBody {
	s.UserMail = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserResponseBody) SetUserName(v string) *CreateSmartAccessGatewayClientUserResponseBody {
	s.UserName = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserResponseBody) Validate() error {
	return dara.Validate(s)
}
