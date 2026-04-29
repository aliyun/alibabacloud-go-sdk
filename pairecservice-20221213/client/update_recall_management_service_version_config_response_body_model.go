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
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
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
