// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMediaByURLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadMediaByURLResponseBody
	GetRequestId() *string
	SetUploadJobs(v []*UploadMediaByURLResponseBodyUploadJobs) *UploadMediaByURLResponseBody
	GetUploadJobs() []*UploadMediaByURLResponseBodyUploadJobs
}

type UploadMediaByURLResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ****83B7-7F87-4792-BFE9-63CD2137****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about upload jobs.
	UploadJobs []*UploadMediaByURLResponseBodyUploadJobs `json:"UploadJobs,omitempty" xml:"UploadJobs,omitempty" type:"Repeated"`
}

func (s UploadMediaByURLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaByURLResponseBody) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadMediaByURLResponseBody) GetUploadJobs() []*UploadMediaByURLResponseBodyUploadJobs {
	return s.UploadJobs
}

func (s *UploadMediaByURLResponseBody) SetRequestId(v string) *UploadMediaByURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadMediaByURLResponseBody) SetUploadJobs(v []*UploadMediaByURLResponseBodyUploadJobs) *UploadMediaByURLResponseBody {
	s.UploadJobs = v
	return s
}

func (s *UploadMediaByURLResponseBody) Validate() error {
	if s.UploadJobs != nil {
		for _, item := range s.UploadJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UploadMediaByURLResponseBodyUploadJobs struct {
	// The ID of the upload job.
	//
	// example:
	//
	// 20ce1e05dba64576b96e9683879f0***
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// f476988629f54a7b8a4ba90d1a6c7***
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The URL of the source file that is uploaded in the upload job.
	//
	// example:
	//
	// http://example****.mp4
	SourceURL *string `json:"SourceURL,omitempty" xml:"SourceURL,omitempty"`
}

func (s UploadMediaByURLResponseBodyUploadJobs) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaByURLResponseBodyUploadJobs) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLResponseBodyUploadJobs) GetJobId() *string {
	return s.JobId
}

func (s *UploadMediaByURLResponseBodyUploadJobs) GetMediaId() *string {
	return s.MediaId
}

func (s *UploadMediaByURLResponseBodyUploadJobs) GetSourceURL() *string {
	return s.SourceURL
}

func (s *UploadMediaByURLResponseBodyUploadJobs) SetJobId(v string) *UploadMediaByURLResponseBodyUploadJobs {
	s.JobId = &v
	return s
}

func (s *UploadMediaByURLResponseBodyUploadJobs) SetMediaId(v string) *UploadMediaByURLResponseBodyUploadJobs {
	s.MediaId = &v
	return s
}

func (s *UploadMediaByURLResponseBodyUploadJobs) SetSourceURL(v string) *UploadMediaByURLResponseBodyUploadJobs {
	s.SourceURL = &v
	return s
}

func (s *UploadMediaByURLResponseBodyUploadJobs) Validate() error {
	return dara.Validate(s)
}
