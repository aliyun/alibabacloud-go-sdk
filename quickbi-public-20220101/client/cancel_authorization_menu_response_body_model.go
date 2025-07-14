// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAuthorizationMenuResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelAuthorizationMenuResponseBody
	GetRequestId() *string
	SetResult(v int32) *CancelAuthorizationMenuResponseBody
	GetResult() *int32
	SetSuccess(v bool) *CancelAuthorizationMenuResponseBody
	GetSuccess() *bool
}

type CancelAuthorizationMenuResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Number of menus successfully unauthorized.
	//
	// example:
	//
	// 2
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelAuthorizationMenuResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAuthorizationMenuResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAuthorizationMenuResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAuthorizationMenuResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *CancelAuthorizationMenuResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelAuthorizationMenuResponseBody) SetRequestId(v string) *CancelAuthorizationMenuResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAuthorizationMenuResponseBody) SetResult(v int32) *CancelAuthorizationMenuResponseBody {
	s.Result = &v
	return s
}

func (s *CancelAuthorizationMenuResponseBody) SetSuccess(v bool) *CancelAuthorizationMenuResponseBody {
	s.Success = &v
	return s
}

func (s *CancelAuthorizationMenuResponseBody) Validate() error {
	return dara.Validate(s)
}
