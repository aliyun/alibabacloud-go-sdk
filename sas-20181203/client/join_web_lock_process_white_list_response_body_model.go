// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinWebLockProcessWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *JoinWebLockProcessWhiteListResponseBody
	GetRequestId() *string
}

type JoinWebLockProcessWhiteListResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A47D621A-193E-5BDA-ADFA-A0D3133E199C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JoinWebLockProcessWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s JoinWebLockProcessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *JoinWebLockProcessWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *JoinWebLockProcessWhiteListResponseBody) SetRequestId(v string) *JoinWebLockProcessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinWebLockProcessWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
