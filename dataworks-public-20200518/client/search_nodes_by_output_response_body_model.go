// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchNodesByOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *SearchNodesByOutputResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v string) *SearchNodesByOutputResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SearchNodesByOutputResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *SearchNodesByOutputResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SearchNodesByOutputResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchNodesByOutputResponseBody
	GetSuccess() *bool
}

type SearchNodesByOutputResponseBody struct {
	// The map returned. The key in the map indicates an output name, and the value in the map indicates the information about the node that generates the output.
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// SDFSDFSDF-asdfDFSDF-SDFSDf-SDfSFD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchNodesByOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchNodesByOutputResponseBody) GoString() string {
	return s.String()
}

func (s *SearchNodesByOutputResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *SearchNodesByOutputResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchNodesByOutputResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SearchNodesByOutputResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SearchNodesByOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchNodesByOutputResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchNodesByOutputResponseBody) SetData(v map[string]interface{}) *SearchNodesByOutputResponseBody {
	s.Data = v
	return s
}

func (s *SearchNodesByOutputResponseBody) SetErrorCode(v string) *SearchNodesByOutputResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchNodesByOutputResponseBody) SetErrorMessage(v string) *SearchNodesByOutputResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchNodesByOutputResponseBody) SetHttpStatusCode(v int32) *SearchNodesByOutputResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchNodesByOutputResponseBody) SetRequestId(v string) *SearchNodesByOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchNodesByOutputResponseBody) SetSuccess(v bool) *SearchNodesByOutputResponseBody {
	s.Success = &v
	return s
}

func (s *SearchNodesByOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
