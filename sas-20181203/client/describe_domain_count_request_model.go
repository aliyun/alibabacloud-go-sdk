// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DescribeDomainCountRequest
	GetSourceIp() *string
}

type DescribeDomainCountRequest struct {
	// The source IP address of the request.
	//
	// example:
	//
	// 192.172.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainCountRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainCountRequest) SetSourceIp(v string) *DescribeDomainCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainCountRequest) Validate() error {
	return dara.Validate(s)
}
