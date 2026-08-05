// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCredentialsResponseBody
	GetRequestId() *string
}

type DeleteCredentialsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2423C841-91C4-5E51-B296-590D367967FC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCredentialsResponseBody) SetRequestId(v string) *DeleteCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCredentialsResponseBody) Validate() error {
	return dara.Validate(s)
}
