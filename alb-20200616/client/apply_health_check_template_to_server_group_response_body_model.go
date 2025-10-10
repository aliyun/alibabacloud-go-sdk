// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyHealthCheckTemplateToServerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *ApplyHealthCheckTemplateToServerGroupResponseBody
	GetJobId() *string
	SetRequestId(v string) *ApplyHealthCheckTemplateToServerGroupResponseBody
	GetRequestId() *string
}

type ApplyHealthCheckTemplateToServerGroupResponseBody struct {
	// The asynchronous task ID.
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

func (s ApplyHealthCheckTemplateToServerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyHealthCheckTemplateToServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyHealthCheckTemplateToServerGroupResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *ApplyHealthCheckTemplateToServerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyHealthCheckTemplateToServerGroupResponseBody) SetJobId(v string) *ApplyHealthCheckTemplateToServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupResponseBody) SetRequestId(v string) *ApplyHealthCheckTemplateToServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyHealthCheckTemplateToServerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
