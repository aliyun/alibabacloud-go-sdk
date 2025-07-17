// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CheckServiceStatusResponseBody
	GetData() *string
	SetRequestId(v string) *CheckServiceStatusResponseBody
	GetRequestId() *string
}

type CheckServiceStatusResponseBody struct {
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true.
	//
	// 	- false.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID. You can use the ID to find logs and troubleshoot issues.
	//
	// example:
	//
	// 5710C923-AD09-4293-9E11-DCBE3D15F8D4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *CheckServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckServiceStatusResponseBody) SetData(v string) *CheckServiceStatusResponseBody {
	s.Data = &v
	return s
}

func (s *CheckServiceStatusResponseBody) SetRequestId(v string) *CheckServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
