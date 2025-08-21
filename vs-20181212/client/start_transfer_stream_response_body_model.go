// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTransferStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartTransferStreamResponseBody
	GetRequestId() *string
}

type StartTransferStreamResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartTransferStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartTransferStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StartTransferStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartTransferStreamResponseBody) SetRequestId(v string) *StartTransferStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTransferStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
