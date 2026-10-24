// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssetScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAssetScopeResponseBody
	GetRequestId() *string
}

type UpdateAssetScopeResponseBody struct {
	// The unique identifier of the request, used to troubleshoot and locate issues.
	//
	// example:
	//
	// 01A157FA-F957-3A5A-A6B8-5BBB8ADB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAssetScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetScopeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAssetScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAssetScopeResponseBody) SetRequestId(v string) *UpdateAssetScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAssetScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
