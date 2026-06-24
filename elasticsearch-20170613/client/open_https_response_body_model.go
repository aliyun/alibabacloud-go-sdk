// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenHttpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenHttpsResponseBody
	GetRequestId() *string
	SetResult(v bool) *OpenHttpsResponseBody
	GetResult() *bool
}

type OpenHttpsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result. Valid values:
	//
	// - true: The HTTPS protocol is enabled.
	//
	// - false: The HTTPS protocol failed to be enabled.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s OpenHttpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenHttpsResponseBody) GoString() string {
	return s.String()
}

func (s *OpenHttpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenHttpsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *OpenHttpsResponseBody) SetRequestId(v string) *OpenHttpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenHttpsResponseBody) SetResult(v bool) *OpenHttpsResponseBody {
	s.Result = &v
	return s
}

func (s *OpenHttpsResponseBody) Validate() error {
	return dara.Validate(s)
}
