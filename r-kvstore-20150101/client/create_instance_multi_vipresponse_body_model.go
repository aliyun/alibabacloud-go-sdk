// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceMultiVIPResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateInstanceMultiVIPResponseBody
	GetRequestId() *string
}

type CreateInstanceMultiVIPResponseBody struct {
	// example:
	//
	// 52D901ED-E0A5-42FB-B9DB-39C295C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceMultiVIPResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceMultiVIPResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceMultiVIPResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceMultiVIPResponseBody) SetRequestId(v string) *CreateInstanceMultiVIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceMultiVIPResponseBody) Validate() error {
	return dara.Validate(s)
}
