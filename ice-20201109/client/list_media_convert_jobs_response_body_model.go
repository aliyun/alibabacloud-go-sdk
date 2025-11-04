// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaConvertJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*MediaConvertJobWithoutDetail) *ListMediaConvertJobsResponseBody
	GetJobs() []*MediaConvertJobWithoutDetail
	SetNextPageToken(v string) *ListMediaConvertJobsResponseBody
	GetNextPageToken() *string
	SetRequestId(v string) *ListMediaConvertJobsResponseBody
	GetRequestId() *string
}

type ListMediaConvertJobsResponseBody struct {
	// The tasks.
	Jobs []*MediaConvertJobWithoutDetail `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// Indicates the read position returned by the current call. An empty value means all data has been read.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMediaConvertJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaConvertJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaConvertJobsResponseBody) GetJobs() []*MediaConvertJobWithoutDetail {
	return s.Jobs
}

func (s *ListMediaConvertJobsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListMediaConvertJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaConvertJobsResponseBody) SetJobs(v []*MediaConvertJobWithoutDetail) *ListMediaConvertJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListMediaConvertJobsResponseBody) SetNextPageToken(v string) *ListMediaConvertJobsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListMediaConvertJobsResponseBody) SetRequestId(v string) *ListMediaConvertJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaConvertJobsResponseBody) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
