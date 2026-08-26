// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrustedOriginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTrustedOriginsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTrustedOriginsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTrustedOriginsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTrustedOriginsResponseBody
	GetTotalCount() *int32
	SetTrustedOrigins(v []*ListTrustedOriginsResponseBodyTrustedOrigins) *ListTrustedOriginsResponseBody
	GetTrustedOrigins() []*ListTrustedOriginsResponseBodyTrustedOrigins
}

type ListTrustedOriginsResponseBody struct {
	// The number of entries per page that takes effect for this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page query.
	//
	// example:
	//
	// NT_example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-example
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of trusted origins.
	TrustedOrigins []*ListTrustedOriginsResponseBodyTrustedOrigins `json:"TrustedOrigins,omitempty" xml:"TrustedOrigins,omitempty" type:"Repeated"`
}

func (s ListTrustedOriginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrustedOriginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrustedOriginsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTrustedOriginsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTrustedOriginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrustedOriginsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTrustedOriginsResponseBody) GetTrustedOrigins() []*ListTrustedOriginsResponseBodyTrustedOrigins {
	return s.TrustedOrigins
}

func (s *ListTrustedOriginsResponseBody) SetMaxResults(v int32) *ListTrustedOriginsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTrustedOriginsResponseBody) SetNextToken(v string) *ListTrustedOriginsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTrustedOriginsResponseBody) SetRequestId(v string) *ListTrustedOriginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrustedOriginsResponseBody) SetTotalCount(v int32) *ListTrustedOriginsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrustedOriginsResponseBody) SetTrustedOrigins(v []*ListTrustedOriginsResponseBodyTrustedOrigins) *ListTrustedOriginsResponseBody {
	s.TrustedOrigins = v
	return s
}

func (s *ListTrustedOriginsResponseBody) Validate() error {
	if s.TrustedOrigins != nil {
		for _, item := range s.TrustedOrigins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTrustedOriginsResponseBodyTrustedOrigins struct {
	// The creation time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-08-20T08:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_example
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The browser origin.
	//
	// example:
	//
	// https://console.qoder.com
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The status.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The trusted origin name.
	//
	// example:
	//
	// Qoder Production Console
	TrustOriginName *string `json:"TrustOriginName,omitempty" xml:"TrustOriginName,omitempty"`
	// The trusted origin ID.
	//
	// example:
	//
	// to_example
	TrustedOriginId *string `json:"TrustedOriginId,omitempty" xml:"TrustedOriginId,omitempty"`
	// The trusted origin scene.
	TrustedOriginScene []*string `json:"TrustedOriginScene,omitempty" xml:"TrustedOriginScene,omitempty" type:"Repeated"`
	// The update time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-08-20T08:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTrustedOriginsResponseBodyTrustedOrigins) String() string {
	return dara.Prettify(s)
}

func (s ListTrustedOriginsResponseBodyTrustedOrigins) GoString() string {
	return s.String()
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) GetOrigin() *string {
	return s.Origin
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) GetStatus() *string {
	return s.Status
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) GetTrustOriginName() *string {
	return s.TrustOriginName
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) GetTrustedOriginId() *string {
	return s.TrustedOriginId
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) GetTrustedOriginScene() []*string {
	return s.TrustedOriginScene
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) SetCreateTime(v string) *ListTrustedOriginsResponseBodyTrustedOrigins {
	s.CreateTime = &v
	return s
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) SetInstanceId(v string) *ListTrustedOriginsResponseBodyTrustedOrigins {
	s.InstanceId = &v
	return s
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) SetOrigin(v string) *ListTrustedOriginsResponseBodyTrustedOrigins {
	s.Origin = &v
	return s
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) SetStatus(v string) *ListTrustedOriginsResponseBodyTrustedOrigins {
	s.Status = &v
	return s
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) SetTrustOriginName(v string) *ListTrustedOriginsResponseBodyTrustedOrigins {
	s.TrustOriginName = &v
	return s
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) SetTrustedOriginId(v string) *ListTrustedOriginsResponseBodyTrustedOrigins {
	s.TrustedOriginId = &v
	return s
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) SetTrustedOriginScene(v []*string) *ListTrustedOriginsResponseBodyTrustedOrigins {
	s.TrustedOriginScene = v
	return s
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) SetUpdateTime(v string) *ListTrustedOriginsResponseBodyTrustedOrigins {
	s.UpdateTime = &v
	return s
}

func (s *ListTrustedOriginsResponseBodyTrustedOrigins) Validate() error {
	return dara.Validate(s)
}
