// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVerifyDownloadTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadTaskId(v string) *QueryVerifyDownloadTaskRequest
	GetDownloadTaskId() *string
}

type QueryVerifyDownloadTaskRequest struct {
	// Download task ID.
	//
	// example:
	//
	// 202411194002618
	DownloadTaskId *string `json:"DownloadTaskId,omitempty" xml:"DownloadTaskId,omitempty"`
}

func (s QueryVerifyDownloadTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyDownloadTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryVerifyDownloadTaskRequest) GetDownloadTaskId() *string {
	return s.DownloadTaskId
}

func (s *QueryVerifyDownloadTaskRequest) SetDownloadTaskId(v string) *QueryVerifyDownloadTaskRequest {
	s.DownloadTaskId = &v
	return s
}

func (s *QueryVerifyDownloadTaskRequest) Validate() error {
	return dara.Validate(s)
}
