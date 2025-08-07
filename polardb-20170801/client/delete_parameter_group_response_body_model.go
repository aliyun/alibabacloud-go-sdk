// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteParameterGroupResponseBody
	GetRequestId() *string
}

type DeleteParameterGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4F7195E7-5F74-479D-AF59-1B4BF9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteParameterGroupResponseBody) SetRequestId(v string) *DeleteParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
