// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnchorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateAnchorResponseBody
	GetData() *string
	SetErrorCode(v string) *CreateAnchorResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateAnchorResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateAnchorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAnchorResponseBody
	GetSuccess() *bool
}

type CreateAnchorResponseBody struct {
	// 123456789
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// PARAM_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Failed to proxy flink ui request, message: An error occurred: Invalid UUID string: jobsn
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 10923AA3-F7A1-5EA0-ACCA-D704269EAA78
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateAnchorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAnchorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAnchorResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateAnchorResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateAnchorResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateAnchorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAnchorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAnchorResponseBody) SetData(v string) *CreateAnchorResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAnchorResponseBody) SetErrorCode(v string) *CreateAnchorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAnchorResponseBody) SetErrorMessage(v string) *CreateAnchorResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAnchorResponseBody) SetRequestId(v string) *CreateAnchorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAnchorResponseBody) SetSuccess(v bool) *CreateAnchorResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAnchorResponseBody) Validate() error {
	return dara.Validate(s)
}
