// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskTemplateViewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTaskTemplateViewsResponseBody
	GetCode() *int32
	SetDetails(v string) *GetTaskTemplateViewsResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetTaskTemplateViewsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTaskTemplateViewsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskTemplateViewsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskTemplateViewsResponseBody
	GetSuccess() *bool
	SetViews(v *GetTaskTemplateViewsResponseBodyViews) *GetTaskTemplateViewsResponseBody
	GetViews() *GetTaskTemplateViewsResponseBodyViews
}

type GetTaskTemplateViewsResponseBody struct {
	// Total amount of data under the current request conditions. This parameter is optional and does not need to be returned by default.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// error code
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Return message of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// is succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// view
	Views *GetTaskTemplateViewsResponseBodyViews `json:"Views,omitempty" xml:"Views,omitempty" type:"Struct"`
}

func (s GetTaskTemplateViewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskTemplateViewsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateViewsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTaskTemplateViewsResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetTaskTemplateViewsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskTemplateViewsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskTemplateViewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskTemplateViewsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskTemplateViewsResponseBody) GetViews() *GetTaskTemplateViewsResponseBodyViews {
	return s.Views
}

func (s *GetTaskTemplateViewsResponseBody) SetCode(v int32) *GetTaskTemplateViewsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) SetDetails(v string) *GetTaskTemplateViewsResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) SetErrorCode(v string) *GetTaskTemplateViewsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) SetMessage(v string) *GetTaskTemplateViewsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) SetRequestId(v string) *GetTaskTemplateViewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) SetSuccess(v bool) *GetTaskTemplateViewsResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) SetViews(v *GetTaskTemplateViewsResponseBodyViews) *GetTaskTemplateViewsResponseBody {
	s.Views = v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) Validate() error {
	if s.Views != nil {
		if err := s.Views.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskTemplateViewsResponseBodyViews struct {
	// List of view plugins
	ViewPlugins []*ViewPlugin `json:"ViewPlugins,omitempty" xml:"ViewPlugins,omitempty" type:"Repeated"`
}

func (s GetTaskTemplateViewsResponseBodyViews) String() string {
	return dara.Prettify(s)
}

func (s GetTaskTemplateViewsResponseBodyViews) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateViewsResponseBodyViews) GetViewPlugins() []*ViewPlugin {
	return s.ViewPlugins
}

func (s *GetTaskTemplateViewsResponseBodyViews) SetViewPlugins(v []*ViewPlugin) *GetTaskTemplateViewsResponseBodyViews {
	s.ViewPlugins = v
	return s
}

func (s *GetTaskTemplateViewsResponseBodyViews) Validate() error {
	if s.ViewPlugins != nil {
		for _, item := range s.ViewPlugins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
