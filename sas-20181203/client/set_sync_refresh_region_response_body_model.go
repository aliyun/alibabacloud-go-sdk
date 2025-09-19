// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSyncRefreshRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSyncRefreshRegionResponseBody
	GetRequestId() *string
}

type SetSyncRefreshRegionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9BB78BC9-07B9-578B-B020-C954E6FC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSyncRefreshRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSyncRefreshRegionResponseBody) GoString() string {
	return s.String()
}

func (s *SetSyncRefreshRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSyncRefreshRegionResponseBody) SetRequestId(v string) *SetSyncRefreshRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSyncRefreshRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
