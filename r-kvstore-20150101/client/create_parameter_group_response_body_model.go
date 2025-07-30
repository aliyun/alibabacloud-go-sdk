// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParamGroupId(v string) *CreateParameterGroupResponseBody
	GetParamGroupId() *string
	SetRequestId(v string) *CreateParameterGroupResponseBody
	GetRequestId() *string
}

type CreateParameterGroupResponseBody struct {
	// The parameter template ID.
	//
	// example:
	//
	// g-51ii2ienn0dg0xi8****
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 62DA5BE5-F9C9-527C-ACCB-4D783C297A3A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParameterGroupResponseBody) GetParamGroupId() *string {
	return s.ParamGroupId
}

func (s *CreateParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateParameterGroupResponseBody) SetParamGroupId(v string) *CreateParameterGroupResponseBody {
	s.ParamGroupId = &v
	return s
}

func (s *CreateParameterGroupResponseBody) SetRequestId(v string) *CreateParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
