// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDcdnObjectCachesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRefreshTaskId(v string) *RefreshDcdnObjectCachesResponseBody
	GetRefreshTaskId() *string
	SetRequestId(v string) *RefreshDcdnObjectCachesResponseBody
	GetRequestId() *string
}

type RefreshDcdnObjectCachesResponseBody struct {
	// The ID of the refresh task. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// 95248880
	RefreshTaskId *string `json:"RefreshTaskId,omitempty" xml:"RefreshTaskId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E5BD4B50-7A02-493A-AE0B-97B9024B4135
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshDcdnObjectCachesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshDcdnObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDcdnObjectCachesResponseBody) GetRefreshTaskId() *string {
	return s.RefreshTaskId
}

func (s *RefreshDcdnObjectCachesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshDcdnObjectCachesResponseBody) SetRefreshTaskId(v string) *RefreshDcdnObjectCachesResponseBody {
	s.RefreshTaskId = &v
	return s
}

func (s *RefreshDcdnObjectCachesResponseBody) SetRequestId(v string) *RefreshDcdnObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshDcdnObjectCachesResponseBody) Validate() error {
	return dara.Validate(s)
}
