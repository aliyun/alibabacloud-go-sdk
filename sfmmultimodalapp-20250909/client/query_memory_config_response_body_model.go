// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMemoryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryMemoryConfigResponseBody
	GetCode() *string
	SetData(v *QueryMemoryConfigResponseBodyData) *QueryMemoryConfigResponseBody
	GetData() *QueryMemoryConfigResponseBodyData
	SetHttpStatusCode(v int32) *QueryMemoryConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryMemoryConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryMemoryConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMemoryConfigResponseBody
	GetSuccess() *bool
}

type QueryMemoryConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryMemoryConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B075711B-7857-5BC3-A953-F04B1755EF67
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMemoryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMemoryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMemoryConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryMemoryConfigResponseBody) GetData() *QueryMemoryConfigResponseBodyData {
	return s.Data
}

func (s *QueryMemoryConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryMemoryConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMemoryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMemoryConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMemoryConfigResponseBody) SetCode(v string) *QueryMemoryConfigResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMemoryConfigResponseBody) SetData(v *QueryMemoryConfigResponseBodyData) *QueryMemoryConfigResponseBody {
	s.Data = v
	return s
}

func (s *QueryMemoryConfigResponseBody) SetHttpStatusCode(v int32) *QueryMemoryConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryMemoryConfigResponseBody) SetMessage(v string) *QueryMemoryConfigResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMemoryConfigResponseBody) SetRequestId(v string) *QueryMemoryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMemoryConfigResponseBody) SetSuccess(v bool) *QueryMemoryConfigResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMemoryConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMemoryConfigResponseBodyData struct {
	// example:
	//
	// mm_bfaf7e110b6d4359977d1686a3f8
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// true
	AutoUpdate *bool `json:"AutoUpdate,omitempty" xml:"AutoUpdate,omitempty"`
	// example:
	//
	// 30
	ExpirationTime *int32  `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 0.03
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// example:
	//
	// we4m6373
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// example:
	//
	// llm-956eawe4m6373633
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMemoryConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryMemoryConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMemoryConfigResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *QueryMemoryConfigResponseBodyData) GetAutoUpdate() *bool {
	return s.AutoUpdate
}

func (s *QueryMemoryConfigResponseBodyData) GetExpirationTime() *int32 {
	return s.ExpirationTime
}

func (s *QueryMemoryConfigResponseBodyData) GetPrompt() *string {
	return s.Prompt
}

func (s *QueryMemoryConfigResponseBodyData) GetThreshold() *float64 {
	return s.Threshold
}

func (s *QueryMemoryConfigResponseBodyData) GetTopK() *int32 {
	return s.TopK
}

func (s *QueryMemoryConfigResponseBodyData) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *QueryMemoryConfigResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMemoryConfigResponseBodyData) SetAppId(v string) *QueryMemoryConfigResponseBodyData {
	s.AppId = &v
	return s
}

func (s *QueryMemoryConfigResponseBodyData) SetAutoUpdate(v bool) *QueryMemoryConfigResponseBodyData {
	s.AutoUpdate = &v
	return s
}

func (s *QueryMemoryConfigResponseBodyData) SetExpirationTime(v int32) *QueryMemoryConfigResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *QueryMemoryConfigResponseBodyData) SetPrompt(v string) *QueryMemoryConfigResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *QueryMemoryConfigResponseBodyData) SetThreshold(v float64) *QueryMemoryConfigResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *QueryMemoryConfigResponseBodyData) SetTopK(v int32) *QueryMemoryConfigResponseBodyData {
	s.TopK = &v
	return s
}

func (s *QueryMemoryConfigResponseBodyData) SetUserDefinedId(v string) *QueryMemoryConfigResponseBodyData {
	s.UserDefinedId = &v
	return s
}

func (s *QueryMemoryConfigResponseBodyData) SetWorkspaceId(v string) *QueryMemoryConfigResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMemoryConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
