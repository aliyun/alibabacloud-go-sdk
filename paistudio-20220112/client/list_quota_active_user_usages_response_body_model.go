// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaActiveUserUsagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaUserUsage(v []*QuotaUser) *ListQuotaActiveUserUsagesResponseBody
	GetQuotaUserUsage() []*QuotaUser
	SetQuotaUserUsages(v []*QuotaUser) *ListQuotaActiveUserUsagesResponseBody
	GetQuotaUserUsages() []*QuotaUser
	SetRequestId(v string) *ListQuotaActiveUserUsagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListQuotaActiveUserUsagesResponseBody
	GetTotalCount() *int32
}

type ListQuotaActiveUserUsagesResponseBody struct {
	QuotaUserUsage  []*QuotaUser `json:"QuotaUserUsage,omitempty" xml:"QuotaUserUsage,omitempty" type:"Repeated"`
	QuotaUserUsages []*QuotaUser `json:"QuotaUserUsages,omitempty" xml:"QuotaUserUsages,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// E7C42CC7-2E85-508A-84F4-923B605FD10F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotaActiveUserUsagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaActiveUserUsagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaActiveUserUsagesResponseBody) GetQuotaUserUsage() []*QuotaUser {
	return s.QuotaUserUsage
}

func (s *ListQuotaActiveUserUsagesResponseBody) GetQuotaUserUsages() []*QuotaUser {
	return s.QuotaUserUsages
}

func (s *ListQuotaActiveUserUsagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotaActiveUserUsagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListQuotaActiveUserUsagesResponseBody) SetQuotaUserUsage(v []*QuotaUser) *ListQuotaActiveUserUsagesResponseBody {
	s.QuotaUserUsage = v
	return s
}

func (s *ListQuotaActiveUserUsagesResponseBody) SetQuotaUserUsages(v []*QuotaUser) *ListQuotaActiveUserUsagesResponseBody {
	s.QuotaUserUsages = v
	return s
}

func (s *ListQuotaActiveUserUsagesResponseBody) SetRequestId(v string) *ListQuotaActiveUserUsagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaActiveUserUsagesResponseBody) SetTotalCount(v int32) *ListQuotaActiveUserUsagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotaActiveUserUsagesResponseBody) Validate() error {
	if s.QuotaUserUsage != nil {
		for _, item := range s.QuotaUserUsage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QuotaUserUsages != nil {
		for _, item := range s.QuotaUserUsages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
