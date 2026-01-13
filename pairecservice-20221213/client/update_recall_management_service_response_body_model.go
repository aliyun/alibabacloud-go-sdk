// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecallManagementServiceResponseBody
	GetRequestId() *string
}

type UpdateRecallManagementServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecallManagementServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecallManagementServiceResponseBody) SetRequestId(v string) *UpdateRecallManagementServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecallManagementServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
