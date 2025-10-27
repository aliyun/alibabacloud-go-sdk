// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthVersionStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAuthVersionStatisticResponseBody
	GetRequestId() *string
	SetStatistics(v []*GetAuthVersionStatisticResponseBodyStatistics) *GetAuthVersionStatisticResponseBody
	GetStatistics() []*GetAuthVersionStatisticResponseBodyStatistics
}

type GetAuthVersionStatisticResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2CA2BDF6-F3BD-51A4-BAAC-30B02F7A3FBB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics about the numbers of assets protected by each edition of Security Center.
	Statistics []*GetAuthVersionStatisticResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s GetAuthVersionStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthVersionStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthVersionStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthVersionStatisticResponseBody) GetStatistics() []*GetAuthVersionStatisticResponseBodyStatistics {
	return s.Statistics
}

func (s *GetAuthVersionStatisticResponseBody) SetRequestId(v string) *GetAuthVersionStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthVersionStatisticResponseBody) SetStatistics(v []*GetAuthVersionStatisticResponseBodyStatistics) *GetAuthVersionStatisticResponseBody {
	s.Statistics = v
	return s
}

func (s *GetAuthVersionStatisticResponseBody) Validate() error {
	if s.Statistics != nil {
		for _, item := range s.Statistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAuthVersionStatisticResponseBodyStatistics struct {
	// The edition of Security Center. Valid values:
	//
	// 	- **1**: Basic edition (Unauthorized)
	//
	// 	- **6**: Anti-virus edition
	//
	// 	- **5**: Advanced edition
	//
	// 	- **3**: Enterprise edition
	//
	// 	- **7**: Ultimate edition
	//
	// 	- **10**: Value-added Plan edition
	//
	// example:
	//
	// 6
	AuthVersion *int32 `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// The number of authorized servers.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetAuthVersionStatisticResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetAuthVersionStatisticResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *GetAuthVersionStatisticResponseBodyStatistics) GetAuthVersion() *int32 {
	return s.AuthVersion
}

func (s *GetAuthVersionStatisticResponseBodyStatistics) GetCount() *int32 {
	return s.Count
}

func (s *GetAuthVersionStatisticResponseBodyStatistics) SetAuthVersion(v int32) *GetAuthVersionStatisticResponseBodyStatistics {
	s.AuthVersion = &v
	return s
}

func (s *GetAuthVersionStatisticResponseBodyStatistics) SetCount(v int32) *GetAuthVersionStatisticResponseBodyStatistics {
	s.Count = &v
	return s
}

func (s *GetAuthVersionStatisticResponseBodyStatistics) Validate() error {
	return dara.Validate(s)
}
