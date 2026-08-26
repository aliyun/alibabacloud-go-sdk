// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOneMetaSqlTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteOneMetaSqlTemplateResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteOneMetaSqlTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteOneMetaSqlTemplateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteOneMetaSqlTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteOneMetaSqlTemplateResponseBody
	GetSuccess() *bool
}

type DeleteOneMetaSqlTemplateResponseBody struct {
	// The response struct.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteOneMetaSqlTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOneMetaSqlTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOneMetaSqlTemplateResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteOneMetaSqlTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteOneMetaSqlTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteOneMetaSqlTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOneMetaSqlTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteOneMetaSqlTemplateResponseBody) SetData(v bool) *DeleteOneMetaSqlTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponseBody) SetErrorCode(v string) *DeleteOneMetaSqlTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponseBody) SetErrorMessage(v string) *DeleteOneMetaSqlTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponseBody) SetRequestId(v string) *DeleteOneMetaSqlTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponseBody) SetSuccess(v bool) *DeleteOneMetaSqlTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
