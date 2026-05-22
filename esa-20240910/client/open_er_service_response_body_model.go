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
	// Id of the request
	//
	// example:
	//
	// 5CC228B4-7A67-4016-9C9F-4A4133494A91
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
