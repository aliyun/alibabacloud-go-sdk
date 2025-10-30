// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnrolledAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountUid(v int64) *GetEnrolledAccountRequest
	GetAccountUid() *int64
	SetRegionId(v string) *GetEnrolledAccountRequest
	GetRegionId() *string
}

type GetEnrolledAccountRequest struct {
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 19534534552****
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetEnrolledAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEnrolledAccountRequest) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountRequest) GetAccountUid() *int64 {
	return s.AccountUid
}

func (s *GetEnrolledAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEnrolledAccountRequest) SetAccountUid(v int64) *GetEnrolledAccountRequest {
	s.AccountUid = &v
	return s
}

func (s *GetEnrolledAccountRequest) SetRegionId(v string) *GetEnrolledAccountRequest {
	s.RegionId = &v
	return s
}

func (s *GetEnrolledAccountRequest) Validate() error {
	return dara.Validate(s)
}
