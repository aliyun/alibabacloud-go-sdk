// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCubeBySqlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCubeBySqlResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateCubeBySqlResponseBody
	GetResult() *string
	SetSuccess(v bool) *CreateCubeBySqlResponseBody
	GetSuccess() *bool
}

type CreateCubeBySqlResponseBody struct {
	// example:
	//
	// 05739b************02522b9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCubeBySqlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCubeBySqlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCubeBySqlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCubeBySqlResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateCubeBySqlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCubeBySqlResponseBody) SetRequestId(v string) *CreateCubeBySqlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCubeBySqlResponseBody) SetResult(v string) *CreateCubeBySqlResponseBody {
	s.Result = &v
	return s
}

func (s *CreateCubeBySqlResponseBody) SetSuccess(v bool) *CreateCubeBySqlResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCubeBySqlResponseBody) Validate() error {
	return dara.Validate(s)
}
