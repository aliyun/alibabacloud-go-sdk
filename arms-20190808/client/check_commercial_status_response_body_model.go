// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCommercialStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CheckCommercialStatusResponseBody
	GetData() *string
	SetRequestId(v string) *CheckCommercialStatusResponseBody
	GetRequestId() *string
}

type CheckCommercialStatusResponseBody struct {
	// The returned struct.
	//
	// example:
	//
	// True
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A474FF8-7861-4D00-81B5-5BC3DA4E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckCommercialStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCommercialStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCommercialStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *CheckCommercialStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCommercialStatusResponseBody) SetData(v string) *CheckCommercialStatusResponseBody {
	s.Data = &v
	return s
}

func (s *CheckCommercialStatusResponseBody) SetRequestId(v string) *CheckCommercialStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCommercialStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
