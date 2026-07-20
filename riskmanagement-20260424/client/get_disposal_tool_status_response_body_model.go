// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDisposalToolStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDisposalToolStatusResponseBody
	GetCode() *string
	SetData(v *GetDisposalToolStatusResponseBodyData) *GetDisposalToolStatusResponseBody
	GetData() *GetDisposalToolStatusResponseBodyData
	SetMessage(v string) *GetDisposalToolStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDisposalToolStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDisposalToolStatusResponseBody
	GetSuccess() *bool
}

type GetDisposalToolStatusResponseBody struct {
	// example:
	//
	// Success
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDisposalToolStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6D462855-7835-5F91-835E-A62E44EC01CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDisposalToolStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDisposalToolStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDisposalToolStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDisposalToolStatusResponseBody) GetData() *GetDisposalToolStatusResponseBodyData {
	return s.Data
}

func (s *GetDisposalToolStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDisposalToolStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDisposalToolStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDisposalToolStatusResponseBody) SetCode(v string) *GetDisposalToolStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetDisposalToolStatusResponseBody) SetData(v *GetDisposalToolStatusResponseBodyData) *GetDisposalToolStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetDisposalToolStatusResponseBody) SetMessage(v string) *GetDisposalToolStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetDisposalToolStatusResponseBody) SetRequestId(v string) *GetDisposalToolStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDisposalToolStatusResponseBody) SetSuccess(v bool) *GetDisposalToolStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetDisposalToolStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDisposalToolStatusResponseBodyData struct {
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDisposalToolStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDisposalToolStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDisposalToolStatusResponseBodyData) GetStatus() *bool {
	return s.Status
}

func (s *GetDisposalToolStatusResponseBodyData) SetStatus(v bool) *GetDisposalToolStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDisposalToolStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
