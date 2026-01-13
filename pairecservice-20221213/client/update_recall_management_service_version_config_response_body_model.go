// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementServiceVersionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecallManagementServiceVersionConfigResponseBody
	GetRequestId() *string
}

type UpdateRecallManagementServiceVersionConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecallManagementServiceVersionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceVersionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceVersionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecallManagementServiceVersionConfigResponseBody) SetRequestId(v string) *UpdateRecallManagementServiceVersionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
