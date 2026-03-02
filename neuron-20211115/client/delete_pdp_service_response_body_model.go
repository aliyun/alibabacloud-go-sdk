// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePdpServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePdpServiceResponseBody
	GetRequestId() *string
}

type DeletePdpServiceResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePdpServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePdpServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePdpServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePdpServiceResponseBody) SetRequestId(v string) *DeletePdpServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePdpServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
