// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataLevelPermissionWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDataLevelPermissionWhiteListResponseBody
	GetRequestId() *string
	SetResult(v bool) *SetDataLevelPermissionWhiteListResponseBody
	GetResult() *bool
	SetSuccess(v bool) *SetDataLevelPermissionWhiteListResponseBody
	GetSuccess() *bool
}

type SetDataLevelPermissionWhiteListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
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

func (s SetDataLevelPermissionWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDataLevelPermissionWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDataLevelPermissionWhiteListResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SetDataLevelPermissionWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetDataLevelPermissionWhiteListResponseBody) SetRequestId(v string) *SetDataLevelPermissionWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDataLevelPermissionWhiteListResponseBody) SetResult(v bool) *SetDataLevelPermissionWhiteListResponseBody {
	s.Result = &v
	return s
}

func (s *SetDataLevelPermissionWhiteListResponseBody) SetSuccess(v bool) *SetDataLevelPermissionWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *SetDataLevelPermissionWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
