// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPlaybookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPlaybookResponseBody
	GetRequestId() *string
}

type ModifyPlaybookResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B584F84-D66A-5525-8E7B-05612A903ABF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPlaybookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPlaybookResponseBody) SetRequestId(v string) *ModifyPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPlaybookResponseBody) Validate() error {
	return dara.Validate(s)
}
