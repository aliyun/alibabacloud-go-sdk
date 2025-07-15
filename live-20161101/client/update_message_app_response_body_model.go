// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMessageAppResponseBody
	GetRequestId() *string
	SetResult(v *UpdateMessageAppResponseBodyResult) *UpdateMessageAppResponseBody
	GetResult() *UpdateMessageAppResponseBodyResult
}

type UpdateMessageAppResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *UpdateMessageAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateMessageAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMessageAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMessageAppResponseBody) GetResult() *UpdateMessageAppResponseBodyResult {
	return s.Result
}

func (s *UpdateMessageAppResponseBody) SetRequestId(v string) *UpdateMessageAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMessageAppResponseBody) SetResult(v *UpdateMessageAppResponseBodyResult) *UpdateMessageAppResponseBody {
	s.Result = v
	return s
}

func (s *UpdateMessageAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateMessageAppResponseBodyResult struct {
	// Indicates whether the update is successful. Valid values:
	//
	// 	- true: The update is successful.
	//
	// 	- false: The update failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMessageAppResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateMessageAppResponseBodyResult) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMessageAppResponseBodyResult) SetSuccess(v bool) *UpdateMessageAppResponseBodyResult {
	s.Success = &v
	return s
}

func (s *UpdateMessageAppResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
