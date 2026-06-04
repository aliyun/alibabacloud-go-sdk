// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomAttributeResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteCustomAttributeResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DeleteCustomAttributeResponseBody
	GetSuccess() *bool
}

type DeleteCustomAttributeResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 7B4D81F1-98CF-500C-8D25-2865EB652EA0
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

func (s DeleteCustomAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomAttributeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteCustomAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomAttributeResponseBody) SetRequestId(v string) *DeleteCustomAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomAttributeResponseBody) SetResult(v bool) *DeleteCustomAttributeResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteCustomAttributeResponseBody) SetSuccess(v bool) *DeleteCustomAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
