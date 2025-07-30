// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParamGroupId(v string) *DeleteParameterGroupResponseBody
	GetParamGroupId() *string
	SetRequestId(v string) *DeleteParameterGroupResponseBody
	GetRequestId() *string
}

type DeleteParameterGroupResponseBody struct {
	// The parameter template ID, which is globally unique.
	//
	// example:
	//
	// sys-001*****
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2BE6E619-A657-42E3-AD2D-18F8428A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParameterGroupResponseBody) GetParamGroupId() *string {
	return s.ParamGroupId
}

func (s *DeleteParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteParameterGroupResponseBody) SetParamGroupId(v string) *DeleteParameterGroupResponseBody {
	s.ParamGroupId = &v
	return s
}

func (s *DeleteParameterGroupResponseBody) SetRequestId(v string) *DeleteParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
