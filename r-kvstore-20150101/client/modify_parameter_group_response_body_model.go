// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParamGroupId(v string) *ModifyParameterGroupResponseBody
	GetParamGroupId() *string
	SetRequestId(v string) *ModifyParameterGroupResponseBody
	GetRequestId() *string
}

type ModifyParameterGroupResponseBody struct {
	// The parameter template ID.
	//
	// example:
	//
	// testGroupName
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AD425AD3-CC7B-4EE2-A5CB-2F61BA73****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParameterGroupResponseBody) GetParamGroupId() *string {
	return s.ParamGroupId
}

func (s *ModifyParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyParameterGroupResponseBody) SetParamGroupId(v string) *ModifyParameterGroupResponseBody {
	s.ParamGroupId = &v
	return s
}

func (s *ModifyParameterGroupResponseBody) SetRequestId(v string) *ModifyParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
