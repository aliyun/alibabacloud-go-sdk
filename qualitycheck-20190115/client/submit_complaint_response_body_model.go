// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitComplaintResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitComplaintResponseBody
	GetCode() *string
	SetData(v string) *SubmitComplaintResponseBody
	GetData() *string
	SetMessage(v string) *SubmitComplaintResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitComplaintResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitComplaintResponseBody
	GetSuccess() *bool
}

type SubmitComplaintResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 90
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitComplaintResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitComplaintResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitComplaintResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitComplaintResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitComplaintResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitComplaintResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitComplaintResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitComplaintResponseBody) SetCode(v string) *SubmitComplaintResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitComplaintResponseBody) SetData(v string) *SubmitComplaintResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitComplaintResponseBody) SetMessage(v string) *SubmitComplaintResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitComplaintResponseBody) SetRequestId(v string) *SubmitComplaintResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitComplaintResponseBody) SetSuccess(v bool) *SubmitComplaintResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitComplaintResponseBody) Validate() error {
	return dara.Validate(s)
}
