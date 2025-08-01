// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEngineNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteEngineNamespaceResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *DeleteEngineNamespaceResponseBody
	GetHttpCode() *string
	SetMessage(v string) *DeleteEngineNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEngineNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEngineNamespaceResponseBody
	GetSuccess() *bool
}

type DeleteEngineNamespaceResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A73AC37C-C617-4E3A-8049-372CF49C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEngineNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEngineNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEngineNamespaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteEngineNamespaceResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *DeleteEngineNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEngineNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEngineNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEngineNamespaceResponseBody) SetErrorCode(v string) *DeleteEngineNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteEngineNamespaceResponseBody) SetHttpCode(v string) *DeleteEngineNamespaceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteEngineNamespaceResponseBody) SetMessage(v string) *DeleteEngineNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEngineNamespaceResponseBody) SetRequestId(v string) *DeleteEngineNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEngineNamespaceResponseBody) SetSuccess(v bool) *DeleteEngineNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEngineNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
