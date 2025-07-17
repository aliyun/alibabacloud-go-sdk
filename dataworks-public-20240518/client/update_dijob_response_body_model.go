// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDIJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDIJobResponseBody
	GetSuccess() *bool
}

type UpdateDIJobResponseBody struct {
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// AAC30B35-820D-5F3E-A42C-E96BB6379325
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

func (s UpdateDIJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDIJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDIJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDIJobResponseBody) SetRequestId(v string) *UpdateDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDIJobResponseBody) SetSuccess(v bool) *UpdateDIJobResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDIJobResponseBody) Validate() error {
	return dara.Validate(s)
}
