// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateIndexResponseBody
	GetCode() *string
	SetData(v *UpdateIndexResponseBodyData) *UpdateIndexResponseBody
	GetData() *UpdateIndexResponseBodyData
	SetMessage(v string) *UpdateIndexResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateIndexResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateIndexResponseBody
	GetStatus() *string
	SetSuccess(v bool) *UpdateIndexResponseBody
	GetSuccess() *bool
}

type UpdateIndexResponseBody struct {
	// example:
	//
	// Success
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateIndexResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 17204B98-7734-4F9A-8464-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIndexResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIndexResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateIndexResponseBody) GetData() *UpdateIndexResponseBodyData {
	return s.Data
}

func (s *UpdateIndexResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIndexResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateIndexResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateIndexResponseBody) SetCode(v string) *UpdateIndexResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateIndexResponseBody) SetData(v *UpdateIndexResponseBodyData) *UpdateIndexResponseBody {
	s.Data = v
	return s
}

func (s *UpdateIndexResponseBody) SetMessage(v string) *UpdateIndexResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateIndexResponseBody) SetRequestId(v string) *UpdateIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIndexResponseBody) SetStatus(v string) *UpdateIndexResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateIndexResponseBody) SetSuccess(v bool) *UpdateIndexResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateIndexResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIndexResponseBodyData struct {
	// example:
	//
	// 79c0alxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateIndexResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateIndexResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateIndexResponseBodyData) GetId() *string {
	return s.Id
}

func (s *UpdateIndexResponseBodyData) SetId(v string) *UpdateIndexResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateIndexResponseBodyData) Validate() error {
	return dara.Validate(s)
}
