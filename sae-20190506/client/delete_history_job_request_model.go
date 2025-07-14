// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHistoryJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteHistoryJobRequest
	GetAppId() *string
	SetJobId(v string) *DeleteHistoryJobRequest
	GetJobId() *string
}

type DeleteHistoryJobRequest struct {
	// The ID of the job template to which the job that you want to delete belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// manual-3db7a8fa-5d40-4edc-92e4-49d50eab****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteHistoryJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHistoryJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteHistoryJobRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteHistoryJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DeleteHistoryJobRequest) SetAppId(v string) *DeleteHistoryJobRequest {
	s.AppId = &v
	return s
}

func (s *DeleteHistoryJobRequest) SetJobId(v string) *DeleteHistoryJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteHistoryJobRequest) Validate() error {
	return dara.Validate(s)
}
