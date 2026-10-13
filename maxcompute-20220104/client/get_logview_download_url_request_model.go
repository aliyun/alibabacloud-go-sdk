// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogviewDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUseVpcEndpoint(v bool) *GetLogviewDownloadUrlRequest
	GetUseVpcEndpoint() *bool
}

type GetLogviewDownloadUrlRequest struct {
	// Specifies whether to use the OSS VPC endpoint to generate the download link. Valid values:
	//
	// - false: Uses the default OSS endpoint.
	//
	// - true: Uses the OSS VPC endpoint configured on the server. The download client must have network access to this endpoint.
	//
	// Default value: false. If this parameter is set to true and no OSS VPC endpoint is configured on the server, link generation fails and does not automatically fall back.
	UseVpcEndpoint *bool `json:"useVpcEndpoint,omitempty" xml:"useVpcEndpoint,omitempty"`
}

func (s GetLogviewDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogviewDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetLogviewDownloadUrlRequest) GetUseVpcEndpoint() *bool {
	return s.UseVpcEndpoint
}

func (s *GetLogviewDownloadUrlRequest) SetUseVpcEndpoint(v bool) *GetLogviewDownloadUrlRequest {
	s.UseVpcEndpoint = &v
	return s
}

func (s *GetLogviewDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
