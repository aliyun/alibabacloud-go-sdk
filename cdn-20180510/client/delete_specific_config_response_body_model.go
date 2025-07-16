// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSpecificConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSpecificConfigResponseBody
	GetRequestId() *string
}

type DeleteSpecificConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSpecificConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSpecificConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSpecificConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSpecificConfigResponseBody) SetRequestId(v string) *DeleteSpecificConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSpecificConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
