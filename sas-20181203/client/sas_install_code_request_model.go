// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSasInstallCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *SasInstallCodeRequest
	GetSourceIp() *string
}

type SasInstallCodeRequest struct {
	// The source IP address of the request.
	//
	// example:
	//
	// 120.41.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s SasInstallCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s SasInstallCodeRequest) GoString() string {
	return s.String()
}

func (s *SasInstallCodeRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *SasInstallCodeRequest) SetSourceIp(v string) *SasInstallCodeRequest {
	s.SourceIp = &v
	return s
}

func (s *SasInstallCodeRequest) Validate() error {
	return dara.Validate(s)
}
