// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePCACertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePCACertResponseBody
	GetRequestId() *string
}

type DeletePCACertResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePCACertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePCACertResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePCACertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePCACertResponseBody) SetRequestId(v string) *DeletePCACertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePCACertResponseBody) Validate() error {
	return dara.Validate(s)
}
