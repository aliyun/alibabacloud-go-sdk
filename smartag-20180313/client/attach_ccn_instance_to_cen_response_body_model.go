// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachCcnInstanceToCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachCcnInstanceToCenResponseBody
	GetRequestId() *string
}

type AttachCcnInstanceToCenResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 551CD836-9E46-4F2C-A167-B4363180A647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachCcnInstanceToCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachCcnInstanceToCenResponseBody) GoString() string {
	return s.String()
}

func (s *AttachCcnInstanceToCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachCcnInstanceToCenResponseBody) SetRequestId(v string) *AttachCcnInstanceToCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachCcnInstanceToCenResponseBody) Validate() error {
	return dara.Validate(s)
}
