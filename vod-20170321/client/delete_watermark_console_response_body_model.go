// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWatermarkConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWatermarkConsoleResponseBody
	GetRequestId() *string
}

type DeleteWatermarkConsoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWatermarkConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWatermarkConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWatermarkConsoleResponseBody) SetRequestId(v string) *DeleteWatermarkConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWatermarkConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}
