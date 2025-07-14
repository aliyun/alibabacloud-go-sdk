// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLevelPermissionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataLevelPermissionStatusResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateDataLevelPermissionStatusResponseBody
	GetResult() *bool
	SetSuccess(v bool) *UpdateDataLevelPermissionStatusResponseBody
	GetSuccess() *bool
}

type UpdateDataLevelPermissionStatusResponseBody struct {
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
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

func (s UpdateDataLevelPermissionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLevelPermissionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataLevelPermissionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataLevelPermissionStatusResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateDataLevelPermissionStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataLevelPermissionStatusResponseBody) SetRequestId(v string) *UpdateDataLevelPermissionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataLevelPermissionStatusResponseBody) SetResult(v bool) *UpdateDataLevelPermissionStatusResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateDataLevelPermissionStatusResponseBody) SetSuccess(v bool) *UpdateDataLevelPermissionStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataLevelPermissionStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
