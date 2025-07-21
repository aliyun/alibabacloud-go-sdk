// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVersionDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersionName(v string) *GetVersionDownloadUrlRequest
	GetVersionName() *string
}

type GetVersionDownloadUrlRequest struct {
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetVersionDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVersionDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetVersionDownloadUrlRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *GetVersionDownloadUrlRequest) SetVersionName(v string) *GetVersionDownloadUrlRequest {
	s.VersionName = &v
	return s
}

func (s *GetVersionDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
