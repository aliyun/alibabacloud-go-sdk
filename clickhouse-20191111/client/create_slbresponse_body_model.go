// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSLBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSLBResponseBody
	GetRequestId() *string
}

type CreateSLBResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9A23C87D-87DF-4DA0-A50E-CB13F4F7923D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSLBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSLBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSLBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSLBResponseBody) SetRequestId(v string) *CreateSLBResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSLBResponseBody) Validate() error {
	return dara.Validate(s)
}
