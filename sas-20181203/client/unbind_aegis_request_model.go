// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAegisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuids(v string) *UnbindAegisRequest
	GetUuids() *string
}

type UnbindAegisRequest struct {
	// The UUID of the server that you want to unbind. Separate multiple UUIDs with commas (,).
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4fe8e1cd-3c37-4851-b9de-124da32c****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s UnbindAegisRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindAegisRequest) GoString() string {
	return s.String()
}

func (s *UnbindAegisRequest) GetUuids() *string {
	return s.Uuids
}

func (s *UnbindAegisRequest) SetUuids(v string) *UnbindAegisRequest {
	s.Uuids = &v
	return s
}

func (s *UnbindAegisRequest) Validate() error {
	return dara.Validate(s)
}
