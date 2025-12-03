// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskWarningLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyDiskWarningLineRequest
	GetClusterId() *string
	SetWarningLine(v int32) *ModifyDiskWarningLineRequest
	GetWarningLine() *int32
}

type ModifyDiskWarningLineRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 80
	WarningLine *int32 `json:"WarningLine,omitempty" xml:"WarningLine,omitempty"`
}

func (s ModifyDiskWarningLineRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskWarningLineRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskWarningLineRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyDiskWarningLineRequest) GetWarningLine() *int32 {
	return s.WarningLine
}

func (s *ModifyDiskWarningLineRequest) SetClusterId(v string) *ModifyDiskWarningLineRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyDiskWarningLineRequest) SetWarningLine(v int32) *ModifyDiskWarningLineRequest {
	s.WarningLine = &v
	return s
}

func (s *ModifyDiskWarningLineRequest) Validate() error {
	return dara.Validate(s)
}
