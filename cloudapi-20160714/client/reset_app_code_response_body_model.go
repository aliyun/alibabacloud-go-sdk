// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAppCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetAppCodeResponseBody
	GetRequestId() *string
}

type ResetAppCodeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D1B18FFE-4A81-59D8-AA02-1817098977CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAppCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAppCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAppCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAppCodeResponseBody) SetRequestId(v string) *ResetAppCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAppCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
