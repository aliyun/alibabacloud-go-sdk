// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTemplateViewResponseBody
	GetCode() *int32
	SetDetails(v string) *GetTemplateViewResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetTemplateViewResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTemplateViewResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTemplateViewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTemplateViewResponseBody
	GetSuccess() *bool
	SetViewConfigs(v *GetTemplateViewResponseBodyViewConfigs) *GetTemplateViewResponseBody
	GetViewConfigs() *GetTemplateViewResponseBodyViewConfigs
}

type GetTemplateViewResponseBody struct {
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
	// ""
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
	// Data display configuration
	ViewConfigs *GetTemplateViewResponseBodyViewConfigs `json:"ViewConfigs,omitempty" xml:"ViewConfigs,omitempty" type:"Struct"`
}

func (s GetTemplateViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateViewResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateViewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTemplateViewResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetTemplateViewResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTemplateViewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTemplateViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateViewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTemplateViewResponseBody) GetViewConfigs() *GetTemplateViewResponseBodyViewConfigs {
	return s.ViewConfigs
}

func (s *GetTemplateViewResponseBody) SetCode(v int32) *GetTemplateViewResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemplateViewResponseBody) SetDetails(v string) *GetTemplateViewResponseBody {
	s.Details = &v
	return s
}

func (s *GetTemplateViewResponseBody) SetErrorCode(v string) *GetTemplateViewResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTemplateViewResponseBody) SetMessage(v string) *GetTemplateViewResponseBody {
	s.Message = &v
	return s
}

func (s *GetTemplateViewResponseBody) SetRequestId(v string) *GetTemplateViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateViewResponseBody) SetSuccess(v bool) *GetTemplateViewResponseBody {
	s.Success = &v
	return s
}

func (s *GetTemplateViewResponseBody) SetViewConfigs(v *GetTemplateViewResponseBodyViewConfigs) *GetTemplateViewResponseBody {
	s.ViewConfigs = v
	return s
}

func (s *GetTemplateViewResponseBody) Validate() error {
	if s.ViewConfigs != nil {
		if err := s.ViewConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTemplateViewResponseBodyViewConfigs struct {
	// List of data display configuration plugins
	ViewPlugins []*ViewPlugin `json:"ViewPlugins,omitempty" xml:"ViewPlugins,omitempty" type:"Repeated"`
}

func (s GetTemplateViewResponseBodyViewConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateViewResponseBodyViewConfigs) GoString() string {
	return s.String()
}

func (s *GetTemplateViewResponseBodyViewConfigs) GetViewPlugins() []*ViewPlugin {
	return s.ViewPlugins
}

func (s *GetTemplateViewResponseBodyViewConfigs) SetViewPlugins(v []*ViewPlugin) *GetTemplateViewResponseBodyViewConfigs {
	s.ViewPlugins = v
	return s
}

func (s *GetTemplateViewResponseBodyViewConfigs) Validate() error {
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
