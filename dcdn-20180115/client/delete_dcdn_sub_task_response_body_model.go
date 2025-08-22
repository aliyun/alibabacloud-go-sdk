// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnSubTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnSubTaskResponseBody
	GetRequestId() *string
}

type DeleteDcdnSubTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnSubTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSubTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnSubTaskResponseBody) SetRequestId(v string) *DeleteDcdnSubTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnSubTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
