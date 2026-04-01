// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterGroupId(v string) *DeleteParameterGroupResponseBody
	GetParameterGroupId() *string
	SetRequestId(v string) *DeleteParameterGroupResponseBody
	GetRequestId() *string
}

type DeleteParameterGroupResponseBody struct {
	// The ID of the parameter template.
	//
	// example:
	//
	// rpg-gfs****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8AF26036-B254-4212-B8E4-EFBE818B7FD6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParameterGroupResponseBody) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DeleteParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteParameterGroupResponseBody) SetParameterGroupId(v string) *DeleteParameterGroupResponseBody {
	s.ParameterGroupId = &v
	return s
}

func (s *DeleteParameterGroupResponseBody) SetRequestId(v string) *DeleteParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
