// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateComputeResourceResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateComputeResourceResponseBody
	GetRequestId() *string
}

type CreateComputeResourceResponseBody struct {
	// Returns the ID of the created computing resource.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateComputeResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComputeResourceResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateComputeResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComputeResourceResponseBody) SetId(v int64) *CreateComputeResourceResponseBody {
	s.Id = &v
	return s
}

func (s *CreateComputeResourceResponseBody) SetRequestId(v string) *CreateComputeResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComputeResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
