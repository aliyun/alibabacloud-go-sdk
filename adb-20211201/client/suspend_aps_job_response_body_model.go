// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendApsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApsJobId(v string) *SuspendApsJobResponseBody
	GetApsJobId() *string
	SetErrCode(v string) *SuspendApsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SuspendApsJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *SuspendApsJobResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SuspendApsJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendApsJobResponseBody
	GetSuccess() *bool
}

type SuspendApsJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// aps-bj1xxxxxx
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The HTTP status code or the error code.
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
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******-3EEC-******-9F06-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendApsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendApsJobResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendApsJobResponseBody) GetApsJobId() *string {
	return s.ApsJobId
}

func (s *SuspendApsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SuspendApsJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SuspendApsJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SuspendApsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendApsJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendApsJobResponseBody) SetApsJobId(v string) *SuspendApsJobResponseBody {
	s.ApsJobId = &v
	return s
}

func (s *SuspendApsJobResponseBody) SetErrCode(v string) *SuspendApsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SuspendApsJobResponseBody) SetErrMessage(v string) *SuspendApsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SuspendApsJobResponseBody) SetHttpStatusCode(v int32) *SuspendApsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SuspendApsJobResponseBody) SetRequestId(v string) *SuspendApsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendApsJobResponseBody) SetSuccess(v bool) *SuspendApsJobResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendApsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
