// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPurgeQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuota(v string) *GetPurgeQuotaResponseBody
	GetQuota() *string
	SetRequestId(v string) *GetPurgeQuotaResponseBody
	GetRequestId() *string
	SetUsage(v string) *GetPurgeQuotaResponseBody
	GetUsage() *string
}

type GetPurgeQuotaResponseBody struct {
	// The total quota.
	//
	// example:
	//
	// 100000
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The quota usage.
	//
	// example:
	//
	// 10
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s GetPurgeQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPurgeQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetPurgeQuotaResponseBody) GetQuota() *string {
	return s.Quota
}

func (s *GetPurgeQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPurgeQuotaResponseBody) GetUsage() *string {
	return s.Usage
}

func (s *GetPurgeQuotaResponseBody) SetQuota(v string) *GetPurgeQuotaResponseBody {
	s.Quota = &v
	return s
}

func (s *GetPurgeQuotaResponseBody) SetRequestId(v string) *GetPurgeQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPurgeQuotaResponseBody) SetUsage(v string) *GetPurgeQuotaResponseBody {
	s.Usage = &v
	return s
}

func (s *GetPurgeQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
