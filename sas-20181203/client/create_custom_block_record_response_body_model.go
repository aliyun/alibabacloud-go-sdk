// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomBlockRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCustomBlockRecordResponseBody
	GetRequestId() *string
}

type CreateCustomBlockRecordResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 30CBF632-109F-596F-97F2-451C8B2A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomBlockRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomBlockRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomBlockRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomBlockRecordResponseBody) SetRequestId(v string) *CreateCustomBlockRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomBlockRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
