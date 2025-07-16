// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProjectResponseBody
	GetRequestId() *string
}

type DeleteProjectResponseBody struct {
	// example:
	//
	// 0DA35264-0877-5852-8971-7735B547C969
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
