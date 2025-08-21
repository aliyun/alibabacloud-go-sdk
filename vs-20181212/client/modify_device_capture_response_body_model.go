// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceCaptureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDeviceCaptureResponseBody
	GetRequestId() *string
}

type ModifyDeviceCaptureResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDeviceCaptureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceCaptureResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceCaptureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDeviceCaptureResponseBody) SetRequestId(v string) *ModifyDeviceCaptureResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDeviceCaptureResponseBody) Validate() error {
	return dara.Validate(s)
}
