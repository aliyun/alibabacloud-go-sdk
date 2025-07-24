// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceDeployableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckResults(v []*CheckServiceDeployableResponseBodyCheckResults) *CheckServiceDeployableResponseBody
	GetCheckResults() []*CheckServiceDeployableResponseBodyCheckResults
	SetRequestId(v string) *CheckServiceDeployableResponseBody
	GetRequestId() *string
}

type CheckServiceDeployableResponseBody struct {
	// Inspection result.
	CheckResults []*CheckServiceDeployableResponseBodyCheckResults `json:"CheckResults,omitempty" xml:"CheckResults,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckServiceDeployableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceDeployableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceDeployableResponseBody) GetCheckResults() []*CheckServiceDeployableResponseBodyCheckResults {
	return s.CheckResults
}

func (s *CheckServiceDeployableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckServiceDeployableResponseBody) SetCheckResults(v []*CheckServiceDeployableResponseBodyCheckResults) *CheckServiceDeployableResponseBody {
	s.CheckResults = v
	return s
}

func (s *CheckServiceDeployableResponseBody) SetRequestId(v string) *CheckServiceDeployableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceDeployableResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckServiceDeployableResponseBodyCheckResults struct {
	// Returns a hint message for the result.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Check type, invalid values:
	//
	// - Balance ：Account balance.
	//
	// - Quota:  Account quota.
	//
	// example:
	//
	// Balance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Inspection result.
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CheckServiceDeployableResponseBodyCheckResults) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceDeployableResponseBodyCheckResults) GoString() string {
	return s.String()
}

func (s *CheckServiceDeployableResponseBodyCheckResults) GetMessage() *string {
	return s.Message
}

func (s *CheckServiceDeployableResponseBodyCheckResults) GetType() *string {
	return s.Type
}

func (s *CheckServiceDeployableResponseBodyCheckResults) GetValue() *string {
	return s.Value
}

func (s *CheckServiceDeployableResponseBodyCheckResults) SetMessage(v string) *CheckServiceDeployableResponseBodyCheckResults {
	s.Message = &v
	return s
}

func (s *CheckServiceDeployableResponseBodyCheckResults) SetType(v string) *CheckServiceDeployableResponseBodyCheckResults {
	s.Type = &v
	return s
}

func (s *CheckServiceDeployableResponseBodyCheckResults) SetValue(v string) *CheckServiceDeployableResponseBodyCheckResults {
	s.Value = &v
	return s
}

func (s *CheckServiceDeployableResponseBodyCheckResults) Validate() error {
	return dara.Validate(s)
}
