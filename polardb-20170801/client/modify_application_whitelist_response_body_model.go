// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationWhitelistResponseBody
	GetApplicationId() *string
	SetComponentId(v string) *ModifyApplicationWhitelistResponseBody
	GetComponentId() *string
	SetRequestId(v string) *ModifyApplicationWhitelistResponseBody
	GetRequestId() *string
}

type ModifyApplicationWhitelistResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// pac-**************
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApplicationWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApplicationWhitelistResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationWhitelistResponseBody) GetComponentId() *string {
	return s.ComponentId
}

func (s *ModifyApplicationWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApplicationWhitelistResponseBody) SetApplicationId(v string) *ModifyApplicationWhitelistResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationWhitelistResponseBody) SetComponentId(v string) *ModifyApplicationWhitelistResponseBody {
	s.ComponentId = &v
	return s
}

func (s *ModifyApplicationWhitelistResponseBody) SetRequestId(v string) *ModifyApplicationWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApplicationWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
