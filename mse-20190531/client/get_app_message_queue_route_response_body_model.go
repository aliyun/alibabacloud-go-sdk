// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppMessageQueueRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAppMessageQueueRouteResponseBody
	GetCode() *int32
	SetData(v *GetAppMessageQueueRouteResponseBodyData) *GetAppMessageQueueRouteResponseBody
	GetData() *GetAppMessageQueueRouteResponseBodyData
	SetHttpStatusCode(v int32) *GetAppMessageQueueRouteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAppMessageQueueRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAppMessageQueueRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAppMessageQueueRouteResponseBody
	GetSuccess() *bool
}

type GetAppMessageQueueRouteResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data.
	Data *GetAppMessageQueueRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// - If the request is successful, a success message is returned.
	//
	// - If the request fails, the cause of the failure is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B00BCB0-105F-5A2A-B75B-641C8E9B18FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAppMessageQueueRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppMessageQueueRouteResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppMessageQueueRouteResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAppMessageQueueRouteResponseBody) GetData() *GetAppMessageQueueRouteResponseBodyData {
	return s.Data
}

func (s *GetAppMessageQueueRouteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAppMessageQueueRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAppMessageQueueRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppMessageQueueRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAppMessageQueueRouteResponseBody) SetCode(v int32) *GetAppMessageQueueRouteResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppMessageQueueRouteResponseBody) SetData(v *GetAppMessageQueueRouteResponseBodyData) *GetAppMessageQueueRouteResponseBody {
	s.Data = v
	return s
}

func (s *GetAppMessageQueueRouteResponseBody) SetHttpStatusCode(v int32) *GetAppMessageQueueRouteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAppMessageQueueRouteResponseBody) SetMessage(v string) *GetAppMessageQueueRouteResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppMessageQueueRouteResponseBody) SetRequestId(v string) *GetAppMessageQueueRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppMessageQueueRouteResponseBody) SetSuccess(v bool) *GetAppMessageQueueRouteResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppMessageQueueRouteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppMessageQueueRouteResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// hkhon1po62@54e1f42f37cd65a
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Indicates whether the feature is enabled. Valid values:
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The filtering side for message canary release.
	//
	// example:
	//
	// Server
	FilterSide *string `json:"FilterSide,omitempty" xml:"FilterSide,omitempty"`
	// The list of baseline environment labels for MSMQ grayscale rules.
	GrayBaseTags []*string `json:"GrayBaseTags,omitempty" xml:"GrayBaseTags,omitempty" type:"Repeated"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The tags that are not consumed by the untagged environments of the application.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetAppMessageQueueRouteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAppMessageQueueRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppMessageQueueRouteResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetAppMessageQueueRouteResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *GetAppMessageQueueRouteResponseBodyData) GetFilterSide() *string {
	return s.FilterSide
}

func (s *GetAppMessageQueueRouteResponseBodyData) GetGrayBaseTags() []*string {
	return s.GrayBaseTags
}

func (s *GetAppMessageQueueRouteResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetAppMessageQueueRouteResponseBodyData) GetTags() []*string {
	return s.Tags
}

func (s *GetAppMessageQueueRouteResponseBodyData) SetAppId(v string) *GetAppMessageQueueRouteResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetAppMessageQueueRouteResponseBodyData) SetEnable(v bool) *GetAppMessageQueueRouteResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetAppMessageQueueRouteResponseBodyData) SetFilterSide(v string) *GetAppMessageQueueRouteResponseBodyData {
	s.FilterSide = &v
	return s
}

func (s *GetAppMessageQueueRouteResponseBodyData) SetGrayBaseTags(v []*string) *GetAppMessageQueueRouteResponseBodyData {
	s.GrayBaseTags = v
	return s
}

func (s *GetAppMessageQueueRouteResponseBodyData) SetRegion(v string) *GetAppMessageQueueRouteResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetAppMessageQueueRouteResponseBodyData) SetTags(v []*string) *GetAppMessageQueueRouteResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetAppMessageQueueRouteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
