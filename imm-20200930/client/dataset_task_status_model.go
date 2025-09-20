// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetTaskStatus interface {
	dara.Model
	String() string
	GoString() string
	SetLastSucceededTime(v string) *DatasetTaskStatus
	GetLastSucceededTime() *string
	SetStartTime(v string) *DatasetTaskStatus
	GetStartTime() *string
	SetStatus(v string) *DatasetTaskStatus
	GetStatus() *string
}

type DatasetTaskStatus struct {
	// example:
	//
	// 2024-06-29T14:50:13.011643661+08:00
	LastSucceededTime *string `json:"LastSucceededTime,omitempty" xml:"LastSucceededTime,omitempty"`
	// example:
	//
	// 2024-06-29T14:50:13.011643661+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DatasetTaskStatus) String() string {
	return dara.Prettify(s)
}

func (s DatasetTaskStatus) GoString() string {
	return s.String()
}

func (s *DatasetTaskStatus) GetLastSucceededTime() *string {
	return s.LastSucceededTime
}

func (s *DatasetTaskStatus) GetStartTime() *string {
	return s.StartTime
}

func (s *DatasetTaskStatus) GetStatus() *string {
	return s.Status
}

func (s *DatasetTaskStatus) SetLastSucceededTime(v string) *DatasetTaskStatus {
	s.LastSucceededTime = &v
	return s
}

func (s *DatasetTaskStatus) SetStartTime(v string) *DatasetTaskStatus {
	s.StartTime = &v
	return s
}

func (s *DatasetTaskStatus) SetStatus(v string) *DatasetTaskStatus {
	s.Status = &v
	return s
}

func (s *DatasetTaskStatus) Validate() error {
	return dara.Validate(s)
}
