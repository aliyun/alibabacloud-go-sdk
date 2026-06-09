// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTableResponseBody
	GetCode() *string
	SetData(v *CreateTableResponseBodyData) *CreateTableResponseBody
	GetData() *CreateTableResponseBodyData
	SetMessage(v string) *CreateTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTableResponseBody
	GetSuccess() *bool
}

type CreateTableResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"TableARN":"acs:eventbridge:cn-hangzhou:123456789:catalog/my_catalog/namespace/my_namespace/table/my_table"}
	Data *CreateTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTableResponseBody) GetData() *CreateTableResponseBodyData {
	return s.Data
}

func (s *CreateTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTableResponseBody) SetCode(v string) *CreateTableResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTableResponseBody) SetData(v *CreateTableResponseBodyData) *CreateTableResponseBody {
	s.Data = v
	return s
}

func (s *CreateTableResponseBody) SetMessage(v string) *CreateTableResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTableResponseBody) SetRequestId(v string) *CreateTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTableResponseBody) SetSuccess(v bool) *CreateTableResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTableResponseBodyData struct {
	// example:
	//
	// acs:eventbridge:cn-hangzhou:123456789:catalog/my_catalog/namespace/my_namespace/table/my_table
	TableARN *string `json:"TableARN,omitempty" xml:"TableARN,omitempty"`
}

func (s CreateTableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTableResponseBodyData) GetTableARN() *string {
	return s.TableARN
}

func (s *CreateTableResponseBodyData) SetTableARN(v string) *CreateTableResponseBodyData {
	s.TableARN = &v
	return s
}

func (s *CreateTableResponseBodyData) Validate() error {
	return dara.Validate(s)
}
