// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnInstallAddonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnInstallAddonResponseBody
	GetRequestId() *string
}

type UnInstallAddonResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnInstallAddonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnInstallAddonResponseBody) GoString() string {
	return s.String()
}

func (s *UnInstallAddonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnInstallAddonResponseBody) SetRequestId(v string) *UnInstallAddonResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnInstallAddonResponseBody) Validate() error {
	return dara.Validate(s)
}
