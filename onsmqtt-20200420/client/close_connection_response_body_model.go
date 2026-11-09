// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CloseConnectionResponseBody
	GetCode() *int32
	SetMessage(v string) *CloseConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloseConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CloseConnectionResponseBody
	GetSuccess() *bool
}

type CloseConnectionResponseBody struct {
	// The return code of the API operation. A value of 200 indicates success. Other values indicate error codes. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result message of the API call.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82B9E503-F4A1-4F30-976F-C6999FF9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. A value of true indicates success, and a value of false indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CloseConnectionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CloseConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloseConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CloseConnectionResponseBody) SetCode(v int32) *CloseConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *CloseConnectionResponseBody) SetMessage(v string) *CloseConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *CloseConnectionResponseBody) SetRequestId(v string) *CloseConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseConnectionResponseBody) SetSuccess(v bool) *CloseConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *CloseConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
