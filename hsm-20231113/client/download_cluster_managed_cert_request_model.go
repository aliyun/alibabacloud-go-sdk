// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadClusterManagedCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DownloadClusterManagedCertRequest
	GetClusterId() *string
}

type DownloadClusterManagedCertRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-001***hui
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DownloadClusterManagedCertRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadClusterManagedCertRequest) GoString() string {
	return s.String()
}

func (s *DownloadClusterManagedCertRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DownloadClusterManagedCertRequest) SetClusterId(v string) *DownloadClusterManagedCertRequest {
	s.ClusterId = &v
	return s
}

func (s *DownloadClusterManagedCertRequest) Validate() error {
	return dara.Validate(s)
}
