// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTableResponseBody
	GetCode() *string
	SetData(v *UpdateTableResponseBodyData) *UpdateTableResponseBody
	GetData() *UpdateTableResponseBodyData
	SetMessage(v string) *UpdateTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTableResponseBody
	GetSuccess() *bool
}

type UpdateTableResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"TableARN":"acs:eventbridge:cn-hangzhou:123456789:catalog/my_catalog/namespace/my_namespace/table/my_table"}
	Data *UpdateTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UpdateTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTableResponseBody) GetData() *UpdateTableResponseBodyData {
	return s.Data
}

func (s *UpdateTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTableResponseBody) SetCode(v string) *UpdateTableResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTableResponseBody) SetData(v *UpdateTableResponseBodyData) *UpdateTableResponseBody {
	s.Data = v
	return s
}

func (s *UpdateTableResponseBody) SetMessage(v string) *UpdateTableResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTableResponseBody) SetRequestId(v string) *UpdateTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTableResponseBody) SetSuccess(v bool) *UpdateTableResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTableResponseBodyData struct {
	// example:
	//
	// acs:eventbridge:cn-hangzhou:123456789:catalog/my_catalog/namespace/my_namespace/table/my_table
	TableARN *string `json:"TableARN,omitempty" xml:"TableARN,omitempty"`
}

func (s UpdateTableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateTableResponseBodyData) GetTableARN() *string {
	return s.TableARN
}

func (s *UpdateTableResponseBodyData) SetTableARN(v string) *UpdateTableResponseBodyData {
	s.TableARN = &v
	return s
}

func (s *UpdateTableResponseBodyData) Validate() error {
	return dara.Validate(s)
}
