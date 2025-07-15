// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAndroidInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAndroidInstanceGroupResponseBody
	GetRequestId() *string
}

type ModifyAndroidInstanceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6C83EBE3-F267-5F11-ABF8-4E7B90B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAndroidInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAndroidInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAndroidInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAndroidInstanceGroupResponseBody) SetRequestId(v string) *ModifyAndroidInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAndroidInstanceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
