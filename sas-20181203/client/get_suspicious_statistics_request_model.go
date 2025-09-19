// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuspiciousStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupIdList(v string) *GetSuspiciousStatisticsRequest
	GetGroupIdList() *string
	SetSourceIp(v string) *GetSuspiciousStatisticsRequest
	GetSourceIp() *string
}

type GetSuspiciousStatisticsRequest struct {
	// The ID of the asset group. Separate multiple IDs with commas (,).
	//
	// > You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the IDs of asset groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9997897
	GroupIdList *string `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 10.12.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s GetSuspiciousStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSuspiciousStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetSuspiciousStatisticsRequest) GetGroupIdList() *string {
	return s.GroupIdList
}

func (s *GetSuspiciousStatisticsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *GetSuspiciousStatisticsRequest) SetGroupIdList(v string) *GetSuspiciousStatisticsRequest {
	s.GroupIdList = &v
	return s
}

func (s *GetSuspiciousStatisticsRequest) SetSourceIp(v string) *GetSuspiciousStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *GetSuspiciousStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
