// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSampleDataByTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddSampleDataByTextResponseBody
	GetRequestId() *string
}

type AddSampleDataByTextResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSampleDataByTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSampleDataByTextResponseBody) GoString() string {
	return s.String()
}

func (s *AddSampleDataByTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSampleDataByTextResponseBody) SetRequestId(v string) *AddSampleDataByTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSampleDataByTextResponseBody) Validate() error {
	return dara.Validate(s)
}
