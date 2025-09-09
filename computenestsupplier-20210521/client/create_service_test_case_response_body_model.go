// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceTestCaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceTestCaseResponseBody
	GetRequestId() *string
	SetTestCaseId(v string) *CreateServiceTestCaseResponseBody
	GetTestCaseId() *string
}

type CreateServiceTestCaseResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E50287CB-AABF-4877-92C0-289B339A1546
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The test case Id
	//
	// example:
	//
	// stc-5ba03a6a9a2746be8739
	TestCaseId *string `json:"TestCaseId,omitempty" xml:"TestCaseId,omitempty"`
}

func (s CreateServiceTestCaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceTestCaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceTestCaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceTestCaseResponseBody) GetTestCaseId() *string {
	return s.TestCaseId
}

func (s *CreateServiceTestCaseResponseBody) SetRequestId(v string) *CreateServiceTestCaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceTestCaseResponseBody) SetTestCaseId(v string) *CreateServiceTestCaseResponseBody {
	s.TestCaseId = &v
	return s
}

func (s *CreateServiceTestCaseResponseBody) Validate() error {
	return dara.Validate(s)
}
