// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDynamicDictUploadInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DescribeDynamicDictUploadInfoRequest
	GetSourceIp() *string
}

type DescribeDynamicDictUploadInfoRequest struct {
	// The source IP address of the request.
	//
	// example:
	//
	// 123.103.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDynamicDictUploadInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicDictUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDynamicDictUploadInfoRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDynamicDictUploadInfoRequest) SetSourceIp(v string) *DescribeDynamicDictUploadInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDynamicDictUploadInfoRequest) Validate() error {
	return dara.Validate(s)
}
