// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSampleDataByCsvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddSampleDataByCsvResponseBody
	GetRequestId() *string
}

type AddSampleDataByCsvResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSampleDataByCsvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSampleDataByCsvResponseBody) GoString() string {
	return s.String()
}

func (s *AddSampleDataByCsvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSampleDataByCsvResponseBody) SetRequestId(v string) *AddSampleDataByCsvResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSampleDataByCsvResponseBody) Validate() error {
	return dara.Validate(s)
}
