// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DeleteAppInstanceResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *DeleteAppInstanceResponseBody
	GetRequestId() *string
}

type DeleteAppInstanceResponseBody struct {
	// The ID of the RDS Supabase instance.
	//
	// example:
	//
	// ra-supabase-8moov5lxba***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeleteAppInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppInstanceResponseBody) SetInstanceName(v string) *DeleteAppInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DeleteAppInstanceResponseBody) SetRequestId(v string) *DeleteAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
