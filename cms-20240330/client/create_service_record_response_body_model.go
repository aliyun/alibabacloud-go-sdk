// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceRecordResponseBody
	GetRequestId() *string
}

type CreateServiceRecordResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateServiceRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceRecordResponseBody) SetRequestId(v string) *CreateServiceRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
