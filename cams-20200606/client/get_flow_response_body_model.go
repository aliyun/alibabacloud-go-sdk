// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetFlowResponseBody
	GetCode() *string
	SetData(v *GetFlowResponseBodyData) *GetFlowResponseBody
	GetData() *GetFlowResponseBodyData
	SetMessage(v string) *GetFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFlowResponseBody
	GetRequestId() *string
}

type GetFlowResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFlowResponseBody) GetData() *GetFlowResponseBodyData {
	return s.Data
}

func (s *GetFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFlowResponseBody) SetAccessDeniedDetail(v string) *GetFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetFlowResponseBody) SetCode(v string) *GetFlowResponseBody {
	s.Code = &v
	return s
}

func (s *GetFlowResponseBody) SetData(v *GetFlowResponseBodyData) *GetFlowResponseBody {
	s.Data = v
	return s
}

func (s *GetFlowResponseBody) SetMessage(v string) *GetFlowResponseBody {
	s.Message = &v
	return s
}

func (s *GetFlowResponseBody) SetRequestId(v string) *GetFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFlowResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFlowResponseBodyData struct {
	// The categories of the Flow.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The version number of the API.
	//
	// example:
	//
	// 3.0
	DataApiVersion *string `json:"DataApiVersion,omitempty" xml:"DataApiVersion,omitempty"`
	// example:
	//
	// http://abc.com
	EndpointUri *string `json:"EndpointUri,omitempty" xml:"EndpointUri,omitempty"`
	// The Flow ID.
	//
	// example:
	//
	// flow_id_arms
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The Flow name.
	//
	// example:
	//
	// dnjn
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The JSON version.
	//
	// example:
	//
	// 2.1
	JSONVersion *string `json:"JSONVersion,omitempty" xml:"JSONVersion,omitempty"`
	// The temporary preview URL.
	//
	// example:
	//
	// https://pre-url
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// The time when the preview URL expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1700617436633
	PreviewUrlExpires *int64 `json:"PreviewUrlExpires,omitempty" xml:"PreviewUrlExpires,omitempty"`
	// The state of the Flow.
	//
	// Valid values:
	//
	// 	- PUBLISHED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DRAFT
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DEPRECATED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// DRAFT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFlowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFlowResponseBodyData) GetCategories() []*string {
	return s.Categories
}

func (s *GetFlowResponseBodyData) GetDataApiVersion() *string {
	return s.DataApiVersion
}

func (s *GetFlowResponseBodyData) GetEndpointUri() *string {
	return s.EndpointUri
}

func (s *GetFlowResponseBodyData) GetFlowId() *string {
	return s.FlowId
}

func (s *GetFlowResponseBodyData) GetFlowName() *string {
	return s.FlowName
}

func (s *GetFlowResponseBodyData) GetJSONVersion() *string {
	return s.JSONVersion
}

func (s *GetFlowResponseBodyData) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *GetFlowResponseBodyData) GetPreviewUrlExpires() *int64 {
	return s.PreviewUrlExpires
}

func (s *GetFlowResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetFlowResponseBodyData) SetCategories(v []*string) *GetFlowResponseBodyData {
	s.Categories = v
	return s
}

func (s *GetFlowResponseBodyData) SetDataApiVersion(v string) *GetFlowResponseBodyData {
	s.DataApiVersion = &v
	return s
}

func (s *GetFlowResponseBodyData) SetEndpointUri(v string) *GetFlowResponseBodyData {
	s.EndpointUri = &v
	return s
}

func (s *GetFlowResponseBodyData) SetFlowId(v string) *GetFlowResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *GetFlowResponseBodyData) SetFlowName(v string) *GetFlowResponseBodyData {
	s.FlowName = &v
	return s
}

func (s *GetFlowResponseBodyData) SetJSONVersion(v string) *GetFlowResponseBodyData {
	s.JSONVersion = &v
	return s
}

func (s *GetFlowResponseBodyData) SetPreviewUrl(v string) *GetFlowResponseBodyData {
	s.PreviewUrl = &v
	return s
}

func (s *GetFlowResponseBodyData) SetPreviewUrlExpires(v int64) *GetFlowResponseBodyData {
	s.PreviewUrlExpires = &v
	return s
}

func (s *GetFlowResponseBodyData) SetStatus(v string) *GetFlowResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetFlowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
