// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterAddonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClusterAddonResponseBody
	GetRequestId() *string
}

type ModifyClusterAddonResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterAddonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterAddonResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterAddonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterAddonResponseBody) SetRequestId(v string) *ModifyClusterAddonResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterAddonResponseBody) Validate() error {
	return dara.Validate(s)
}
