// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryTokenResponseBody
	GetRequestId() *string
	SetTokenStatus(v bool) *QueryTokenResponseBody
	GetTokenStatus() *bool
}

type QueryTokenResponseBody struct {
	// The unique ID that the system generates for the request. This parameter is a common parameter.
	//
	// example:
	//
	// 5C8AADD0-6A95-436D-AFA0-3405CCE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the queried token. Valid values:
	//
	// 	- **true**: indicates the token is valid.
	//
	// 	- **false**: indicates the token is invalid.
	//
	// example:
	//
	// true
	TokenStatus *bool `json:"TokenStatus,omitempty" xml:"TokenStatus,omitempty"`
}

func (s QueryTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTokenResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTokenResponseBody) GetTokenStatus() *bool {
	return s.TokenStatus
}

func (s *QueryTokenResponseBody) SetRequestId(v string) *QueryTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTokenResponseBody) SetTokenStatus(v bool) *QueryTokenResponseBody {
	s.TokenStatus = &v
	return s
}

func (s *QueryTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
