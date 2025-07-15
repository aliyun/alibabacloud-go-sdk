// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainMultiStreamConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeLiveDomainMultiStreamConfigRequest
	GetDomain() *string
	SetOwnerId(v int64) *DescribeLiveDomainMultiStreamConfigRequest
	GetOwnerId() *int64
}

type DescribeLiveDomainMultiStreamConfigRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// pliveplay.gstv.com.cn
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeLiveDomainMultiStreamConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMultiStreamConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMultiStreamConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveDomainMultiStreamConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainMultiStreamConfigRequest) SetDomain(v string) *DescribeLiveDomainMultiStreamConfigRequest {
	s.Domain = &v
	return s
}

func (s *DescribeLiveDomainMultiStreamConfigRequest) SetOwnerId(v int64) *DescribeLiveDomainMultiStreamConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainMultiStreamConfigRequest) Validate() error {
	return dara.Validate(s)
}
