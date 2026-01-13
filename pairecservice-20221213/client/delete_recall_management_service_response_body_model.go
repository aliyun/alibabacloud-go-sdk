// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecallManagementServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRecallManagementServiceResponseBody
	GetRequestId() *string
}

type DeleteRecallManagementServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecallManagementServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecallManagementServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecallManagementServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecallManagementServiceResponseBody) SetRequestId(v string) *DeleteRecallManagementServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecallManagementServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
