// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteServiceRecordResponseBody
	GetRequestId() *string
}

type DeleteServiceRecordResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteServiceRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceRecordResponseBody) SetRequestId(v string) *DeleteServiceRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
