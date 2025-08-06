// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackVodStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *RollbackVodStagingConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *RollbackVodStagingConfigRequest
	GetOwnerId() *int64
}

type RollbackVodStagingConfigRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s RollbackVodStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackVodStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *RollbackVodStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *RollbackVodStagingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RollbackVodStagingConfigRequest) SetDomainName(v string) *RollbackVodStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *RollbackVodStagingConfigRequest) SetOwnerId(v int64) *RollbackVodStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *RollbackVodStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
