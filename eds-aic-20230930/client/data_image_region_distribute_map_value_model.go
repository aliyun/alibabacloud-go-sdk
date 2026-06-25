// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataImageRegionDistributeMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetDistributeStatus(v string) *DataImageRegionDistributeMapValue
	GetDistributeStatus() *string
	SetProgress(v string) *DataImageRegionDistributeMapValue
	GetProgress() *string
}

type DataImageRegionDistributeMapValue struct {
	// The distribution status.
	//
	// example:
	//
	// AVAILABLE
	DistributeStatus *string `json:"DistributeStatus,omitempty" xml:"DistributeStatus,omitempty"`
	// The distribution progress.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DataImageRegionDistributeMapValue) String() string {
	return dara.Prettify(s)
}

func (s DataImageRegionDistributeMapValue) GoString() string {
	return s.String()
}

func (s *DataImageRegionDistributeMapValue) GetDistributeStatus() *string {
	return s.DistributeStatus
}

func (s *DataImageRegionDistributeMapValue) GetProgress() *string {
	return s.Progress
}

func (s *DataImageRegionDistributeMapValue) SetDistributeStatus(v string) *DataImageRegionDistributeMapValue {
	s.DistributeStatus = &v
	return s
}

func (s *DataImageRegionDistributeMapValue) SetProgress(v string) *DataImageRegionDistributeMapValue {
	s.Progress = &v
	return s
}

func (s *DataImageRegionDistributeMapValue) Validate() error {
	return dara.Validate(s)
}
