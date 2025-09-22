// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAuthConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ModifyInstanceAuthConfigResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *ModifyInstanceAuthConfigResponseBody
	GetRequestId() *string
}

type ModifyInstanceAuthConfigResponseBody struct {
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAuthConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAuthConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAuthConfigResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceAuthConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceAuthConfigResponseBody) SetInstanceName(v string) *ModifyInstanceAuthConfigResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceAuthConfigResponseBody) SetRequestId(v string) *ModifyInstanceAuthConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceAuthConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
