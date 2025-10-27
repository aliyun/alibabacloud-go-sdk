// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateDnsEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceId(v string) *CreatePrivateDnsEndpointResponseBody
	GetAccessInstanceId() *string
	SetRequestId(v string) *CreatePrivateDnsEndpointResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreatePrivateDnsEndpointResponseBody
	GetTaskId() *string
}

type CreatePrivateDnsEndpointResponseBody struct {
	// example:
	//
	// pd-12345
	AccessInstanceId *string `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 850A84D6************00090125EEB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 132
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreatePrivateDnsEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateDnsEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrivateDnsEndpointResponseBody) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *CreatePrivateDnsEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrivateDnsEndpointResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreatePrivateDnsEndpointResponseBody) SetAccessInstanceId(v string) *CreatePrivateDnsEndpointResponseBody {
	s.AccessInstanceId = &v
	return s
}

func (s *CreatePrivateDnsEndpointResponseBody) SetRequestId(v string) *CreatePrivateDnsEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrivateDnsEndpointResponseBody) SetTaskId(v string) *CreatePrivateDnsEndpointResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreatePrivateDnsEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
