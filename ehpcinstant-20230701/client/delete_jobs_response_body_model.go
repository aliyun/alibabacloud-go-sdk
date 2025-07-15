// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteJobsResponseBody
	GetRequestId() *string
}

type DeleteJobsResponseBody struct {
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJobsResponseBody) SetRequestId(v string) *DeleteJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
