// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHttpApiResponseBody
	GetCode() *string
	SetData(v *HttpApiApiInfo) *GetHttpApiResponseBody
	GetData() *HttpApiApiInfo
	SetMessage(v string) *GetHttpApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHttpApiResponseBody
	GetRequestId() *string
}

type GetHttpApiResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// API information.
	Data *HttpApiApiInfo `json:"data,omitempty" xml:"data,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 8FA9BB94-915B-5299-A694-49FCC7F5DD00
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHttpApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHttpApiResponseBody) GetData() *HttpApiApiInfo {
	return s.Data
}

func (s *GetHttpApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHttpApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHttpApiResponseBody) SetCode(v string) *GetHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *GetHttpApiResponseBody) SetData(v *HttpApiApiInfo) *GetHttpApiResponseBody {
	s.Data = v
	return s
}

func (s *GetHttpApiResponseBody) SetMessage(v string) *GetHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *GetHttpApiResponseBody) SetRequestId(v string) *GetHttpApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHttpApiResponseBody) Validate() error {
	return dara.Validate(s)
}
