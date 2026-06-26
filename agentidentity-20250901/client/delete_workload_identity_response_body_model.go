// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkloadIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkloadIdentityResponseBody
	GetRequestId() *string
}

type DeleteWorkloadIdentityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWorkloadIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkloadIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkloadIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkloadIdentityResponseBody) SetRequestId(v string) *DeleteWorkloadIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkloadIdentityResponseBody) Validate() error {
	return dara.Validate(s)
}
