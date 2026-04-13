// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddTableResponseBody
	GetCode() *string
	SetData(v *AddTableResponseBodyData) *AddTableResponseBody
	GetData() *AddTableResponseBodyData
	SetMessage(v string) *AddTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddTableResponseBody
	GetRequestId() *string
	SetStatus(v string) *AddTableResponseBody
	GetStatus() *string
	SetSuccess(v bool) *AddTableResponseBody
	GetSuccess() *bool
}

type AddTableResponseBody struct {
	// example:
	//
	// Index.Forbidden
	Code *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 778C0B3B-03C1-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTableResponseBody) GoString() string {
	return s.String()
}

func (s *AddTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddTableResponseBody) GetData() *AddTableResponseBodyData {
	return s.Data
}

func (s *AddTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTableResponseBody) GetStatus() *string {
	return s.Status
}

func (s *AddTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddTableResponseBody) SetCode(v string) *AddTableResponseBody {
	s.Code = &v
	return s
}

func (s *AddTableResponseBody) SetData(v *AddTableResponseBodyData) *AddTableResponseBody {
	s.Data = v
	return s
}

func (s *AddTableResponseBody) SetMessage(v string) *AddTableResponseBody {
	s.Message = &v
	return s
}

func (s *AddTableResponseBody) SetRequestId(v string) *AddTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTableResponseBody) SetStatus(v string) *AddTableResponseBody {
	s.Status = &v
	return s
}

func (s *AddTableResponseBody) SetSuccess(v bool) *AddTableResponseBody {
	s.Success = &v
	return s
}

func (s *AddTableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddTableResponseBodyData struct {
	// example:
	//
	// table_b6ddc67e7df14db38b74ef5e2e0fe24e_12792097
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
}

func (s AddTableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddTableResponseBodyData) GetTableId() *string {
	return s.TableId
}

func (s *AddTableResponseBodyData) SetTableId(v string) *AddTableResponseBodyData {
	s.TableId = &v
	return s
}

func (s *AddTableResponseBodyData) Validate() error {
	return dara.Validate(s)
}
