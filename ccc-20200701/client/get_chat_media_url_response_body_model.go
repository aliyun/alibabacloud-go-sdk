// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatMediaUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetChatMediaUrlResponseBody
	GetCode() *string
	SetData(v string) *GetChatMediaUrlResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetChatMediaUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetChatMediaUrlResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetChatMediaUrlResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetChatMediaUrlResponseBody
	GetRequestId() *string
}

type GetChatMediaUrlResponseBody struct {
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
	// Internal service issue. Detail:.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 9FBA26B0-462B-4D77-B78F-AF35560DBC71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatMediaUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatMediaUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatMediaUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChatMediaUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *GetChatMediaUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetChatMediaUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatMediaUrlResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetChatMediaUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatMediaUrlResponseBody) SetCode(v string) *GetChatMediaUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatMediaUrlResponseBody) SetData(v string) *GetChatMediaUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetChatMediaUrlResponseBody) SetHttpStatusCode(v int32) *GetChatMediaUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetChatMediaUrlResponseBody) SetMessage(v string) *GetChatMediaUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatMediaUrlResponseBody) SetParams(v []*string) *GetChatMediaUrlResponseBody {
	s.Params = v
	return s
}

func (s *GetChatMediaUrlResponseBody) SetRequestId(v string) *GetChatMediaUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatMediaUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
