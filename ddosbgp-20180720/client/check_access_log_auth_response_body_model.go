// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccessLogAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLogAuth(v bool) *CheckAccessLogAuthResponseBody
	GetAccessLogAuth() *bool
	SetRequestId(v string) *CheckAccessLogAuthResponseBody
	GetRequestId() *string
}

type CheckAccessLogAuthResponseBody struct {
	// Indicates whether Anti-DDoS Origin was authorized to access Simple Log Service. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AccessLogAuth *bool `json:"AccessLogAuth,omitempty" xml:"AccessLogAuth,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 864FE2F4-CB2E-4024-B9EF-D59FD08ABD41
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAccessLogAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAccessLogAuthResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccessLogAuthResponseBody) GetAccessLogAuth() *bool {
	return s.AccessLogAuth
}

func (s *CheckAccessLogAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAccessLogAuthResponseBody) SetAccessLogAuth(v bool) *CheckAccessLogAuthResponseBody {
	s.AccessLogAuth = &v
	return s
}

func (s *CheckAccessLogAuthResponseBody) SetRequestId(v string) *CheckAccessLogAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAccessLogAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
