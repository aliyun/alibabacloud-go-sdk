// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApisecAbnormalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApisecAbnormalsResponseBody
	GetRequestId() *string
}

type DeleteApisecAbnormalsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApisecAbnormalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApisecAbnormalsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApisecAbnormalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApisecAbnormalsResponseBody) SetRequestId(v string) *DeleteApisecAbnormalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApisecAbnormalsResponseBody) Validate() error {
	return dara.Validate(s)
}
