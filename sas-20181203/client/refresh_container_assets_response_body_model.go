// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshContainerAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefreshContainerAssetsResponseBody
	GetRequestId() *string
}

type RefreshContainerAssetsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2090F329-3658-49AF-820B-C4157FC31BCB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshContainerAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshContainerAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshContainerAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshContainerAssetsResponseBody) SetRequestId(v string) *RefreshContainerAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshContainerAssetsResponseBody) Validate() error {
	return dara.Validate(s)
}
