// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTableResponseBody
	GetCode() *string
	SetData(v *Table) *GetTableResponseBody
	GetData() *Table
	SetMessage(v string) *GetTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableResponseBody
	GetSuccess() *bool
}

type GetTableResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *Table  `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTableResponseBody) GetData() *Table {
	return s.Data
}

func (s *GetTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableResponseBody) SetCode(v string) *GetTableResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableResponseBody) SetData(v *Table) *GetTableResponseBody {
	s.Data = v
	return s
}

func (s *GetTableResponseBody) SetMessage(v string) *GetTableResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableResponseBody) SetRequestId(v string) *GetTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableResponseBody) SetSuccess(v bool) *GetTableResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
