// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPlaybookInputOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPlaybookInputOutputResponseBody
	GetRequestId() *string
}

type ModifyPlaybookInputOutputResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8DDC07CE-D41B-5142-8D91-469462719C77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPlaybookInputOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlaybookInputOutputResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInputOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPlaybookInputOutputResponseBody) SetRequestId(v string) *ModifyPlaybookInputOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPlaybookInputOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
