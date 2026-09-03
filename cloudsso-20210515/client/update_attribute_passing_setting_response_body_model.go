// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAttributePassingSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAttributePassingSettingResponseBody
	GetRequestId() *string
}

type UpdateAttributePassingSettingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 63160579-2E1B-57B0-8273-B27427172385
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAttributePassingSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttributePassingSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAttributePassingSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAttributePassingSettingResponseBody) SetRequestId(v string) *UpdateAttributePassingSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAttributePassingSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
