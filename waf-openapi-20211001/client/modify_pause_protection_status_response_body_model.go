// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPauseProtectionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPauseProtectionStatusResponseBody
	GetRequestId() *string
}

type ModifyPauseProtectionStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-*****-******-D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPauseProtectionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPauseProtectionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPauseProtectionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPauseProtectionStatusResponseBody) SetRequestId(v string) *ModifyPauseProtectionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPauseProtectionStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
