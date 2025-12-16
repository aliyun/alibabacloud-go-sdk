// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenewAppGroupResponseBody
	GetRequestId() *string
	SetResult(v bool) *RenewAppGroupResponseBody
	GetResult() *bool
}

type RenewAppGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the application was renewed.
	//
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RenewAppGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAppGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewAppGroupResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RenewAppGroupResponseBody) SetRequestId(v string) *RenewAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewAppGroupResponseBody) SetResult(v bool) *RenewAppGroupResponseBody {
	s.Result = &v
	return s
}

func (s *RenewAppGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
