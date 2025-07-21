// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowPreviewUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetFlowPreviewUrlResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetFlowPreviewUrlResponseBody
	GetCode() *string
	SetData(v *GetFlowPreviewUrlResponseBodyData) *GetFlowPreviewUrlResponseBody
	GetData() *GetFlowPreviewUrlResponseBodyData
	SetMessage(v string) *GetFlowPreviewUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFlowPreviewUrlResponseBody
	GetRequestId() *string
}

type GetFlowPreviewUrlResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetFlowPreviewUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFlowPreviewUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFlowPreviewUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowPreviewUrlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetFlowPreviewUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFlowPreviewUrlResponseBody) GetData() *GetFlowPreviewUrlResponseBodyData {
	return s.Data
}

func (s *GetFlowPreviewUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFlowPreviewUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFlowPreviewUrlResponseBody) SetAccessDeniedDetail(v string) *GetFlowPreviewUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetFlowPreviewUrlResponseBody) SetCode(v string) *GetFlowPreviewUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetFlowPreviewUrlResponseBody) SetData(v *GetFlowPreviewUrlResponseBodyData) *GetFlowPreviewUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetFlowPreviewUrlResponseBody) SetMessage(v string) *GetFlowPreviewUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetFlowPreviewUrlResponseBody) SetRequestId(v string) *GetFlowPreviewUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFlowPreviewUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFlowPreviewUrlResponseBodyData struct {
	// The Flow ID.
	//
	// example:
	//
	// 6dd31e1b7cc940fc99e293d9952b5b79
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The temporary preview URL.
	//
	// example:
	//
	// https://url
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// The time when the preview URL expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1700617436633
	PreviewUrlExpires *int64 `json:"PreviewUrlExpires,omitempty" xml:"PreviewUrlExpires,omitempty"`
}

func (s GetFlowPreviewUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFlowPreviewUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFlowPreviewUrlResponseBodyData) GetFlowId() *string {
	return s.FlowId
}

func (s *GetFlowPreviewUrlResponseBodyData) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *GetFlowPreviewUrlResponseBodyData) GetPreviewUrlExpires() *int64 {
	return s.PreviewUrlExpires
}

func (s *GetFlowPreviewUrlResponseBodyData) SetFlowId(v string) *GetFlowPreviewUrlResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *GetFlowPreviewUrlResponseBodyData) SetPreviewUrl(v string) *GetFlowPreviewUrlResponseBodyData {
	s.PreviewUrl = &v
	return s
}

func (s *GetFlowPreviewUrlResponseBodyData) SetPreviewUrlExpires(v int64) *GetFlowPreviewUrlResponseBodyData {
	s.PreviewUrlExpires = &v
	return s
}

func (s *GetFlowPreviewUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
