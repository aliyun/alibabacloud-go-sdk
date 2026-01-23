// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityWatchTaskLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityWatchTaskLogResponseBody
	GetCode() *string
	SetData(v string) *GetQualityWatchTaskLogResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetQualityWatchTaskLogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualityWatchTaskLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityWatchTaskLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityWatchTaskLogResponseBody
	GetSuccess() *bool
}

type GetQualityWatchTaskLogResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// test
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityWatchTaskLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchTaskLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityWatchTaskLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityWatchTaskLogResponseBody) GetData() *string {
	return s.Data
}

func (s *GetQualityWatchTaskLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityWatchTaskLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityWatchTaskLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityWatchTaskLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityWatchTaskLogResponseBody) SetCode(v string) *GetQualityWatchTaskLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityWatchTaskLogResponseBody) SetData(v string) *GetQualityWatchTaskLogResponseBody {
	s.Data = &v
	return s
}

func (s *GetQualityWatchTaskLogResponseBody) SetHttpStatusCode(v int32) *GetQualityWatchTaskLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityWatchTaskLogResponseBody) SetMessage(v string) *GetQualityWatchTaskLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityWatchTaskLogResponseBody) SetRequestId(v string) *GetQualityWatchTaskLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityWatchTaskLogResponseBody) SetSuccess(v bool) *GetQualityWatchTaskLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityWatchTaskLogResponseBody) Validate() error {
	return dara.Validate(s)
}
