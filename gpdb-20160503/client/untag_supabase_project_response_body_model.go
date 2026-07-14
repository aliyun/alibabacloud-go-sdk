// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagSupabaseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UntagSupabaseProjectResponseBody
	GetRequestId() *string
}

type UntagSupabaseProjectResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 34b32a0a-08ef-4a87-b6be-cdd9f56fc3ad
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagSupabaseProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagSupabaseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UntagSupabaseProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagSupabaseProjectResponseBody) SetRequestId(v string) *UntagSupabaseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagSupabaseProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
