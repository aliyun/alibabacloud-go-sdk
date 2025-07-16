// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCdnSubTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCdnSubTaskResponseBody
	GetRequestId() *string
}

type DeleteCdnSubTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCdnSubTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCdnSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCdnSubTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCdnSubTaskResponseBody) SetRequestId(v string) *DeleteCdnSubTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCdnSubTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
