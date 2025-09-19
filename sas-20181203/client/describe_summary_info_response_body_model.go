// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSummaryInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAegisClientOfflineCount(v int32) *DescribeSummaryInfoResponseBody
	GetAegisClientOfflineCount() *int32
	SetAegisClientOnlineCount(v int32) *DescribeSummaryInfoResponseBody
	GetAegisClientOnlineCount() *int32
	SetRequestId(v string) *DescribeSummaryInfoResponseBody
	GetRequestId() *string
	SetSecurityScore(v int32) *DescribeSummaryInfoResponseBody
	GetSecurityScore() *int32
	SetSuccess(v bool) *DescribeSummaryInfoResponseBody
	GetSuccess() *bool
}

type DescribeSummaryInfoResponseBody struct {
	// The number of unprotected assets.
	//
	// example:
	//
	// 12
	AegisClientOfflineCount *int32 `json:"AegisClientOfflineCount,omitempty" xml:"AegisClientOfflineCount,omitempty"`
	// The number of protected assets.
	//
	// example:
	//
	// 127
	AegisClientOnlineCount *int32 `json:"AegisClientOnlineCount,omitempty" xml:"AegisClientOnlineCount,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D40198E1-6EA8-482E-B3C7-D9573D75C0CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security score of the assets. Valid values:
	//
	// 	- 95 to 100: The assets are secure.
	//
	// 	- 85 to 94: The assets are exposed to a few security risks. We recommend that you reinforce your security system in a timely manner.
	//
	// 	- 70 to 84: The assets are exposed to multiple security risks. We recommend that you reinforce your security system in a timely manner.
	//
	// 	- 69 or lower: The current security system is unable to protect the assets against intrusions. We recommend that you reinforce your security system at the earliest opportunity.
	//
	// example:
	//
	// 44
	SecurityScore *int32 `json:"SecurityScore,omitempty" xml:"SecurityScore,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSummaryInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSummaryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSummaryInfoResponseBody) GetAegisClientOfflineCount() *int32 {
	return s.AegisClientOfflineCount
}

func (s *DescribeSummaryInfoResponseBody) GetAegisClientOnlineCount() *int32 {
	return s.AegisClientOnlineCount
}

func (s *DescribeSummaryInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSummaryInfoResponseBody) GetSecurityScore() *int32 {
	return s.SecurityScore
}

func (s *DescribeSummaryInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSummaryInfoResponseBody) SetAegisClientOfflineCount(v int32) *DescribeSummaryInfoResponseBody {
	s.AegisClientOfflineCount = &v
	return s
}

func (s *DescribeSummaryInfoResponseBody) SetAegisClientOnlineCount(v int32) *DescribeSummaryInfoResponseBody {
	s.AegisClientOnlineCount = &v
	return s
}

func (s *DescribeSummaryInfoResponseBody) SetRequestId(v string) *DescribeSummaryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSummaryInfoResponseBody) SetSecurityScore(v int32) *DescribeSummaryInfoResponseBody {
	s.SecurityScore = &v
	return s
}

func (s *DescribeSummaryInfoResponseBody) SetSuccess(v bool) *DescribeSummaryInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSummaryInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
