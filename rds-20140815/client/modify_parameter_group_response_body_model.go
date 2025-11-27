// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterGroupId(v string) *ModifyParameterGroupResponseBody
	GetParameterGroupId() *string
	SetRequestId(v string) *ModifyParameterGroupResponseBody
	GetRequestId() *string
}

type ModifyParameterGroupResponseBody struct {
	// The parameter template ID.
	//
	// example:
	//
	// rpg-13ppdh****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 857DC00B-7B85-4853-8B27-AD65EB618BC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParameterGroupResponseBody) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyParameterGroupResponseBody) SetParameterGroupId(v string) *ModifyParameterGroupResponseBody {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyParameterGroupResponseBody) SetRequestId(v string) *ModifyParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
