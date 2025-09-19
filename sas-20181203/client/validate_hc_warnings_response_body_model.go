// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateHcWarningsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidateHcWarningsResponseBody
	GetRequestId() *string
}

type ValidateHcWarningsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 722C4F88-7867-4E7B-8ADE-7451053
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidateHcWarningsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateHcWarningsResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateHcWarningsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateHcWarningsResponseBody) SetRequestId(v string) *ValidateHcWarningsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateHcWarningsResponseBody) Validate() error {
	return dara.Validate(s)
}
