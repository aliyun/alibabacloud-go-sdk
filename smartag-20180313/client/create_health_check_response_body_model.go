// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHealthCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateHealthCheckResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *CreateHealthCheckResponseBody
	GetRequestId() *string
}

type CreateHealthCheckResponseBody struct {
	// The ID of the health check.
	//
	// example:
	//
	// hc-rrqoucina3gmpn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E38E950D-28A4-4C41-9428-A8908EC6AE5C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHealthCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHealthCheckResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateHealthCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHealthCheckResponseBody) SetInstanceId(v string) *CreateHealthCheckResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateHealthCheckResponseBody) SetRequestId(v string) *CreateHealthCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHealthCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
