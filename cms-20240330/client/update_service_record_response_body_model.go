// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateServiceRecordResponseBody
	GetRequestId() *string
}

type UpdateServiceRecordResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0CEC5375-XXXX-XXXX-XXXX-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateServiceRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceRecordResponseBody) SetRequestId(v string) *UpdateServiceRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
