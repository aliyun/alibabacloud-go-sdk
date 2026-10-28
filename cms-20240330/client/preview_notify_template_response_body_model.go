// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewNotifyTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PreviewNotifyTemplateResponseBodyData) *PreviewNotifyTemplateResponseBody
	GetData() *PreviewNotifyTemplateResponseBodyData
	SetMsg(v string) *PreviewNotifyTemplateResponseBody
	GetMsg() *string
	SetRequestId(v string) *PreviewNotifyTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PreviewNotifyTemplateResponseBody
	GetSuccess() *bool
}

type PreviewNotifyTemplateResponseBody struct {
	// The rendering result object, which contains two mappings: success for successfully rendered items and failed for items that failed to render.
	Data *PreviewNotifyTemplateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The processing result message. The value is success if the request is successful, or an error message if the request fails.
	//
	// example:
	//
	// success
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 452faed9-ab3b-4340-8a92-bc7d9b98f17a
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PreviewNotifyTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewNotifyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewNotifyTemplateResponseBody) GetData() *PreviewNotifyTemplateResponseBodyData {
	return s.Data
}

func (s *PreviewNotifyTemplateResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *PreviewNotifyTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewNotifyTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PreviewNotifyTemplateResponseBody) SetData(v *PreviewNotifyTemplateResponseBodyData) *PreviewNotifyTemplateResponseBody {
	s.Data = v
	return s
}

func (s *PreviewNotifyTemplateResponseBody) SetMsg(v string) *PreviewNotifyTemplateResponseBody {
	s.Msg = &v
	return s
}

func (s *PreviewNotifyTemplateResponseBody) SetRequestId(v string) *PreviewNotifyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewNotifyTemplateResponseBody) SetSuccess(v bool) *PreviewNotifyTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *PreviewNotifyTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewNotifyTemplateResponseBodyData struct {
	// The mapping of failed rendering results. The key is the template identifier, and the value is the cause of the failure, such as a template syntax error or a reference to an undefined variable.
	Failed map[string]interface{} `json:"failed,omitempty" xml:"failed,omitempty"`
	// The mapping of successful rendering results. The key is the template identifier, which corresponds to templates[].key in the request, and the value is the rendered string content.
	Success map[string]interface{} `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PreviewNotifyTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PreviewNotifyTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *PreviewNotifyTemplateResponseBodyData) GetFailed() map[string]interface{} {
	return s.Failed
}

func (s *PreviewNotifyTemplateResponseBodyData) GetSuccess() map[string]interface{} {
	return s.Success
}

func (s *PreviewNotifyTemplateResponseBodyData) SetFailed(v map[string]interface{}) *PreviewNotifyTemplateResponseBodyData {
	s.Failed = v
	return s
}

func (s *PreviewNotifyTemplateResponseBodyData) SetSuccess(v map[string]interface{}) *PreviewNotifyTemplateResponseBodyData {
	s.Success = v
	return s
}

func (s *PreviewNotifyTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
