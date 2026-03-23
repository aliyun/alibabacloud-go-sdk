// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterGroupId(v string) *CreateParameterGroupResponseBody
	GetParameterGroupId() *string
	SetRequestId(v string) *CreateParameterGroupResponseBody
	GetRequestId() *string
}

type CreateParameterGroupResponseBody struct {
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParameterGroupResponseBody) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *CreateParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateParameterGroupResponseBody) SetParameterGroupId(v string) *CreateParameterGroupResponseBody {
	s.ParameterGroupId = &v
	return s
}

func (s *CreateParameterGroupResponseBody) SetRequestId(v string) *CreateParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
