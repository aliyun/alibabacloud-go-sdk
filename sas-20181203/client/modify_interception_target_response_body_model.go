// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInterceptionTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInterceptionTargetResponseBody
	GetRequestId() *string
	SetResult(v bool) *ModifyInterceptionTargetResponseBody
	GetResult() *bool
}

type ModifyInterceptionTargetResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 16CB4467-56AE-546C-BF19-AD4584C0DD03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ModifyInterceptionTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInterceptionTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInterceptionTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInterceptionTargetResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ModifyInterceptionTargetResponseBody) SetRequestId(v string) *ModifyInterceptionTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInterceptionTargetResponseBody) SetResult(v bool) *ModifyInterceptionTargetResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyInterceptionTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
