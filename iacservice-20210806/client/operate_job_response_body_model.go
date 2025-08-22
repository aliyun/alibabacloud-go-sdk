// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateJobResponseBody
	GetRequestId() *string
}

type OperateJobResponseBody struct {
	// example:
	//
	// E602681C-A811-5787-9DC3-48BED7537071
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OperateJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateJobResponseBody) GoString() string {
	return s.String()
}

func (s *OperateJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateJobResponseBody) SetRequestId(v string) *OperateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateJobResponseBody) Validate() error {
	return dara.Validate(s)
}
