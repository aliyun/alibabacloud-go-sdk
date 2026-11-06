// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshProxyLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefreshProxyLevelResponseBody
	GetRequestId() *string
}

type RefreshProxyLevelResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshProxyLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshProxyLevelResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshProxyLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshProxyLevelResponseBody) SetRequestId(v string) *RefreshProxyLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshProxyLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
