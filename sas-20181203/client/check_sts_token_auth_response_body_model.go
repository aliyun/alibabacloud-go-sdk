// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckStsTokenAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CheckStsTokenAuthResponseBody
	GetData() *int64
	SetRequestId(v string) *CheckStsTokenAuthResponseBody
	GetRequestId() *string
}

type CheckStsTokenAuthResponseBody struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 185685871307****
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 898F7AA7-CECD-5EC7-AF4D-664C601B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckStsTokenAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckStsTokenAuthResponseBody) GoString() string {
	return s.String()
}

func (s *CheckStsTokenAuthResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CheckStsTokenAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckStsTokenAuthResponseBody) SetData(v int64) *CheckStsTokenAuthResponseBody {
	s.Data = &v
	return s
}

func (s *CheckStsTokenAuthResponseBody) SetRequestId(v string) *CheckStsTokenAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckStsTokenAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
