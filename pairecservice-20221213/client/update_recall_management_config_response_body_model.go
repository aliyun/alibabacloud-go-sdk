// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecallManagementConfigResponseBody
	GetRequestId() *string
}

type UpdateRecallManagementConfigResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecallManagementConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecallManagementConfigResponseBody) SetRequestId(v string) *UpdateRecallManagementConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecallManagementConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
