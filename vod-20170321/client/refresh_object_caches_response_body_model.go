// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshObjectCachesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRefreshTaskId(v string) *RefreshObjectCachesResponseBody
	GetRefreshTaskId() *string
	SetRequestId(v string) *RefreshObjectCachesResponseBody
	GetRequestId() *string
}

type RefreshObjectCachesResponseBody struct {
	RefreshTaskId *string `json:"RefreshTaskId,omitempty" xml:"RefreshTaskId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshObjectCachesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshObjectCachesResponseBody) GetRefreshTaskId() *string {
	return s.RefreshTaskId
}

func (s *RefreshObjectCachesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshObjectCachesResponseBody) SetRefreshTaskId(v string) *RefreshObjectCachesResponseBody {
	s.RefreshTaskId = &v
	return s
}

func (s *RefreshObjectCachesResponseBody) SetRequestId(v string) *RefreshObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshObjectCachesResponseBody) Validate() error {
	return dara.Validate(s)
}
