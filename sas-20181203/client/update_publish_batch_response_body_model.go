// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublishBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePublishBatchResponseBody
	GetRequestId() *string
}

type UpdatePublishBatchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3956048F-9D73-5EDB-834B-4827BB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePublishBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublishBatchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublishBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePublishBatchResponseBody) SetRequestId(v string) *UpdatePublishBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublishBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
