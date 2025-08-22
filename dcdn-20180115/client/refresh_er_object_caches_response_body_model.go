// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshErObjectCachesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRefreshTaskId(v string) *RefreshErObjectCachesResponseBody
	GetRefreshTaskId() *string
	SetRequestId(v string) *RefreshErObjectCachesResponseBody
	GetRequestId() *string
}

type RefreshErObjectCachesResponseBody struct {
	// The ID of the refresh task. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 95248880
	RefreshTaskId *string `json:"RefreshTaskId,omitempty" xml:"RefreshTaskId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshErObjectCachesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshErObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshErObjectCachesResponseBody) GetRefreshTaskId() *string {
	return s.RefreshTaskId
}

func (s *RefreshErObjectCachesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshErObjectCachesResponseBody) SetRefreshTaskId(v string) *RefreshErObjectCachesResponseBody {
	s.RefreshTaskId = &v
	return s
}

func (s *RefreshErObjectCachesResponseBody) SetRequestId(v string) *RefreshErObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshErObjectCachesResponseBody) Validate() error {
	return dara.Validate(s)
}
