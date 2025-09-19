// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefreshAssetsResponseBody
	GetRequestId() *string
}

type RefreshAssetsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 340D7FC4-D575-1661-8ACD-CFA7BE57B795
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshAssetsResponseBody) SetRequestId(v string) *RefreshAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshAssetsResponseBody) Validate() error {
	return dara.Validate(s)
}
