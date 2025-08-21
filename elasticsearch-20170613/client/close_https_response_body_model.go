// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseHttpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseHttpsResponseBody
	GetRequestId() *string
	SetResult(v bool) *CloseHttpsResponseBody
	GetResult() *bool
}

type CloseHttpsResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloseHttpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseHttpsResponseBody) GoString() string {
	return s.String()
}

func (s *CloseHttpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseHttpsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CloseHttpsResponseBody) SetRequestId(v string) *CloseHttpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseHttpsResponseBody) SetResult(v bool) *CloseHttpsResponseBody {
	s.Result = &v
	return s
}

func (s *CloseHttpsResponseBody) Validate() error {
	return dara.Validate(s)
}
