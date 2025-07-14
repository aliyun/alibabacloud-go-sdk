// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeVisibilityModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeVisibilityModelResponseBody
	GetRequestId() *string
	SetResult(v int32) *ChangeVisibilityModelResponseBody
	GetResult() *int32
	SetSuccess(v bool) *ChangeVisibilityModelResponseBody
	GetSuccess() *bool
}

type ChangeVisibilityModelResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of menus that are successfully modified.
	//
	// example:
	//
	// 2
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeVisibilityModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeVisibilityModelResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeVisibilityModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeVisibilityModelResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *ChangeVisibilityModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeVisibilityModelResponseBody) SetRequestId(v string) *ChangeVisibilityModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeVisibilityModelResponseBody) SetResult(v int32) *ChangeVisibilityModelResponseBody {
	s.Result = &v
	return s
}

func (s *ChangeVisibilityModelResponseBody) SetSuccess(v bool) *ChangeVisibilityModelResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeVisibilityModelResponseBody) Validate() error {
	return dara.Validate(s)
}
