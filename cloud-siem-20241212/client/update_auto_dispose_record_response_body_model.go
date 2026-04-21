// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoDisposeRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAutoDisposeRecordResponseBody
	GetRequestId() *string
}

type UpdateAutoDisposeRecordResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAutoDisposeRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoDisposeRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoDisposeRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAutoDisposeRecordResponseBody) SetRequestId(v string) *UpdateAutoDisposeRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutoDisposeRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
