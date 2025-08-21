// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseManagedIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseManagedIndexResponseBody
	GetRequestId() *string
	SetResult(v bool) *CloseManagedIndexResponseBody
	GetResult() *bool
}

type CloseManagedIndexResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the index\\"s cloud management has been successfully closed:
	//
	// - true: Success
	//
	// - false: Failure
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloseManagedIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseManagedIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CloseManagedIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseManagedIndexResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CloseManagedIndexResponseBody) SetRequestId(v string) *CloseManagedIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseManagedIndexResponseBody) SetResult(v bool) *CloseManagedIndexResponseBody {
	s.Result = &v
	return s
}

func (s *CloseManagedIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
