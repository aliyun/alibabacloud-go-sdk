// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaName(v string) *ListInstancesQuotaResponseBody
	GetQuotaName() *string
	SetQuotaValueType(v string) *ListInstancesQuotaResponseBody
	GetQuotaValueType() *string
	SetQuotaValues(v []*ListInstancesQuotaResponseBodyQuotaValues) *ListInstancesQuotaResponseBody
	GetQuotaValues() []*ListInstancesQuotaResponseBodyQuotaValues
	SetRequestId(v string) *ListInstancesQuotaResponseBody
	GetRequestId() *string
}

type ListInstancesQuotaResponseBody struct {
	// The quota name.
	//
	// example:
	//
	// siteCount
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The threshold type of the quota. Valid values:
	//
	// - **value**: Enumeration type. The enumeration range of quota values.
	//
	// - **bool**: Boolean type. Indicates whether the quota is available.
	//
	// - **num**: Numeric type. The upper limit of the quota usage.
	//
	// - **range**: Range type. The value range of the quota.
	//
	// - **custom**: Custom type. Other types beyond the four threshold types above.
	//
	// example:
	//
	// bool
	QuotaValueType *string `json:"QuotaValueType,omitempty" xml:"QuotaValueType,omitempty"`
	// The list of quota values.
	QuotaValues []*ListInstancesQuotaResponseBodyQuotaValues `json:"QuotaValues,omitempty" xml:"QuotaValues,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 156A6B-677B1A-4297B7-9187B7-2B44792
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancesQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesQuotaResponseBody) GetQuotaName() *string {
	return s.QuotaName
}

func (s *ListInstancesQuotaResponseBody) GetQuotaValueType() *string {
	return s.QuotaValueType
}

func (s *ListInstancesQuotaResponseBody) GetQuotaValues() []*ListInstancesQuotaResponseBodyQuotaValues {
	return s.QuotaValues
}

func (s *ListInstancesQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesQuotaResponseBody) SetQuotaName(v string) *ListInstancesQuotaResponseBody {
	s.QuotaName = &v
	return s
}

func (s *ListInstancesQuotaResponseBody) SetQuotaValueType(v string) *ListInstancesQuotaResponseBody {
	s.QuotaValueType = &v
	return s
}

func (s *ListInstancesQuotaResponseBody) SetQuotaValues(v []*ListInstancesQuotaResponseBodyQuotaValues) *ListInstancesQuotaResponseBody {
	s.QuotaValues = v
	return s
}

func (s *ListInstancesQuotaResponseBody) SetRequestId(v string) *ListInstancesQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesQuotaResponseBody) Validate() error {
	if s.QuotaValues != nil {
		for _, item := range s.QuotaValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesQuotaResponseBodyQuotaValues struct {
	// The instance ID.
	//
	// example:
	//
	// esa-site-blea5hv7m0ow
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The quota value.
	//
	// example:
	//
	// 10
	QuotaValue *string `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
}

func (s ListInstancesQuotaResponseBodyQuotaValues) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesQuotaResponseBodyQuotaValues) GoString() string {
	return s.String()
}

func (s *ListInstancesQuotaResponseBodyQuotaValues) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesQuotaResponseBodyQuotaValues) GetQuotaValue() *string {
	return s.QuotaValue
}

func (s *ListInstancesQuotaResponseBodyQuotaValues) SetInstanceId(v string) *ListInstancesQuotaResponseBodyQuotaValues {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesQuotaResponseBodyQuotaValues) SetQuotaValue(v string) *ListInstancesQuotaResponseBodyQuotaValues {
	s.QuotaValue = &v
	return s
}

func (s *ListInstancesQuotaResponseBodyQuotaValues) Validate() error {
	return dara.Validate(s)
}
