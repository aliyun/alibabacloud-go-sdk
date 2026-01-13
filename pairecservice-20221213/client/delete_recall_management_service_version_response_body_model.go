// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecallManagementServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRecallManagementServiceVersionResponseBody
	GetRequestId() *string
}

type DeleteRecallManagementServiceVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecallManagementServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecallManagementServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecallManagementServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecallManagementServiceVersionResponseBody) SetRequestId(v string) *DeleteRecallManagementServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecallManagementServiceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
