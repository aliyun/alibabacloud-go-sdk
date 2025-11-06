// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMaxYearOfServerLockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxYear(v int32) *CheckMaxYearOfServerLockResponseBody
	GetMaxYear() *int32
	SetRequestId(v string) *CheckMaxYearOfServerLockResponseBody
	GetRequestId() *string
}

type CheckMaxYearOfServerLockResponseBody struct {
	// example:
	//
	// 10
	MaxYear *int32 `json:"MaxYear,omitempty" xml:"MaxYear,omitempty"`
	// example:
	//
	// 9DFCF6F8-243C-****-8035-4B12FEFD7D48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckMaxYearOfServerLockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckMaxYearOfServerLockResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMaxYearOfServerLockResponseBody) GetMaxYear() *int32 {
	return s.MaxYear
}

func (s *CheckMaxYearOfServerLockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckMaxYearOfServerLockResponseBody) SetMaxYear(v int32) *CheckMaxYearOfServerLockResponseBody {
	s.MaxYear = &v
	return s
}

func (s *CheckMaxYearOfServerLockResponseBody) SetRequestId(v string) *CheckMaxYearOfServerLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckMaxYearOfServerLockResponseBody) Validate() error {
	return dara.Validate(s)
}
