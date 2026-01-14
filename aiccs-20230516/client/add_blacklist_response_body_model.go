// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBlacklistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *AddBlacklistResponseBody
	GetCode() *int64
	SetMessage(v string) *AddBlacklistResponseBody
	GetMessage() *string
	SetModel(v *AddBlacklistResponseBodyModel) *AddBlacklistResponseBody
	GetModel() *AddBlacklistResponseBodyModel
	SetRequestId(v string) *AddBlacklistResponseBody
	GetRequestId() *string
	SetSuccess(v string) *AddBlacklistResponseBody
	GetSuccess() *string
	SetTimestamp(v int64) *AddBlacklistResponseBody
	GetTimestamp() *int64
}

type AddBlacklistResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *AddBlacklistResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AddBlacklistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *AddBlacklistResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AddBlacklistResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddBlacklistResponseBody) GetModel() *AddBlacklistResponseBodyModel {
	return s.Model
}

func (s *AddBlacklistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBlacklistResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AddBlacklistResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AddBlacklistResponseBody) SetCode(v int64) *AddBlacklistResponseBody {
	s.Code = &v
	return s
}

func (s *AddBlacklistResponseBody) SetMessage(v string) *AddBlacklistResponseBody {
	s.Message = &v
	return s
}

func (s *AddBlacklistResponseBody) SetModel(v *AddBlacklistResponseBodyModel) *AddBlacklistResponseBody {
	s.Model = v
	return s
}

func (s *AddBlacklistResponseBody) SetRequestId(v string) *AddBlacklistResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBlacklistResponseBody) SetSuccess(v string) *AddBlacklistResponseBody {
	s.Success = &v
	return s
}

func (s *AddBlacklistResponseBody) SetTimestamp(v int64) *AddBlacklistResponseBody {
	s.Timestamp = &v
	return s
}

func (s *AddBlacklistResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddBlacklistResponseBodyModel struct {
	// 错误手机号
	UnHandleNumbers []*string `json:"UnHandleNumbers,omitempty" xml:"UnHandleNumbers,omitempty" type:"Repeated"`
}

func (s AddBlacklistResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s AddBlacklistResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AddBlacklistResponseBodyModel) GetUnHandleNumbers() []*string {
	return s.UnHandleNumbers
}

func (s *AddBlacklistResponseBodyModel) SetUnHandleNumbers(v []*string) *AddBlacklistResponseBodyModel {
	s.UnHandleNumbers = v
	return s
}

func (s *AddBlacklistResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
