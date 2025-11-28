// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadPrivacyKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DownloadPrivacyKeyRequest
	GetRegionId() *string
}

type DownloadPrivacyKeyRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DownloadPrivacyKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadPrivacyKeyRequest) GoString() string {
	return s.String()
}

func (s *DownloadPrivacyKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DownloadPrivacyKeyRequest) SetRegionId(v string) *DownloadPrivacyKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DownloadPrivacyKeyRequest) Validate() error {
	return dara.Validate(s)
}
