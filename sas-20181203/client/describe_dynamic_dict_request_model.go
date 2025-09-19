// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDynamicDictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DescribeDynamicDictRequest
	GetSourceIp() *string
}

type DescribeDynamicDictRequest struct {
	// The source IP address of the request.
	//
	// example:
	//
	// 101.204.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDynamicDictRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicDictRequest) GoString() string {
	return s.String()
}

func (s *DescribeDynamicDictRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDynamicDictRequest) SetSourceIp(v string) *DescribeDynamicDictRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDynamicDictRequest) Validate() error {
	return dara.Validate(s)
}
