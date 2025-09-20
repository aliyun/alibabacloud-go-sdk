// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPointInt64 interface {
	dara.Model
	String() string
	GoString() string
	SetX(v int64) *PointInt64
	GetX() *int64
	SetY(v int64) *PointInt64
	GetY() *int64
}

type PointInt64 struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s PointInt64) String() string {
	return dara.Prettify(s)
}

func (s PointInt64) GoString() string {
	return s.String()
}

func (s *PointInt64) GetX() *int64 {
	return s.X
}

func (s *PointInt64) GetY() *int64 {
	return s.Y
}

func (s *PointInt64) SetX(v int64) *PointInt64 {
	s.X = &v
	return s
}

func (s *PointInt64) SetY(v int64) *PointInt64 {
	s.Y = &v
	return s
}

func (s *PointInt64) Validate() error {
	return dara.Validate(s)
}
