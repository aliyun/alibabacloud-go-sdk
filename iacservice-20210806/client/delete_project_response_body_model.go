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
	// BF72A6FB-B071-5F2E-A036-9D62545B962C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
