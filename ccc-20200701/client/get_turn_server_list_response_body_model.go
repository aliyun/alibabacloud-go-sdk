// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTurnServerListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTurnServerListResponseBody
	GetCode() *string
	SetData(v string) *GetTurnServerListResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetTurnServerListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTurnServerListResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetTurnServerListResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetTurnServerListResponseBody
	GetRequestId() *string
}

type GetTurnServerListResponseBody struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTurnServerListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTurnServerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTurnServerListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTurnServerListResponseBody) GetData() *string {
	return s.Data
}

func (s *GetTurnServerListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTurnServerListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTurnServerListResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetTurnServerListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTurnServerListResponseBody) SetCode(v string) *GetTurnServerListResponseBody {
	s.Code = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetData(v string) *GetTurnServerListResponseBody {
	s.Data = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetHttpStatusCode(v int32) *GetTurnServerListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetMessage(v string) *GetTurnServerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetParams(v []*string) *GetTurnServerListResponseBody {
	s.Params = v
	return s
}

func (s *GetTurnServerListResponseBody) SetRequestId(v string) *GetTurnServerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTurnServerListResponseBody) Validate() error {
	return dara.Validate(s)
}
