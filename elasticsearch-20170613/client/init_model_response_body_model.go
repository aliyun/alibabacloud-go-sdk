// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InitModelResponseBody
	GetRequestId() *string
	SetResult(v bool) *InitModelResponseBody
	GetResult() *bool
}

type InitModelResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InitModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitModelResponseBody) GoString() string {
	return s.String()
}

func (s *InitModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitModelResponseBody) GetResult() *bool {
	return s.Result
}

func (s *InitModelResponseBody) SetRequestId(v string) *InitModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitModelResponseBody) SetResult(v bool) *InitModelResponseBody {
	s.Result = &v
	return s
}

func (s *InitModelResponseBody) Validate() error {
	return dara.Validate(s)
}
