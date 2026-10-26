// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceInspectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOk(v bool) *DeleteInstanceInspectionResponseBody
	GetOk() *bool
}

type DeleteInstanceInspectionResponseBody struct {
	// example:
	//
	// true
	Ok *bool `json:"ok,omitempty" xml:"ok,omitempty"`
}

func (s DeleteInstanceInspectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceInspectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceInspectionResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *DeleteInstanceInspectionResponseBody) SetOk(v bool) *DeleteInstanceInspectionResponseBody {
	s.Ok = &v
	return s
}

func (s *DeleteInstanceInspectionResponseBody) Validate() error {
	return dara.Validate(s)
}
