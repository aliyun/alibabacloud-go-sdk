// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorksAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddWorksAuthorizationResponseBody
	GetRequestId() *string
	SetResult(v string) *AddWorksAuthorizationResponseBody
	GetResult() *string
	SetSuccess(v bool) *AddWorksAuthorizationResponseBody
	GetSuccess() *bool
}

type AddWorksAuthorizationResponseBody struct {
	// example:
	//
	// 78C1AA***************C462E57A95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// asdasf8****sda
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddWorksAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWorksAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorksAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWorksAuthorizationResponseBody) GetResult() *string {
	return s.Result
}

func (s *AddWorksAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddWorksAuthorizationResponseBody) SetRequestId(v string) *AddWorksAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWorksAuthorizationResponseBody) SetResult(v string) *AddWorksAuthorizationResponseBody {
	s.Result = &v
	return s
}

func (s *AddWorksAuthorizationResponseBody) SetSuccess(v bool) *AddWorksAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *AddWorksAuthorizationResponseBody) Validate() error {
	return dara.Validate(s)
}
