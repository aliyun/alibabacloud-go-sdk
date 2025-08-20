// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApsJobId(v string) *DeleteApsJobResponseBody
	GetApsJobId() *string
	SetCode(v string) *DeleteApsJobResponseBody
	GetCode() *string
	SetErrCode(v string) *DeleteApsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteApsJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DeleteApsJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteApsJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteApsJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteApsJobResponseBody
	GetSuccess() *bool
}

type DeleteApsJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// aps-*****
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The HTTP status code or the error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// OK
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ****-****-5D14-AC9F-*********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteApsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApsJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApsJobResponseBody) GetApsJobId() *string {
	return s.ApsJobId
}

func (s *DeleteApsJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteApsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteApsJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteApsJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteApsJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteApsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApsJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteApsJobResponseBody) SetApsJobId(v string) *DeleteApsJobResponseBody {
	s.ApsJobId = &v
	return s
}

func (s *DeleteApsJobResponseBody) SetCode(v string) *DeleteApsJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApsJobResponseBody) SetErrCode(v string) *DeleteApsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteApsJobResponseBody) SetErrMessage(v string) *DeleteApsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteApsJobResponseBody) SetHttpStatusCode(v int32) *DeleteApsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteApsJobResponseBody) SetMessage(v string) *DeleteApsJobResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApsJobResponseBody) SetRequestId(v string) *DeleteApsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApsJobResponseBody) SetSuccess(v bool) *DeleteApsJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteApsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
