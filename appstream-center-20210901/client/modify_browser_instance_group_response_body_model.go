// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBrowserInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBrowserInstanceGroupResponseBody
	GetRequestId() *string
}

type ModifyBrowserInstanceGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBrowserInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBrowserInstanceGroupResponseBody) SetRequestId(v string) *ModifyBrowserInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBrowserInstanceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
