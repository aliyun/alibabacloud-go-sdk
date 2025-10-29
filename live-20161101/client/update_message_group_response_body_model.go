// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMessageGroupResponseBody
	GetRequestId() *string
	SetResult(v *UpdateMessageGroupResponseBodyResult) *UpdateMessageGroupResponseBody
	GetResult() *UpdateMessageGroupResponseBodyResult
}

type UpdateMessageGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *UpdateMessageGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMessageGroupResponseBody) GetResult() *UpdateMessageGroupResponseBodyResult {
	return s.Result
}

func (s *UpdateMessageGroupResponseBody) SetRequestId(v string) *UpdateMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMessageGroupResponseBody) SetResult(v *UpdateMessageGroupResponseBodyResult) *UpdateMessageGroupResponseBody {
	s.Result = v
	return s
}

func (s *UpdateMessageGroupResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMessageGroupResponseBodyResult struct {
	// Indicates whether the update was successful. Valid values:
	//
	// 	- true: The update was successful.
	//
	// 	- false: The update failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMessageGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateMessageGroupResponseBodyResult) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMessageGroupResponseBodyResult) SetSuccess(v bool) *UpdateMessageGroupResponseBodyResult {
	s.Success = &v
	return s
}

func (s *UpdateMessageGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
