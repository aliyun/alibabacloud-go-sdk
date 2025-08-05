// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRestoreJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelRestoreJobResponseBody
	GetCode() *string
	SetMessage(v string) *CancelRestoreJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelRestoreJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelRestoreJobResponseBody
	GetSuccess() *bool
}

type CancelRestoreJobResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
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
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelRestoreJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelRestoreJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRestoreJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelRestoreJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelRestoreJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelRestoreJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelRestoreJobResponseBody) SetCode(v string) *CancelRestoreJobResponseBody {
	s.Code = &v
	return s
}

func (s *CancelRestoreJobResponseBody) SetMessage(v string) *CancelRestoreJobResponseBody {
	s.Message = &v
	return s
}

func (s *CancelRestoreJobResponseBody) SetRequestId(v string) *CancelRestoreJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelRestoreJobResponseBody) SetSuccess(v bool) *CancelRestoreJobResponseBody {
	s.Success = &v
	return s
}

func (s *CancelRestoreJobResponseBody) Validate() error {
	return dara.Validate(s)
}
