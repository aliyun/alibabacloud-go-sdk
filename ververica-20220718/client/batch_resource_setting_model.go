// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchResourceSetting interface {
	dara.Model
	String() string
	GoString() string
	SetBasicResourceSetting(v *BasicResourceSetting) *BatchResourceSetting
	GetBasicResourceSetting() *BasicResourceSetting
	SetMaxSlot(v int64) *BatchResourceSetting
	GetMaxSlot() *int64
}

type BatchResourceSetting struct {
	BasicResourceSetting *BasicResourceSetting `json:"basicResourceSetting,omitempty" xml:"basicResourceSetting,omitempty"`
	// example:
	//
	// 10
	MaxSlot *int64 `json:"maxSlot,omitempty" xml:"maxSlot,omitempty"`
}

func (s BatchResourceSetting) String() string {
	return dara.Prettify(s)
}

func (s BatchResourceSetting) GoString() string {
	return s.String()
}

func (s *BatchResourceSetting) GetBasicResourceSetting() *BasicResourceSetting {
	return s.BasicResourceSetting
}

func (s *BatchResourceSetting) GetMaxSlot() *int64 {
	return s.MaxSlot
}

func (s *BatchResourceSetting) SetBasicResourceSetting(v *BasicResourceSetting) *BatchResourceSetting {
	s.BasicResourceSetting = v
	return s
}

func (s *BatchResourceSetting) SetMaxSlot(v int64) *BatchResourceSetting {
	s.MaxSlot = &v
	return s
}

func (s *BatchResourceSetting) Validate() error {
	if s.BasicResourceSetting != nil {
		if err := s.BasicResourceSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}
