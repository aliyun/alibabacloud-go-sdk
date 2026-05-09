// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponents(v []*ModifyAppInstanceResponseBodyComponents) *ModifyAppInstanceResponseBody
	GetComponents() []*ModifyAppInstanceResponseBodyComponents
	SetInstanceName(v string) *ModifyAppInstanceResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *ModifyAppInstanceResponseBody
	GetRequestId() *string
}

type ModifyAppInstanceResponseBody struct {
	Components []*ModifyAppInstanceResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// example:
	//
	// ra-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceResponseBody) GetComponents() []*ModifyAppInstanceResponseBodyComponents {
	return s.Components
}

func (s *ModifyAppInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyAppInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppInstanceResponseBody) SetComponents(v []*ModifyAppInstanceResponseBodyComponents) *ModifyAppInstanceResponseBody {
	s.Components = v
	return s
}

func (s *ModifyAppInstanceResponseBody) SetInstanceName(v string) *ModifyAppInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ModifyAppInstanceResponseBody) SetRequestId(v string) *ModifyAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppInstanceResponseBody) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyAppInstanceResponseBodyComponents struct {
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// supabase
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyAppInstanceResponseBodyComponents) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceResponseBodyComponents) GetStatus() *string {
	return s.Status
}

func (s *ModifyAppInstanceResponseBodyComponents) GetType() *string {
	return s.Type
}

func (s *ModifyAppInstanceResponseBodyComponents) SetStatus(v string) *ModifyAppInstanceResponseBodyComponents {
	s.Status = &v
	return s
}

func (s *ModifyAppInstanceResponseBodyComponents) SetType(v string) *ModifyAppInstanceResponseBodyComponents {
	s.Type = &v
	return s
}

func (s *ModifyAppInstanceResponseBodyComponents) Validate() error {
	return dara.Validate(s)
}
