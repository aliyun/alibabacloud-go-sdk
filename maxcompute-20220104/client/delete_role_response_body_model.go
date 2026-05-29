// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeleteRoleResponseBody
	GetData() *string
	SetRequestId(v string) *DeleteRoleResponseBody
	GetRequestId() *string
}

type DeleteRoleResponseBody struct {
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 0adb901117579891946416405d0409
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoleResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRoleResponseBody) SetData(v string) *DeleteRoleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteRoleResponseBody) SetRequestId(v string) *DeleteRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
