// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppDomainRedirectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DeleteAppDomainRedirectRequest
	GetBizId() *string
	SetRecordId(v int64) *DeleteAppDomainRedirectRequest
	GetRecordId() *int64
}

type DeleteAppDomainRedirectRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 803858889404426240
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DeleteAppDomainRedirectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppDomainRedirectRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppDomainRedirectRequest) GetBizId() *string {
	return s.BizId
}

func (s *DeleteAppDomainRedirectRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DeleteAppDomainRedirectRequest) SetBizId(v string) *DeleteAppDomainRedirectRequest {
	s.BizId = &v
	return s
}

func (s *DeleteAppDomainRedirectRequest) SetRecordId(v int64) *DeleteAppDomainRedirectRequest {
	s.RecordId = &v
	return s
}

func (s *DeleteAppDomainRedirectRequest) Validate() error {
	return dara.Validate(s)
}
