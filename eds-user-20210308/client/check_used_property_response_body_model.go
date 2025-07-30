// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUsedPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckUsedPropertyResponseBody
	GetRequestId() *string
	SetUseCount(v int64) *CheckUsedPropertyResponseBody
	GetUseCount() *int64
}

type CheckUsedPropertyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 22C97624-2405-54AC-BD44-A63FBE97CC2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of convenience users that are associated with the property.
	//
	// example:
	//
	// 7
	UseCount *int64 `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
}

func (s CheckUsedPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckUsedPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckUsedPropertyResponseBody) GetUseCount() *int64 {
	return s.UseCount
}

func (s *CheckUsedPropertyResponseBody) SetRequestId(v string) *CheckUsedPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUsedPropertyResponseBody) SetUseCount(v int64) *CheckUsedPropertyResponseBody {
	s.UseCount = &v
	return s
}

func (s *CheckUsedPropertyResponseBody) Validate() error {
	return dara.Validate(s)
}
