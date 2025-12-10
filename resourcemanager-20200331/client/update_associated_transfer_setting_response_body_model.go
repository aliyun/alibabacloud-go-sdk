// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssociatedTransferSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAssociatedTransferSettingResponseBody
	GetRequestId() *string
}

type UpdateAssociatedTransferSettingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAssociatedTransferSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssociatedTransferSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAssociatedTransferSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAssociatedTransferSettingResponseBody) SetRequestId(v string) *UpdateAssociatedTransferSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAssociatedTransferSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
