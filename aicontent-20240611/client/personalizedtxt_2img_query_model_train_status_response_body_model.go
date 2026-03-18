// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgQueryModelTrainStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) *Personalizedtxt2imgQueryModelTrainStatusResponseBody
	GetData() *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData
	SetErrCode(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainStatusResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *Personalizedtxt2imgQueryModelTrainStatusResponseBody
	GetSuccess() *bool
}

type Personalizedtxt2imgQueryModelTrainStatusResponseBody struct {
	// example:
	//
	// []
	Data *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponseBody) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) GetData() *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData {
	return s.Data
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetData(v *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.Data = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetErrCode(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetErrMessage(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetHttpStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetRequestId(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetSuccess(v bool) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.Success = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Personalizedtxt2imgQueryModelTrainStatusResponseBodyData struct {
	// example:
	//
	// FINISHED
	ModelTrainStatus *string `json:"modelTrainStatus,omitempty" xml:"modelTrainStatus,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) GetModelTrainStatus() *string {
	return s.ModelTrainStatus
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) SetModelTrainStatus(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData {
	s.ModelTrainStatus = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
