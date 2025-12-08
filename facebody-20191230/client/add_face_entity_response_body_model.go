// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddFaceEntityResponseBody
	GetRequestId() *string
}

type AddFaceEntityResponseBody struct {
	// example:
	//
	// DA7CAFEB-0A37-4496-9CDF-E3DBB6309CB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFaceEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFaceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFaceEntityResponseBody) SetRequestId(v string) *AddFaceEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
