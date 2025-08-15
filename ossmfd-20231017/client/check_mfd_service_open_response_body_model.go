// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMfdServiceOpenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CheckMfdServiceOpenResponseBody
	GetData() *bool
	SetRequestId(v string) *CheckMfdServiceOpenResponseBody
	GetRequestId() *string
}

type CheckMfdServiceOpenResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAC546A5-5EDC-5939-86A3-56DFAF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckMfdServiceOpenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckMfdServiceOpenResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMfdServiceOpenResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckMfdServiceOpenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckMfdServiceOpenResponseBody) SetData(v bool) *CheckMfdServiceOpenResponseBody {
	s.Data = &v
	return s
}

func (s *CheckMfdServiceOpenResponseBody) SetRequestId(v string) *CheckMfdServiceOpenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckMfdServiceOpenResponseBody) Validate() error {
	return dara.Validate(s)
}
