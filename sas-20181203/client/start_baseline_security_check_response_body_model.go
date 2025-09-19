// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBaselineSecurityCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartBaselineSecurityCheckResponseBody
	GetRequestId() *string
}

type StartBaselineSecurityCheckResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 48D2E9A9-A1B0-4295-B727-0995757C47E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartBaselineSecurityCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartBaselineSecurityCheckResponseBody) GoString() string {
	return s.String()
}

func (s *StartBaselineSecurityCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartBaselineSecurityCheckResponseBody) SetRequestId(v string) *StartBaselineSecurityCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartBaselineSecurityCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
