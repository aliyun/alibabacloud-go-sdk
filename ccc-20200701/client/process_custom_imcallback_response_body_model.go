// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessCustomIMCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ProcessCustomIMCallbackResponseBody
	GetCode() *string
	SetData(v string) *ProcessCustomIMCallbackResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ProcessCustomIMCallbackResponseBody
	GetHttpStatusCode() *int32
	SetJobId(v string) *ProcessCustomIMCallbackResponseBody
	GetJobId() *string
	SetMessage(v string) *ProcessCustomIMCallbackResponseBody
	GetMessage() *string
	SetParams(v []*string) *ProcessCustomIMCallbackResponseBody
	GetParams() []*string
	SetRequestId(v string) *ProcessCustomIMCallbackResponseBody
	GetRequestId() *string
}

type ProcessCustomIMCallbackResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// chat-65382141036853491
	JobId   *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 2778FA12-EDD6-42AA-9B15-AF855072E5E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ProcessCustomIMCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProcessCustomIMCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessCustomIMCallbackResponseBody) GetCode() *string {
	return s.Code
}

func (s *ProcessCustomIMCallbackResponseBody) GetData() *string {
	return s.Data
}

func (s *ProcessCustomIMCallbackResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ProcessCustomIMCallbackResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *ProcessCustomIMCallbackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ProcessCustomIMCallbackResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ProcessCustomIMCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProcessCustomIMCallbackResponseBody) SetCode(v string) *ProcessCustomIMCallbackResponseBody {
	s.Code = &v
	return s
}

func (s *ProcessCustomIMCallbackResponseBody) SetData(v string) *ProcessCustomIMCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *ProcessCustomIMCallbackResponseBody) SetHttpStatusCode(v int32) *ProcessCustomIMCallbackResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ProcessCustomIMCallbackResponseBody) SetJobId(v string) *ProcessCustomIMCallbackResponseBody {
	s.JobId = &v
	return s
}

func (s *ProcessCustomIMCallbackResponseBody) SetMessage(v string) *ProcessCustomIMCallbackResponseBody {
	s.Message = &v
	return s
}

func (s *ProcessCustomIMCallbackResponseBody) SetParams(v []*string) *ProcessCustomIMCallbackResponseBody {
	s.Params = v
	return s
}

func (s *ProcessCustomIMCallbackResponseBody) SetRequestId(v string) *ProcessCustomIMCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProcessCustomIMCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
