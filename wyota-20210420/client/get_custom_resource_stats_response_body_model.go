// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomResourceStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomResourceStatsResponseBody
	GetCode() *string
	SetData(v *GetCustomResourceStatsResponseBodyData) *GetCustomResourceStatsResponseBody
	GetData() *GetCustomResourceStatsResponseBodyData
	SetHttpStatusCode(v int32) *GetCustomResourceStatsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCustomResourceStatsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomResourceStatsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomResourceStatsResponseBody
	GetSuccess() *bool
}

type GetCustomResourceStatsResponseBody struct {
	// The status code. 200 is returned if the call is successful. An error code is returned if the call fails.
	//
	// example:
	//
	// PARAM_ERROR
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The custom resource statistics information.
	Data *GetCustomResourceStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. This parameter is empty if the call is successful.
	//
	// example:
	//
	// parameter error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C5DCE54A-B266-522E-A6ED-468AF45F5AAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomResourceStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomResourceStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomResourceStatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomResourceStatsResponseBody) GetData() *GetCustomResourceStatsResponseBodyData {
	return s.Data
}

func (s *GetCustomResourceStatsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCustomResourceStatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomResourceStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomResourceStatsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomResourceStatsResponseBody) SetCode(v string) *GetCustomResourceStatsResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomResourceStatsResponseBody) SetData(v *GetCustomResourceStatsResponseBodyData) *GetCustomResourceStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomResourceStatsResponseBody) SetHttpStatusCode(v int32) *GetCustomResourceStatsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCustomResourceStatsResponseBody) SetMessage(v string) *GetCustomResourceStatsResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomResourceStatsResponseBody) SetRequestId(v string) *GetCustomResourceStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomResourceStatsResponseBody) SetSuccess(v bool) *GetCustomResourceStatsResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomResourceStatsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomResourceStatsResponseBodyData struct {
	// The number of terminals with custom resources configured.
	//
	// example:
	//
	// 10
	CustomResourceCount *int64 `json:"CustomResourceCount,omitempty" xml:"CustomResourceCount,omitempty"`
	// The number of terminals on which custom resources have taken effect.
	//
	// example:
	//
	// 8
	EffectiveCount *int64 `json:"EffectiveCount,omitempty" xml:"EffectiveCount,omitempty"`
	// The number of terminals without custom resources configured.
	//
	// example:
	//
	// 90
	NoCustomResourceCount *int64 `json:"NoCustomResourceCount,omitempty" xml:"NoCustomResourceCount,omitempty"`
	// The number of terminals on which custom resources have not taken effect.
	//
	// example:
	//
	// 2
	UnEffectiveCount *int64 `json:"UnEffectiveCount,omitempty" xml:"UnEffectiveCount,omitempty"`
}

func (s GetCustomResourceStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomResourceStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomResourceStatsResponseBodyData) GetCustomResourceCount() *int64 {
	return s.CustomResourceCount
}

func (s *GetCustomResourceStatsResponseBodyData) GetEffectiveCount() *int64 {
	return s.EffectiveCount
}

func (s *GetCustomResourceStatsResponseBodyData) GetNoCustomResourceCount() *int64 {
	return s.NoCustomResourceCount
}

func (s *GetCustomResourceStatsResponseBodyData) GetUnEffectiveCount() *int64 {
	return s.UnEffectiveCount
}

func (s *GetCustomResourceStatsResponseBodyData) SetCustomResourceCount(v int64) *GetCustomResourceStatsResponseBodyData {
	s.CustomResourceCount = &v
	return s
}

func (s *GetCustomResourceStatsResponseBodyData) SetEffectiveCount(v int64) *GetCustomResourceStatsResponseBodyData {
	s.EffectiveCount = &v
	return s
}

func (s *GetCustomResourceStatsResponseBodyData) SetNoCustomResourceCount(v int64) *GetCustomResourceStatsResponseBodyData {
	s.NoCustomResourceCount = &v
	return s
}

func (s *GetCustomResourceStatsResponseBodyData) SetUnEffectiveCount(v int64) *GetCustomResourceStatsResponseBodyData {
	s.UnEffectiveCount = &v
	return s
}

func (s *GetCustomResourceStatsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
