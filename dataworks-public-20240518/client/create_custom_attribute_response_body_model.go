// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCustomAttributeResponseBody
	GetRequestId() *string
	SetResult(v bool) *CreateCustomAttributeResponseBody
	GetResult() *bool
	SetSuccess(v bool) *CreateCustomAttributeResponseBody
	GetSuccess() *bool
}

type CreateCustomAttributeResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5A1E9EBB-FEA6-5BBB-B7BE-BFC0FB3F8C71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCustomAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomAttributeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CreateCustomAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCustomAttributeResponseBody) SetRequestId(v string) *CreateCustomAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomAttributeResponseBody) SetResult(v bool) *CreateCustomAttributeResponseBody {
	s.Result = &v
	return s
}

func (s *CreateCustomAttributeResponseBody) SetSuccess(v bool) *CreateCustomAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
