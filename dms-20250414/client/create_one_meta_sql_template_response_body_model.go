// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOneMetaSqlTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OneMetaSqlTemplateView) *CreateOneMetaSqlTemplateResponseBody
	GetData() *OneMetaSqlTemplateView
	SetErrorCode(v string) *CreateOneMetaSqlTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateOneMetaSqlTemplateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateOneMetaSqlTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOneMetaSqlTemplateResponseBody
	GetSuccess() *bool
}

type CreateOneMetaSqlTemplateResponseBody struct {
	// The response struct.
	Data *OneMetaSqlTemplateView `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// ServerUnrecognizedException
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOneMetaSqlTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOneMetaSqlTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOneMetaSqlTemplateResponseBody) GetData() *OneMetaSqlTemplateView {
	return s.Data
}

func (s *CreateOneMetaSqlTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateOneMetaSqlTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateOneMetaSqlTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOneMetaSqlTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOneMetaSqlTemplateResponseBody) SetData(v *OneMetaSqlTemplateView) *CreateOneMetaSqlTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateOneMetaSqlTemplateResponseBody) SetErrorCode(v string) *CreateOneMetaSqlTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateOneMetaSqlTemplateResponseBody) SetErrorMessage(v string) *CreateOneMetaSqlTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateOneMetaSqlTemplateResponseBody) SetRequestId(v string) *CreateOneMetaSqlTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOneMetaSqlTemplateResponseBody) SetSuccess(v bool) *CreateOneMetaSqlTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOneMetaSqlTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
