// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRCInstanceDescriptionResponseBody
	GetRequestId() *string
}

type ModifyRCInstanceDescriptionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CCECD3CD-AB2D-4F6D-BEDE-47BC90A398D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCInstanceDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyRCInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCInstanceDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
