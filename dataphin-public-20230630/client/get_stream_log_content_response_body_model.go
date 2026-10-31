// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamLogContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStreamLogContentResponseBody
	GetCode() *string
	SetData(v *GetStreamLogContentResponseBodyData) *GetStreamLogContentResponseBody
	GetData() *GetStreamLogContentResponseBodyData
	SetHttpStatusCode(v int32) *GetStreamLogContentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetStreamLogContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStreamLogContentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStreamLogContentResponseBody
	GetSuccess() *bool
}

type GetStreamLogContentResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The body of the JobManager or TaskManager log file. Pagination information is not included.
	Data *GetStreamLogContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend response exception.
	//
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
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStreamLogContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStreamLogContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetStreamLogContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStreamLogContentResponseBody) GetData() *GetStreamLogContentResponseBodyData {
	return s.Data
}

func (s *GetStreamLogContentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetStreamLogContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStreamLogContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStreamLogContentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStreamLogContentResponseBody) SetCode(v string) *GetStreamLogContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetStreamLogContentResponseBody) SetData(v *GetStreamLogContentResponseBodyData) *GetStreamLogContentResponseBody {
	s.Data = v
	return s
}

func (s *GetStreamLogContentResponseBody) SetHttpStatusCode(v int32) *GetStreamLogContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetStreamLogContentResponseBody) SetMessage(v string) *GetStreamLogContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetStreamLogContentResponseBody) SetRequestId(v string) *GetStreamLogContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStreamLogContentResponseBody) SetSuccess(v bool) *GetStreamLogContentResponseBody {
	s.Success = &v
	return s
}

func (s *GetStreamLogContentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStreamLogContentResponseBodyData struct {
	// The text content of the log file, which may contain line breaks. This parameter may be null if no log content is available.
	//
	// example:
	//
	// 2026-10-15 13:46:26 INFO Starting Flink session cluster
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s GetStreamLogContentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStreamLogContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStreamLogContentResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetStreamLogContentResponseBodyData) SetContent(v string) *GetStreamLogContentResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetStreamLogContentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
