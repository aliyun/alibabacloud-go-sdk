// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackApplicationResponseBody
	GetRequestId() *string
}

type RollbackApplicationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackApplicationResponseBody) SetRequestId(v string) *RollbackApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
