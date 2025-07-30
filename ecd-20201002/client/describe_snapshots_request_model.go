// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DescribeSnapshotsRequest
	GetClientId() *string
	SetDesktopId(v string) *DescribeSnapshotsRequest
	GetDesktopId() *string
	SetLoginToken(v string) *DescribeSnapshotsRequest
	GetLoginToken() *string
	SetMaxResults(v int32) *DescribeSnapshotsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSnapshotsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeSnapshotsRequest
	GetRegionId() *string
	SetSessionId(v string) *DescribeSnapshotsRequest
	GetSessionId() *string
	SetSnapshotId(v string) *DescribeSnapshotsRequest
	GetSnapshotId() *string
}

type DescribeSnapshotsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 16dad2b6-3c6d-4e4c-b057-78ecb13c****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v16abfb945208fc5745061668654680853da4a25202d1a394fcad57bba484e9827ad43ea7d10fb6bf13d44a4adc0e9****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// a99b9aca-bac5-4da2-819e-400ce98f****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// s-2ze81owrnv9pity4****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeSnapshotsRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeSnapshotsRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *DescribeSnapshotsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSnapshotsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnapshotsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnapshotsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeSnapshotsRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSnapshotsRequest) SetClientId(v string) *DescribeSnapshotsRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetDesktopId(v string) *DescribeSnapshotsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetLoginToken(v string) *DescribeSnapshotsRequest {
	s.LoginToken = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetMaxResults(v int32) *DescribeSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetNextToken(v string) *DescribeSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetRegionId(v string) *DescribeSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSessionId(v string) *DescribeSnapshotsRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotId(v string) *DescribeSnapshotsRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
