// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchMemoryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PatchMemoryConfigResponseBody
	GetCode() *string
	SetData(v *PatchMemoryConfigResponseBodyData) *PatchMemoryConfigResponseBody
	GetData() *PatchMemoryConfigResponseBodyData
	SetHttpStatusCode(v int32) *PatchMemoryConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PatchMemoryConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *PatchMemoryConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PatchMemoryConfigResponseBody
	GetSuccess() *bool
}

type PatchMemoryConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *PatchMemoryConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 82296D89-6895-574B-8AA1-64959957CB41
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PatchMemoryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PatchMemoryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PatchMemoryConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *PatchMemoryConfigResponseBody) GetData() *PatchMemoryConfigResponseBodyData {
	return s.Data
}

func (s *PatchMemoryConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PatchMemoryConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PatchMemoryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PatchMemoryConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PatchMemoryConfigResponseBody) SetCode(v string) *PatchMemoryConfigResponseBody {
	s.Code = &v
	return s
}

func (s *PatchMemoryConfigResponseBody) SetData(v *PatchMemoryConfigResponseBodyData) *PatchMemoryConfigResponseBody {
	s.Data = v
	return s
}

func (s *PatchMemoryConfigResponseBody) SetHttpStatusCode(v int32) *PatchMemoryConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PatchMemoryConfigResponseBody) SetMessage(v string) *PatchMemoryConfigResponseBody {
	s.Message = &v
	return s
}

func (s *PatchMemoryConfigResponseBody) SetRequestId(v string) *PatchMemoryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PatchMemoryConfigResponseBody) SetSuccess(v bool) *PatchMemoryConfigResponseBody {
	s.Success = &v
	return s
}

func (s *PatchMemoryConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PatchMemoryConfigResponseBodyData struct {
	// example:
	//
	// 1020d93c-9f17-4a39-9fe8-50b1fb1198d7
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
	// 3
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// example:
	//
	// 4a39-9fe8-50b1fb1
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// example:
	//
	// llm-ycc5m37pi2if9nwu
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PatchMemoryConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PatchMemoryConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *PatchMemoryConfigResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *PatchMemoryConfigResponseBodyData) GetAutoUpdate() *bool {
	return s.AutoUpdate
}

func (s *PatchMemoryConfigResponseBodyData) GetExpirationTime() *int32 {
	return s.ExpirationTime
}

func (s *PatchMemoryConfigResponseBodyData) GetPrompt() *string {
	return s.Prompt
}

func (s *PatchMemoryConfigResponseBodyData) GetThreshold() *float64 {
	return s.Threshold
}

func (s *PatchMemoryConfigResponseBodyData) GetTopK() *int32 {
	return s.TopK
}

func (s *PatchMemoryConfigResponseBodyData) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *PatchMemoryConfigResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PatchMemoryConfigResponseBodyData) SetAppId(v string) *PatchMemoryConfigResponseBodyData {
	s.AppId = &v
	return s
}

func (s *PatchMemoryConfigResponseBodyData) SetAutoUpdate(v bool) *PatchMemoryConfigResponseBodyData {
	s.AutoUpdate = &v
	return s
}

func (s *PatchMemoryConfigResponseBodyData) SetExpirationTime(v int32) *PatchMemoryConfigResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *PatchMemoryConfigResponseBodyData) SetPrompt(v string) *PatchMemoryConfigResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *PatchMemoryConfigResponseBodyData) SetThreshold(v float64) *PatchMemoryConfigResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *PatchMemoryConfigResponseBodyData) SetTopK(v int32) *PatchMemoryConfigResponseBodyData {
	s.TopK = &v
	return s
}

func (s *PatchMemoryConfigResponseBodyData) SetUserDefinedId(v string) *PatchMemoryConfigResponseBodyData {
	s.UserDefinedId = &v
	return s
}

func (s *PatchMemoryConfigResponseBodyData) SetWorkspaceId(v string) *PatchMemoryConfigResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *PatchMemoryConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
