// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobResp interface {
	dara.Model
	String() string
	GoString() string
	SetImportJob(v []*GetJobResp) *ListJobResp
	GetImportJob() []*GetJobResp
	SetNextMarker(v string) *ListJobResp
	GetNextMarker() *string
	SetTruncated(v bool) *ListJobResp
	GetTruncated() *bool
}

type ListJobResp struct {
	ImportJob []*GetJobResp `json:"ImportJob,omitempty" xml:"ImportJob,omitempty" type:"Repeated"`
	// example:
	//
	// <your-next-job-name>
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// example:
	//
	// true
	Truncated *bool `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListJobResp) String() string {
	return dara.Prettify(s)
}

func (s ListJobResp) GoString() string {
	return s.String()
}

func (s *ListJobResp) GetImportJob() []*GetJobResp {
	return s.ImportJob
}

func (s *ListJobResp) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListJobResp) GetTruncated() *bool {
	return s.Truncated
}

func (s *ListJobResp) SetImportJob(v []*GetJobResp) *ListJobResp {
	s.ImportJob = v
	return s
}

func (s *ListJobResp) SetNextMarker(v string) *ListJobResp {
	s.NextMarker = &v
	return s
}

func (s *ListJobResp) SetTruncated(v bool) *ListJobResp {
	s.Truncated = &v
	return s
}

func (s *ListJobResp) Validate() error {
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
