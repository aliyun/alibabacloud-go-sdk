// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPassed(v bool) *VerifySiteResponseBody
	GetPassed() *bool
	SetRequestId(v string) *VerifySiteResponseBody
	GetRequestId() *string
}

type VerifySiteResponseBody struct {
	Passed    *bool   `json:"Passed,omitempty" xml:"Passed,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifySiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifySiteResponseBody) GoString() string {
	return s.String()
}

func (s *VerifySiteResponseBody) GetPassed() *bool {
	return s.Passed
}

func (s *VerifySiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifySiteResponseBody) SetPassed(v bool) *VerifySiteResponseBody {
	s.Passed = &v
	return s
}

func (s *VerifySiteResponseBody) SetRequestId(v string) *VerifySiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifySiteResponseBody) Validate() error {
	return dara.Validate(s)
}
