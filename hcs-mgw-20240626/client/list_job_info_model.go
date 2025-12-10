// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobInfo interface {
	dara.Model
	String() string
	GoString() string
	SetImportJob(v []*CreateJobInfo) *ListJobInfo
	GetImportJob() []*CreateJobInfo
	SetNextMarker(v string) *ListJobInfo
	GetNextMarker() *string
	SetTruncated(v bool) *ListJobInfo
	GetTruncated() *bool
}

type ListJobInfo struct {
	ImportJob []*CreateJobInfo `json:"ImportJob,omitempty" xml:"ImportJob,omitempty" type:"Repeated"`
	// example:
	//
	// <your-next-job-name>
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// example:
	//
	// true
	Truncated *bool `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListJobInfo) String() string {
	return dara.Prettify(s)
}

func (s ListJobInfo) GoString() string {
	return s.String()
}

func (s *ListJobInfo) GetImportJob() []*CreateJobInfo {
	return s.ImportJob
}

func (s *ListJobInfo) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListJobInfo) GetTruncated() *bool {
	return s.Truncated
}

func (s *ListJobInfo) SetImportJob(v []*CreateJobInfo) *ListJobInfo {
	s.ImportJob = v
	return s
}

func (s *ListJobInfo) SetNextMarker(v string) *ListJobInfo {
	s.NextMarker = &v
	return s
}

func (s *ListJobInfo) SetTruncated(v bool) *ListJobInfo {
	s.Truncated = &v
	return s
}

func (s *ListJobInfo) Validate() error {
	if s.ImportJob != nil {
		for _, item := range s.ImportJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
