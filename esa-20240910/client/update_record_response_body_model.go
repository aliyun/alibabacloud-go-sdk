// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecordResponseBody
	GetRequestId() *string
}

type UpdateRecordResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecordResponseBody) SetRequestId(v string) *UpdateRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
