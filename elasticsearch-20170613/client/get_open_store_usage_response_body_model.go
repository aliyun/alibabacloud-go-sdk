// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenStoreUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetOpenStoreUsageResponseBody
	GetRequestId() *string
	SetResult(v *GetOpenStoreUsageResponseBodyResult) *GetOpenStoreUsageResponseBody
	GetResult() *GetOpenStoreUsageResponseBodyResult
}

type GetOpenStoreUsageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E1DE2491-804F-4C86-BAB4-548DD70B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The current request result.
	Result *GetOpenStoreUsageResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetOpenStoreUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpenStoreUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenStoreUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpenStoreUsageResponseBody) GetResult() *GetOpenStoreUsageResponseBodyResult {
	return s.Result
}

func (s *GetOpenStoreUsageResponseBody) SetRequestId(v string) *GetOpenStoreUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpenStoreUsageResponseBody) SetResult(v *GetOpenStoreUsageResponseBodyResult) *GetOpenStoreUsageResponseBody {
	s.Result = v
	return s
}

func (s *GetOpenStoreUsageResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOpenStoreUsageResponseBodyResult struct {
	// The current OpenStore storage capacity (estimated value based on actual indexes). Unit: Byte.
	//
	// example:
	//
	// 204800
	CurrentUsage *int64 `json:"currentUsage,omitempty" xml:"currentUsage,omitempty"`
	// The storage capacity of OpenStore yesterday. Unit: bytes.
	//
	// example:
	//
	// 184320
	LastDayUsage *int64 `json:"lastDayUsage,omitempty" xml:"lastDayUsage,omitempty"`
}

func (s GetOpenStoreUsageResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetOpenStoreUsageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOpenStoreUsageResponseBodyResult) GetCurrentUsage() *int64 {
	return s.CurrentUsage
}

func (s *GetOpenStoreUsageResponseBodyResult) GetLastDayUsage() *int64 {
	return s.LastDayUsage
}

func (s *GetOpenStoreUsageResponseBodyResult) SetCurrentUsage(v int64) *GetOpenStoreUsageResponseBodyResult {
	s.CurrentUsage = &v
	return s
}

func (s *GetOpenStoreUsageResponseBodyResult) SetLastDayUsage(v int64) *GetOpenStoreUsageResponseBodyResult {
	s.LastDayUsage = &v
	return s
}

func (s *GetOpenStoreUsageResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
