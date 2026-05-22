// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenErServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenErServiceResponseBody
	GetRequestId() *string
}

type OpenErServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenErServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenErServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenErServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenErServiceResponseBody) SetRequestId(v string) *OpenErServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenErServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
