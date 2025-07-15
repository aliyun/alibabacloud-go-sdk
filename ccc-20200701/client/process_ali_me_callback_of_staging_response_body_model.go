// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessAliMeCallbackOfStagingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ProcessAliMeCallbackOfStagingResponseBody
	GetCode() *string
	SetData(v *ProcessAliMeCallbackOfStagingResponseBodyData) *ProcessAliMeCallbackOfStagingResponseBody
	GetData() *ProcessAliMeCallbackOfStagingResponseBodyData
	SetHttpStatusCode(v int32) *ProcessAliMeCallbackOfStagingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ProcessAliMeCallbackOfStagingResponseBody
	GetMessage() *string
	SetRequestId(v string) *ProcessAliMeCallbackOfStagingResponseBody
	GetRequestId() *string
}

type ProcessAliMeCallbackOfStagingResponseBody struct {
	Code           *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ProcessAliMeCallbackOfStagingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ProcessAliMeCallbackOfStagingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProcessAliMeCallbackOfStagingResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessAliMeCallbackOfStagingResponseBody) GetCode() *string {
	return s.Code
}

func (s *ProcessAliMeCallbackOfStagingResponseBody) GetData() *ProcessAliMeCallbackOfStagingResponseBodyData {
	return s.Data
}

func (s *ProcessAliMeCallbackOfStagingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ProcessAliMeCallbackOfStagingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ProcessAliMeCallbackOfStagingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProcessAliMeCallbackOfStagingResponseBody) SetCode(v string) *ProcessAliMeCallbackOfStagingResponseBody {
	s.Code = &v
	return s
}

func (s *ProcessAliMeCallbackOfStagingResponseBody) SetData(v *ProcessAliMeCallbackOfStagingResponseBodyData) *ProcessAliMeCallbackOfStagingResponseBody {
	s.Data = v
	return s
}

func (s *ProcessAliMeCallbackOfStagingResponseBody) SetHttpStatusCode(v int32) *ProcessAliMeCallbackOfStagingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ProcessAliMeCallbackOfStagingResponseBody) SetMessage(v string) *ProcessAliMeCallbackOfStagingResponseBody {
	s.Message = &v
	return s
}

func (s *ProcessAliMeCallbackOfStagingResponseBody) SetRequestId(v string) *ProcessAliMeCallbackOfStagingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProcessAliMeCallbackOfStagingResponseBody) Validate() error {
	return dara.Validate(s)
}

type ProcessAliMeCallbackOfStagingResponseBodyData struct {
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ProcessAliMeCallbackOfStagingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ProcessAliMeCallbackOfStagingResponseBodyData) GoString() string {
	return s.String()
}

func (s *ProcessAliMeCallbackOfStagingResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *ProcessAliMeCallbackOfStagingResponseBodyData) SetResult(v string) *ProcessAliMeCallbackOfStagingResponseBodyData {
	s.Result = &v
	return s
}

func (s *ProcessAliMeCallbackOfStagingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
