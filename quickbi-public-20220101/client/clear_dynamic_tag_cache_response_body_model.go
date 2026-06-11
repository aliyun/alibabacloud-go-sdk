// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearDynamicTagCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ClearDynamicTagCacheResponseBody
	GetRequestId() *string
	SetResult(v bool) *ClearDynamicTagCacheResponseBody
	GetResult() *bool
	SetSuccess(v bool) *ClearDynamicTagCacheResponseBody
	GetSuccess() *bool
}

type ClearDynamicTagCacheResponseBody struct {
	// example:
	//
	// 78C17888****C462E57A95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ClearDynamicTagCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearDynamicTagCacheResponseBody) GoString() string {
	return s.String()
}

func (s *ClearDynamicTagCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearDynamicTagCacheResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ClearDynamicTagCacheResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ClearDynamicTagCacheResponseBody) SetRequestId(v string) *ClearDynamicTagCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearDynamicTagCacheResponseBody) SetResult(v bool) *ClearDynamicTagCacheResponseBody {
	s.Result = &v
	return s
}

func (s *ClearDynamicTagCacheResponseBody) SetSuccess(v bool) *ClearDynamicTagCacheResponseBody {
	s.Success = &v
	return s
}

func (s *ClearDynamicTagCacheResponseBody) Validate() error {
	return dara.Validate(s)
}
