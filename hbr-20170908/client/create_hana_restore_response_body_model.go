// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHanaRestoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHanaRestoreResponseBody
	GetCode() *string
	SetMessage(v string) *CreateHanaRestoreResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHanaRestoreResponseBody
	GetRequestId() *string
	SetRestoreId(v string) *CreateHanaRestoreResponseBody
	GetRestoreId() *string
	SetSuccess(v bool) *CreateHanaRestoreResponseBody
	GetSuccess() *bool
}

type CreateHanaRestoreResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EEC65C22-2152-5E31-8AD6-D6CBF1BFF49F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the restore job.
	//
	// example:
	//
	// hr-000fb9bz190p1rse6jwv
	RestoreId *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateHanaRestoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHanaRestoreResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHanaRestoreResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHanaRestoreResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHanaRestoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHanaRestoreResponseBody) GetRestoreId() *string {
	return s.RestoreId
}

func (s *CreateHanaRestoreResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateHanaRestoreResponseBody) SetCode(v string) *CreateHanaRestoreResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHanaRestoreResponseBody) SetMessage(v string) *CreateHanaRestoreResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHanaRestoreResponseBody) SetRequestId(v string) *CreateHanaRestoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHanaRestoreResponseBody) SetRestoreId(v string) *CreateHanaRestoreResponseBody {
	s.RestoreId = &v
	return s
}

func (s *CreateHanaRestoreResponseBody) SetSuccess(v bool) *CreateHanaRestoreResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHanaRestoreResponseBody) Validate() error {
	return dara.Validate(s)
}
