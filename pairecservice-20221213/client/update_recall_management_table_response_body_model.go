// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecallManagementTableResponseBody
	GetRequestId() *string
}

type UpdateRecallManagementTableResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecallManagementTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecallManagementTableResponseBody) SetRequestId(v string) *UpdateRecallManagementTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecallManagementTableResponseBody) Validate() error {
	return dara.Validate(s)
}
