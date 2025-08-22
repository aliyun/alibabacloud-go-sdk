// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnSpecificConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnSpecificConfigResponseBody
	GetRequestId() *string
}

type DeleteDcdnSpecificConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnSpecificConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnSpecificConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSpecificConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnSpecificConfigResponseBody) SetRequestId(v string) *DeleteDcdnSpecificConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnSpecificConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
