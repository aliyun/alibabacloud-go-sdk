// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRemindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteRemindResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteRemindResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteRemindResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteRemindResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteRemindResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRemindResponseBody
	GetSuccess() *bool
}

type DeleteRemindResponseBody struct {
	// Indicates whether the custom alert rule is deleted. Valid values: true and false. The value true indicates that the custom alert rule is deleted. The value false indicates that the custom alert rule fails to be deleted.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRemindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRemindResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRemindResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteRemindResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteRemindResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteRemindResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteRemindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRemindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRemindResponseBody) SetData(v bool) *DeleteRemindResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteRemindResponseBody) SetErrorCode(v string) *DeleteRemindResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRemindResponseBody) SetErrorMessage(v string) *DeleteRemindResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRemindResponseBody) SetHttpStatusCode(v int32) *DeleteRemindResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteRemindResponseBody) SetRequestId(v string) *DeleteRemindResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRemindResponseBody) SetSuccess(v bool) *DeleteRemindResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRemindResponseBody) Validate() error {
	return dara.Validate(s)
}
