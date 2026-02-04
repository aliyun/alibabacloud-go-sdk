// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuota(v *GetInstanceQuotaResponseBodyQuota) *GetInstanceQuotaResponseBody
	GetQuota() *GetInstanceQuotaResponseBodyQuota
	SetRequestId(v string) *GetInstanceQuotaResponseBody
	GetRequestId() *string
}

type GetInstanceQuotaResponseBody struct {
	Quota *GetInstanceQuotaResponseBodyQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceQuotaResponseBody) GetQuota() *GetInstanceQuotaResponseBodyQuota {
	return s.Quota
}

func (s *GetInstanceQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceQuotaResponseBody) SetQuota(v *GetInstanceQuotaResponseBodyQuota) *GetInstanceQuotaResponseBody {
	s.Quota = v
	return s
}

func (s *GetInstanceQuotaResponseBody) SetRequestId(v string) *GetInstanceQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceQuotaResponseBody) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceQuotaResponseBodyQuota struct {
	// Quota 配额的Key，同请求参数
	//
	// example:
	//
	// userMaxNumber
	QuotaKey *string `json:"QuotaKey,omitempty" xml:"QuotaKey,omitempty"`
	// Quota 配额的值。
	//
	// example:
	//
	// 5
	QuotaValue *int32 `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
}

func (s GetInstanceQuotaResponseBodyQuota) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceQuotaResponseBodyQuota) GoString() string {
	return s.String()
}

func (s *GetInstanceQuotaResponseBodyQuota) GetQuotaKey() *string {
	return s.QuotaKey
}

func (s *GetInstanceQuotaResponseBodyQuota) GetQuotaValue() *int32 {
	return s.QuotaValue
}

func (s *GetInstanceQuotaResponseBodyQuota) SetQuotaKey(v string) *GetInstanceQuotaResponseBodyQuota {
	s.QuotaKey = &v
	return s
}

func (s *GetInstanceQuotaResponseBodyQuota) SetQuotaValue(v int32) *GetInstanceQuotaResponseBodyQuota {
	s.QuotaValue = &v
	return s
}

func (s *GetInstanceQuotaResponseBodyQuota) Validate() error {
	return dara.Validate(s)
}
