// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultWatermarkConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultWatermarkConsoleResponseBody
	GetRequestId() *string
}

type SetDefaultWatermarkConsoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultWatermarkConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultWatermarkConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultWatermarkConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultWatermarkConsoleResponseBody) SetRequestId(v string) *SetDefaultWatermarkConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultWatermarkConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}
