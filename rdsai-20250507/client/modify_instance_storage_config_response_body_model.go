// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceStorageConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ModifyInstanceStorageConfigResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *ModifyInstanceStorageConfigResponseBody
	GetRequestId() *string
}

type ModifyInstanceStorageConfigResponseBody struct {
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceStorageConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceStorageConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageConfigResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceStorageConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceStorageConfigResponseBody) SetInstanceName(v string) *ModifyInstanceStorageConfigResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceStorageConfigResponseBody) SetRequestId(v string) *ModifyInstanceStorageConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceStorageConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
