// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestoreJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRestoreJobResponseBody
	GetCode() *string
	SetMessage(v string) *CreateRestoreJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRestoreJobResponseBody
	GetRequestId() *string
	SetRestoreId(v string) *CreateRestoreJobResponseBody
	GetRestoreId() *string
	SetSuccess(v bool) *CreateRestoreJobResponseBody
	GetSuccess() *bool
}

type CreateRestoreJobResponseBody struct {
	// The response code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message. The value "successful" is returned for a success. An error message is returned for a failure.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the restore job.
	//
	// example:
	//
	// r-*********************
	RestoreId *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	// Indicates whether the request was successful.
	//
	//   - true: The request was successful.
	//
	//   - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRestoreJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRestoreJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRestoreJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRestoreJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRestoreJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRestoreJobResponseBody) GetRestoreId() *string {
	return s.RestoreId
}

func (s *CreateRestoreJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRestoreJobResponseBody) SetCode(v string) *CreateRestoreJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRestoreJobResponseBody) SetMessage(v string) *CreateRestoreJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRestoreJobResponseBody) SetRequestId(v string) *CreateRestoreJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRestoreJobResponseBody) SetRestoreId(v string) *CreateRestoreJobResponseBody {
	s.RestoreId = &v
	return s
}

func (s *CreateRestoreJobResponseBody) SetSuccess(v bool) *CreateRestoreJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRestoreJobResponseBody) Validate() error {
	return dara.Validate(s)
}
