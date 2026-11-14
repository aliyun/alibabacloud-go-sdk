// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResourceQuotaResponseBody
	GetCode() *string
	SetData(v int64) *GetResourceQuotaResponseBody
	GetData() *int64
	SetMessage(v string) *GetResourceQuotaResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetResourceQuotaResponseBody
	GetRequestId() *string
}

type GetResourceQuotaResponseBody struct {
	// The error code returned when the call fails. A normal result is not guaranteed to contain this field. You should not use the return of a specific success code as a prerequisite for obtaining the quota.
	//
	// example:
	//
	// InvalidProductType.ValueNotSupported
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The maximum resource count quota applicable under the current query conditions. It is not the usage or remaining quota.
	//
	// - `ResourceType=AppInstanceGroup`: Unit: groups.
	//
	// - `ResourceType=AppAuthUser`: Unit: users.
	//
	// - `ResourceType=PoolNodeAmount`: Unit: nodes.
	//
	// If no smaller limit is obtained, `2147483647` is returned. This does not mean that the resource inventory is unlimited.
	//
	// example:
	//
	// 30
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned when the call fails. It is used together with the error code to help locate the cause.
	//
	// example:
	//
	// The ProductType is not supported.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceQuotaResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResourceQuotaResponseBody) GetData() *int64 {
	return s.Data
}

func (s *GetResourceQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResourceQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceQuotaResponseBody) SetCode(v string) *GetResourceQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourceQuotaResponseBody) SetData(v int64) *GetResourceQuotaResponseBody {
	s.Data = &v
	return s
}

func (s *GetResourceQuotaResponseBody) SetMessage(v string) *GetResourceQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourceQuotaResponseBody) SetRequestId(v string) *GetResourceQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
