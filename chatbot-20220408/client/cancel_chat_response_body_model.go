// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCancelResult(v bool) *CancelChatResponseBody
	GetCancelResult() *bool
	SetRequestId(v string) *CancelChatResponseBody
	GetRequestId() *string
}

type CancelChatResponseBody struct {
	CancelResult *bool `json:"CancelResult,omitempty" xml:"CancelResult,omitempty"`
	// example:
	//
	// E3E5C779-A630-45AC-B0F2-A4506A4212F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelChatResponseBody) GoString() string {
	return s.String()
}

func (s *CancelChatResponseBody) GetCancelResult() *bool {
	return s.CancelResult
}

func (s *CancelChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelChatResponseBody) SetCancelResult(v bool) *CancelChatResponseBody {
	s.CancelResult = &v
	return s
}

func (s *CancelChatResponseBody) SetRequestId(v string) *CancelChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelChatResponseBody) Validate() error {
	return dara.Validate(s)
}
