// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshOssBucketScanInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *RefreshOssBucketScanInfoRequest
	GetSource() *string
}

type RefreshOssBucketScanInfoRequest struct {
	// The service source. Valid values:
	//
	// - **OSS**: OSS
	//
	// - **NAS**: NAS
	//
	// example:
	//
	// OSS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s RefreshOssBucketScanInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshOssBucketScanInfoRequest) GoString() string {
	return s.String()
}

func (s *RefreshOssBucketScanInfoRequest) GetSource() *string {
	return s.Source
}

func (s *RefreshOssBucketScanInfoRequest) SetSource(v string) *RefreshOssBucketScanInfoRequest {
	s.Source = &v
	return s
}

func (s *RefreshOssBucketScanInfoRequest) Validate() error {
	return dara.Validate(s)
}
