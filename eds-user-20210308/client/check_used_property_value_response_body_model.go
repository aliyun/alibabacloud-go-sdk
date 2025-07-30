// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUsedPropertyValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckUsedPropertyValueResponseBody
	GetRequestId() *string
	SetUseCount(v int64) *CheckUsedPropertyValueResponseBody
	GetUseCount() *int64
}

type CheckUsedPropertyValueResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of convenience accounts that are associated with the specified custom property value.
	//
	// example:
	//
	// 1
	UseCount *int64 `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
}

func (s CheckUsedPropertyValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckUsedPropertyValueResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckUsedPropertyValueResponseBody) GetUseCount() *int64 {
	return s.UseCount
}

func (s *CheckUsedPropertyValueResponseBody) SetRequestId(v string) *CheckUsedPropertyValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUsedPropertyValueResponseBody) SetUseCount(v int64) *CheckUsedPropertyValueResponseBody {
	s.UseCount = &v
	return s
}

func (s *CheckUsedPropertyValueResponseBody) Validate() error {
	return dara.Validate(s)
}
