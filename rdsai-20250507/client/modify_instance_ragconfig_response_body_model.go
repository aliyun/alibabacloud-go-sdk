// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceRAGConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ModifyInstanceRAGConfigResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *ModifyInstanceRAGConfigResponseBody
	GetRequestId() *string
	SetStatus(v string) *ModifyInstanceRAGConfigResponseBody
	GetStatus() *string
}

type ModifyInstanceRAGConfigResponseBody struct {
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyInstanceRAGConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRAGConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRAGConfigResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceRAGConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceRAGConfigResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ModifyInstanceRAGConfigResponseBody) SetInstanceName(v string) *ModifyInstanceRAGConfigResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceRAGConfigResponseBody) SetRequestId(v string) *ModifyInstanceRAGConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceRAGConfigResponseBody) SetStatus(v string) *ModifyInstanceRAGConfigResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyInstanceRAGConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
