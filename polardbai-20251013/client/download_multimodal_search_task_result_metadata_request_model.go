// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadMultimodalSearchTaskResultMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DownloadMultimodalSearchTaskResultMetadataRequest
	GetDBClusterId() *string
	SetTaskId(v string) *DownloadMultimodalSearchTaskResultMetadataRequest
	GetTaskId() *string
}

type DownloadMultimodalSearchTaskResultMetadataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 白色汽车
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DownloadMultimodalSearchTaskResultMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadMultimodalSearchTaskResultMetadataRequest) GoString() string {
	return s.String()
}

func (s *DownloadMultimodalSearchTaskResultMetadataRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DownloadMultimodalSearchTaskResultMetadataRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DownloadMultimodalSearchTaskResultMetadataRequest) SetDBClusterId(v string) *DownloadMultimodalSearchTaskResultMetadataRequest {
	s.DBClusterId = &v
	return s
}

func (s *DownloadMultimodalSearchTaskResultMetadataRequest) SetTaskId(v string) *DownloadMultimodalSearchTaskResultMetadataRequest {
	s.TaskId = &v
	return s
}

func (s *DownloadMultimodalSearchTaskResultMetadataRequest) Validate() error {
	return dara.Validate(s)
}
