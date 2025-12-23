// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIpcQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpcQuotaInfos(v []*QueryIpcQuotaResponseBodyIpcQuotaInfos) *QueryIpcQuotaResponseBody
	GetIpcQuotaInfos() []*QueryIpcQuotaResponseBodyIpcQuotaInfos
	SetRequestId(v string) *QueryIpcQuotaResponseBody
	GetRequestId() *string
	SetTotal(v string) *QueryIpcQuotaResponseBody
	GetTotal() *string
}

type QueryIpcQuotaResponseBody struct {
	IpcQuotaInfos []*QueryIpcQuotaResponseBodyIpcQuotaInfos `json:"IpcQuotaInfos,omitempty" xml:"IpcQuotaInfos,omitempty" type:"Repeated"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 39
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryIpcQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryIpcQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIpcQuotaResponseBody) GetIpcQuotaInfos() []*QueryIpcQuotaResponseBodyIpcQuotaInfos {
	return s.IpcQuotaInfos
}

func (s *QueryIpcQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryIpcQuotaResponseBody) GetTotal() *string {
	return s.Total
}

func (s *QueryIpcQuotaResponseBody) SetIpcQuotaInfos(v []*QueryIpcQuotaResponseBodyIpcQuotaInfos) *QueryIpcQuotaResponseBody {
	s.IpcQuotaInfos = v
	return s
}

func (s *QueryIpcQuotaResponseBody) SetRequestId(v string) *QueryIpcQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIpcQuotaResponseBody) SetTotal(v string) *QueryIpcQuotaResponseBody {
	s.Total = &v
	return s
}

func (s *QueryIpcQuotaResponseBody) Validate() error {
	if s.IpcQuotaInfos != nil {
		for _, item := range s.IpcQuotaInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryIpcQuotaResponseBodyIpcQuotaInfos struct {
	// example:
	//
	// understand
	Capability *string `json:"Capability,omitempty" xml:"Capability,omitempty"`
	// example:
	//
	// 32
	ConsumedQuota *int64 `json:"ConsumedQuota,omitempty" xml:"ConsumedQuota,omitempty"`
	// example:
	//
	// 2025-12-21T16:00:00Z
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// example:
	//
	// 10000
	MaxQuota *int64 `json:"MaxQuota,omitempty" xml:"MaxQuota,omitempty"`
}

func (s QueryIpcQuotaResponseBodyIpcQuotaInfos) String() string {
	return dara.Prettify(s)
}

func (s QueryIpcQuotaResponseBodyIpcQuotaInfos) GoString() string {
	return s.String()
}

func (s *QueryIpcQuotaResponseBodyIpcQuotaInfos) GetCapability() *string {
	return s.Capability
}

func (s *QueryIpcQuotaResponseBodyIpcQuotaInfos) GetConsumedQuota() *int64 {
	return s.ConsumedQuota
}

func (s *QueryIpcQuotaResponseBodyIpcQuotaInfos) GetDateTime() *string {
	return s.DateTime
}

func (s *QueryIpcQuotaResponseBodyIpcQuotaInfos) GetMaxQuota() *int64 {
	return s.MaxQuota
}

func (s *QueryIpcQuotaResponseBodyIpcQuotaInfos) SetCapability(v string) *QueryIpcQuotaResponseBodyIpcQuotaInfos {
	s.Capability = &v
	return s
}

func (s *QueryIpcQuotaResponseBodyIpcQuotaInfos) SetConsumedQuota(v int64) *QueryIpcQuotaResponseBodyIpcQuotaInfos {
	s.ConsumedQuota = &v
	return s
}

func (s *QueryIpcQuotaResponseBodyIpcQuotaInfos) SetDateTime(v string) *QueryIpcQuotaResponseBodyIpcQuotaInfos {
	s.DateTime = &v
	return s
}

func (s *QueryIpcQuotaResponseBodyIpcQuotaInfos) SetMaxQuota(v int64) *QueryIpcQuotaResponseBodyIpcQuotaInfos {
	s.MaxQuota = &v
	return s
}

func (s *QueryIpcQuotaResponseBodyIpcQuotaInfos) Validate() error {
	return dara.Validate(s)
}
