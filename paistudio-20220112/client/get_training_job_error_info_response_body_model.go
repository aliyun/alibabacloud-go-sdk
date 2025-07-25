// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrainingJobErrorInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorInfo(v *GetTrainingJobErrorInfoResponseBodyErrorInfo) *GetTrainingJobErrorInfoResponseBody
	GetErrorInfo() *GetTrainingJobErrorInfoResponseBodyErrorInfo
	SetRequestId(v string) *GetTrainingJobErrorInfoResponseBody
	GetRequestId() *string
}

type GetTrainingJobErrorInfoResponseBody struct {
	ErrorInfo *GetTrainingJobErrorInfoResponseBodyErrorInfo `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" type:"Struct"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTrainingJobErrorInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobErrorInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrainingJobErrorInfoResponseBody) GetErrorInfo() *GetTrainingJobErrorInfoResponseBodyErrorInfo {
	return s.ErrorInfo
}

func (s *GetTrainingJobErrorInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrainingJobErrorInfoResponseBody) SetErrorInfo(v *GetTrainingJobErrorInfoResponseBodyErrorInfo) *GetTrainingJobErrorInfoResponseBody {
	s.ErrorInfo = v
	return s
}

func (s *GetTrainingJobErrorInfoResponseBody) SetRequestId(v string) *GetTrainingJobErrorInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrainingJobErrorInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobErrorInfoResponseBodyErrorInfo struct {
	// example:
	//
	// additional info
	AdditionalInfo *string `json:"AdditionalInfo,omitempty" xml:"AdditionalInfo,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetTrainingJobErrorInfoResponseBodyErrorInfo) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobErrorInfoResponseBodyErrorInfo) GoString() string {
	return s.String()
}

func (s *GetTrainingJobErrorInfoResponseBodyErrorInfo) GetAdditionalInfo() *string {
	return s.AdditionalInfo
}

func (s *GetTrainingJobErrorInfoResponseBodyErrorInfo) GetCode() *string {
	return s.Code
}

func (s *GetTrainingJobErrorInfoResponseBodyErrorInfo) GetMessage() *string {
	return s.Message
}

func (s *GetTrainingJobErrorInfoResponseBodyErrorInfo) SetAdditionalInfo(v string) *GetTrainingJobErrorInfoResponseBodyErrorInfo {
	s.AdditionalInfo = &v
	return s
}

func (s *GetTrainingJobErrorInfoResponseBodyErrorInfo) SetCode(v string) *GetTrainingJobErrorInfoResponseBodyErrorInfo {
	s.Code = &v
	return s
}

func (s *GetTrainingJobErrorInfoResponseBodyErrorInfo) SetMessage(v string) *GetTrainingJobErrorInfoResponseBodyErrorInfo {
	s.Message = &v
	return s
}

func (s *GetTrainingJobErrorInfoResponseBodyErrorInfo) Validate() error {
	return dara.Validate(s)
}
