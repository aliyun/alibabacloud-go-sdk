// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSampleBase interface {
	dara.Model
	String() string
	GoString() string
	SetFullSampleDeviceIds(v []*FullSampleItem) *SampleBase
	GetFullSampleDeviceIds() []*FullSampleItem
	SetFullSampleUsers(v []*FullSampleItem) *SampleBase
	GetFullSampleUsers() []*FullSampleItem
	SetSampleMethod(v string) *SampleBase
	GetSampleMethod() *string
	SetSampleRate(v float32) *SampleBase
	GetSampleRate() *float32
}

type SampleBase struct {
	FullSampleDeviceIds []*FullSampleItem `json:"FullSampleDeviceIds,omitempty" xml:"FullSampleDeviceIds,omitempty" type:"Repeated"`
	FullSampleUsers     []*FullSampleItem `json:"FullSampleUsers,omitempty" xml:"FullSampleUsers,omitempty" type:"Repeated"`
	// example:
	//
	// event
	SampleMethod *string `json:"SampleMethod,omitempty" xml:"SampleMethod,omitempty"`
	// example:
	//
	// 1
	SampleRate *float32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
}

func (s SampleBase) String() string {
	return dara.Prettify(s)
}

func (s SampleBase) GoString() string {
	return s.String()
}

func (s *SampleBase) GetFullSampleDeviceIds() []*FullSampleItem {
	return s.FullSampleDeviceIds
}

func (s *SampleBase) GetFullSampleUsers() []*FullSampleItem {
	return s.FullSampleUsers
}

func (s *SampleBase) GetSampleMethod() *string {
	return s.SampleMethod
}

func (s *SampleBase) GetSampleRate() *float32 {
	return s.SampleRate
}

func (s *SampleBase) SetFullSampleDeviceIds(v []*FullSampleItem) *SampleBase {
	s.FullSampleDeviceIds = v
	return s
}

func (s *SampleBase) SetFullSampleUsers(v []*FullSampleItem) *SampleBase {
	s.FullSampleUsers = v
	return s
}

func (s *SampleBase) SetSampleMethod(v string) *SampleBase {
	s.SampleMethod = &v
	return s
}

func (s *SampleBase) SetSampleRate(v float32) *SampleBase {
	s.SampleRate = &v
	return s
}

func (s *SampleBase) Validate() error {
	if s.FullSampleDeviceIds != nil {
		for _, item := range s.FullSampleDeviceIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FullSampleUsers != nil {
		for _, item := range s.FullSampleUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
