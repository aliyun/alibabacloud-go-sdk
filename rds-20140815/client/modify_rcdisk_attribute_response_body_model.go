// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDiskAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRCDiskAttributeResponseBody
	GetRequestId() *string
}

type ModifyRCDiskAttributeResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1E43AAE0-BEE8-43DA-860D-EAF2AA0724DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCDiskAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDiskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCDiskAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCDiskAttributeResponseBody) SetRequestId(v string) *ModifyRCDiskAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCDiskAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
