// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateShrinkNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidateShrinkNodesResponseBody
	GetRequestId() *string
	SetResult(v bool) *ValidateShrinkNodesResponseBody
	GetResult() *bool
}

type ValidateShrinkNodesResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateShrinkNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateShrinkNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateShrinkNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateShrinkNodesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ValidateShrinkNodesResponseBody) SetRequestId(v string) *ValidateShrinkNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateShrinkNodesResponseBody) SetResult(v bool) *ValidateShrinkNodesResponseBody {
	s.Result = &v
	return s
}

func (s *ValidateShrinkNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
