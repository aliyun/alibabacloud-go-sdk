// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueEngineVersionByEnvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQueueEngineVersionByEnvResponseBody
	GetCode() *string
	SetData(v []*string) *GetQueueEngineVersionByEnvResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *GetQueueEngineVersionByEnvResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQueueEngineVersionByEnvResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueueEngineVersionByEnvResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQueueEngineVersionByEnvResponseBody
	GetSuccess() *bool
}

type GetQueueEngineVersionByEnvResponseBody struct {
	// example:
	//
	// OK
	Code *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueueEngineVersionByEnvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueueEngineVersionByEnvResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueueEngineVersionByEnvResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQueueEngineVersionByEnvResponseBody) GetData() []*string {
	return s.Data
}

func (s *GetQueueEngineVersionByEnvResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQueueEngineVersionByEnvResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueueEngineVersionByEnvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueueEngineVersionByEnvResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQueueEngineVersionByEnvResponseBody) SetCode(v string) *GetQueueEngineVersionByEnvResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueueEngineVersionByEnvResponseBody) SetData(v []*string) *GetQueueEngineVersionByEnvResponseBody {
	s.Data = v
	return s
}

func (s *GetQueueEngineVersionByEnvResponseBody) SetHttpStatusCode(v int32) *GetQueueEngineVersionByEnvResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQueueEngineVersionByEnvResponseBody) SetMessage(v string) *GetQueueEngineVersionByEnvResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueueEngineVersionByEnvResponseBody) SetRequestId(v string) *GetQueueEngineVersionByEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueueEngineVersionByEnvResponseBody) SetSuccess(v bool) *GetQueueEngineVersionByEnvResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueueEngineVersionByEnvResponseBody) Validate() error {
	return dara.Validate(s)
}
