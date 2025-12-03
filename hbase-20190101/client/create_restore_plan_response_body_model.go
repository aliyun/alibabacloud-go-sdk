// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestorePlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRestorePlanResponseBody
	GetRequestId() *string
}

type CreateRestorePlanResponseBody struct {
	// example:
	//
	// A0598673-EB6E-4F6D-9961-E0F2012090C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRestorePlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRestorePlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRestorePlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRestorePlanResponseBody) SetRequestId(v string) *CreateRestorePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRestorePlanResponseBody) Validate() error {
	return dara.Validate(s)
}
