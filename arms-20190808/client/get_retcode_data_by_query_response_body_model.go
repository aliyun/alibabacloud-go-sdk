// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeDataByQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRetcodeDataByQueryResponseBody
	GetCode() *string
	SetData(v string) *GetRetcodeDataByQueryResponseBody
	GetData() *string
	SetMessage(v string) *GetRetcodeDataByQueryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRetcodeDataByQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRetcodeDataByQueryResponseBody
	GetSuccess() *bool
}

type GetRetcodeDataByQueryResponseBody struct {
	// The HTTP status code returned for the request. Valid values:
	//
	// 	- 2XX: The request was successful.
	//
	// 	- 3XX: A redirection message was returned.
	//
	// 	- 4XX: The request was invalid.
	//
	// 	- 5XX: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The Browser Monitoring data returned.
	//
	// example:
	//
	// {"total":0,"auth":false,"pageSize":20,"completed":true,"page":1,"items":[]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// Internal error. Please try again. Contact the DingTalk service account if the issue                              persists after multiple retries.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2A0CEDF1-06FE-44AC-8E21-21A5BE65****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRetcodeDataByQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeDataByQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GetRetcodeDataByQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRetcodeDataByQueryResponseBody) GetData() *string {
	return s.Data
}

func (s *GetRetcodeDataByQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRetcodeDataByQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRetcodeDataByQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRetcodeDataByQueryResponseBody) SetCode(v string) *GetRetcodeDataByQueryResponseBody {
	s.Code = &v
	return s
}

func (s *GetRetcodeDataByQueryResponseBody) SetData(v string) *GetRetcodeDataByQueryResponseBody {
	s.Data = &v
	return s
}

func (s *GetRetcodeDataByQueryResponseBody) SetMessage(v string) *GetRetcodeDataByQueryResponseBody {
	s.Message = &v
	return s
}

func (s *GetRetcodeDataByQueryResponseBody) SetRequestId(v string) *GetRetcodeDataByQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRetcodeDataByQueryResponseBody) SetSuccess(v bool) *GetRetcodeDataByQueryResponseBody {
	s.Success = &v
	return s
}

func (s *GetRetcodeDataByQueryResponseBody) Validate() error {
	return dara.Validate(s)
}
