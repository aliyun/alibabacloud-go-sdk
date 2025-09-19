// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoDelConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDays(v int32) *ModifyAutoDelConfigRequest
	GetDays() *int32
}

type ModifyAutoDelConfigRequest struct {
	// The number of days after which a detected vulnerability is automatically deleted. Unit: days. Valid values:
	//
	// 	- 7
	//
	// 	- 30
	//
	// 	- 90
	//
	// example:
	//
	// 30
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
}

func (s ModifyAutoDelConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoDelConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoDelConfigRequest) GetDays() *int32 {
	return s.Days
}

func (s *ModifyAutoDelConfigRequest) SetDays(v int32) *ModifyAutoDelConfigRequest {
	s.Days = &v
	return s
}

func (s *ModifyAutoDelConfigRequest) Validate() error {
	return dara.Validate(s)
}
