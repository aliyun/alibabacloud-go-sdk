// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServerGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateServerGroupAttributeResponseBody
	GetJobId() *string
	SetRequestId(v string) *UpdateServerGroupAttributeResponseBody
	GetRequestId() *string
}

type UpdateServerGroupAttributeResponseBody struct {
	// The ID of the asynchronous job.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServerGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpdateServerGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServerGroupAttributeResponseBody) SetJobId(v string) *UpdateServerGroupAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) SetRequestId(v string) *UpdateServerGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
