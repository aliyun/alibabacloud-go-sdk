// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServerGroupServersAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateServerGroupServersAttributeResponseBody
	GetJobId() *string
	SetRequestId(v string) *UpdateServerGroupServersAttributeResponseBody
	GetRequestId() *string
}

type UpdateServerGroupServersAttributeResponseBody struct {
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

func (s UpdateServerGroupServersAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerGroupServersAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpdateServerGroupServersAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetJobId(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetRequestId(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
