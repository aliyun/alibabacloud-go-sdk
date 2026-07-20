// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSasTrialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSasTrialResponseBody
	GetCode() *string
	SetData(v *CreateSasTrialResponseBodyData) *CreateSasTrialResponseBody
	GetData() *CreateSasTrialResponseBodyData
	SetMessage(v string) *CreateSasTrialResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSasTrialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSasTrialResponseBody
	GetSuccess() *bool
}

type CreateSasTrialResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateSasTrialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6B57D35D-9DAC-5393-AE39-07697E37C2E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSasTrialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSasTrialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSasTrialResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSasTrialResponseBody) GetData() *CreateSasTrialResponseBodyData {
	return s.Data
}

func (s *CreateSasTrialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSasTrialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSasTrialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSasTrialResponseBody) SetCode(v string) *CreateSasTrialResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSasTrialResponseBody) SetData(v *CreateSasTrialResponseBodyData) *CreateSasTrialResponseBody {
	s.Data = v
	return s
}

func (s *CreateSasTrialResponseBody) SetMessage(v string) *CreateSasTrialResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSasTrialResponseBody) SetRequestId(v string) *CreateSasTrialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSasTrialResponseBody) SetSuccess(v bool) *CreateSasTrialResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSasTrialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSasTrialResponseBodyData struct {
	Body *CreateSasTrialResponseBodyDataBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
}

func (s CreateSasTrialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSasTrialResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSasTrialResponseBodyData) GetBody() *CreateSasTrialResponseBodyDataBody {
	return s.Body
}

func (s *CreateSasTrialResponseBodyData) SetBody(v *CreateSasTrialResponseBodyDataBody) *CreateSasTrialResponseBodyData {
	s.Body = v
	return s
}

func (s *CreateSasTrialResponseBodyData) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSasTrialResponseBodyDataBody struct {
	// example:
	//
	// F7C74264-DF12-56D5-869B-C4B11DD88BA2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSasTrialResponseBodyDataBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSasTrialResponseBodyDataBody) GoString() string {
	return s.String()
}

func (s *CreateSasTrialResponseBodyDataBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSasTrialResponseBodyDataBody) SetRequestId(v string) *CreateSasTrialResponseBodyDataBody {
	s.RequestId = &v
	return s
}

func (s *CreateSasTrialResponseBodyDataBody) Validate() error {
	return dara.Validate(s)
}
