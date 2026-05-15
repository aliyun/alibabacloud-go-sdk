// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRobotParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRobotParamsResponseBody
	GetCode() *string
	SetData(v []*ListRobotParamsResponseBodyData) *ListRobotParamsResponseBody
	GetData() []*ListRobotParamsResponseBodyData
	SetMessage(v string) *ListRobotParamsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRobotParamsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRobotParamsResponseBody
	GetSuccess() *bool
}

type ListRobotParamsResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListRobotParamsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FF67D4D5-4E90-1DF5-BB8F-060BBFAD72DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRobotParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRobotParamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRobotParamsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRobotParamsResponseBody) GetData() []*ListRobotParamsResponseBodyData {
	return s.Data
}

func (s *ListRobotParamsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRobotParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRobotParamsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRobotParamsResponseBody) SetCode(v string) *ListRobotParamsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRobotParamsResponseBody) SetData(v []*ListRobotParamsResponseBodyData) *ListRobotParamsResponseBody {
	s.Data = v
	return s
}

func (s *ListRobotParamsResponseBody) SetMessage(v string) *ListRobotParamsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRobotParamsResponseBody) SetRequestId(v string) *ListRobotParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRobotParamsResponseBody) SetSuccess(v bool) *ListRobotParamsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRobotParamsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRobotParamsResponseBodyData struct {
	// example:
	//
	// 0
	IsEmpty *int32 `json:"IsEmpty,omitempty" xml:"IsEmpty,omitempty"`
	// example:
	//
	// name
	ParamCode *string `json:"ParamCode,omitempty" xml:"ParamCode,omitempty"`
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
}

func (s ListRobotParamsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRobotParamsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRobotParamsResponseBodyData) GetIsEmpty() *int32 {
	return s.IsEmpty
}

func (s *ListRobotParamsResponseBodyData) GetParamCode() *string {
	return s.ParamCode
}

func (s *ListRobotParamsResponseBodyData) GetParamName() *string {
	return s.ParamName
}

func (s *ListRobotParamsResponseBodyData) SetIsEmpty(v int32) *ListRobotParamsResponseBodyData {
	s.IsEmpty = &v
	return s
}

func (s *ListRobotParamsResponseBodyData) SetParamCode(v string) *ListRobotParamsResponseBodyData {
	s.ParamCode = &v
	return s
}

func (s *ListRobotParamsResponseBodyData) SetParamName(v string) *ListRobotParamsResponseBodyData {
	s.ParamName = &v
	return s
}

func (s *ListRobotParamsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
