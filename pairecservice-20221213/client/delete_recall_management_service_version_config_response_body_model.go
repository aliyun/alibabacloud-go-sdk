// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecallManagementServiceVersionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRecallManagementServiceVersionConfigResponseBody
	GetRequestId() *string
}

type DeleteRecallManagementServiceVersionConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecallManagementServiceVersionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecallManagementServiceVersionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecallManagementServiceVersionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecallManagementServiceVersionConfigResponseBody) SetRequestId(v string) *DeleteRecallManagementServiceVersionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecallManagementServiceVersionConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
