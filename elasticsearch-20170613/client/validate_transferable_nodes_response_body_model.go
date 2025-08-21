// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateTransferableNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidateTransferableNodesResponseBody
	GetRequestId() *string
	SetResult(v bool) *ValidateTransferableNodesResponseBody
	GetResult() *bool
}

type ValidateTransferableNodesResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateTransferableNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateTransferableNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateTransferableNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateTransferableNodesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ValidateTransferableNodesResponseBody) SetRequestId(v string) *ValidateTransferableNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateTransferableNodesResponseBody) SetResult(v bool) *ValidateTransferableNodesResponseBody {
	s.Result = &v
	return s
}

func (s *ValidateTransferableNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
