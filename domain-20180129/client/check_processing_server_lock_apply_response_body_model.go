// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckProcessingServerLockApplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExists(v bool) *CheckProcessingServerLockApplyResponseBody
	GetExists() *bool
	SetRequestId(v string) *CheckProcessingServerLockApplyResponseBody
	GetRequestId() *string
}

type CheckProcessingServerLockApplyResponseBody struct {
	// example:
	//
	// true
	Exists *bool `json:"Exists,omitempty" xml:"Exists,omitempty"`
	// example:
	//
	// 9DFCF6F8-243C-****-8035-4B12FEFD7D48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckProcessingServerLockApplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckProcessingServerLockApplyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckProcessingServerLockApplyResponseBody) GetExists() *bool {
	return s.Exists
}

func (s *CheckProcessingServerLockApplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckProcessingServerLockApplyResponseBody) SetExists(v bool) *CheckProcessingServerLockApplyResponseBody {
	s.Exists = &v
	return s
}

func (s *CheckProcessingServerLockApplyResponseBody) SetRequestId(v string) *CheckProcessingServerLockApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckProcessingServerLockApplyResponseBody) Validate() error {
	return dara.Validate(s)
}
