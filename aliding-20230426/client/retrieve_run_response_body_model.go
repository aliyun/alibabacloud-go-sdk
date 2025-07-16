// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RetrieveRunResponseBody
	GetRequestId() *string
	SetRun(v *RetrieveRunResponseBodyRun) *RetrieveRunResponseBody
	GetRun() *RetrieveRunResponseBodyRun
}

type RetrieveRunResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Run       *RetrieveRunResponseBodyRun `json:"run,omitempty" xml:"run,omitempty" type:"Struct"`
}

func (s RetrieveRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetrieveRunResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetrieveRunResponseBody) GetRun() *RetrieveRunResponseBodyRun {
	return s.Run
}

func (s *RetrieveRunResponseBody) SetRequestId(v string) *RetrieveRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrieveRunResponseBody) SetRun(v *RetrieveRunResponseBodyRun) *RetrieveRunResponseBody {
	s.Run = v
	return s
}

func (s *RetrieveRunResponseBody) Validate() error {
	return dara.Validate(s)
}

type RetrieveRunResponseBodyRun struct {
	CancelledAt  *int64  `json:"cancelledAt,omitempty" xml:"cancelledAt,omitempty"`
	CompletedAt  *int64  `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	CreateAt     *int64  `json:"createAt,omitempty" xml:"createAt,omitempty"`
	ExpiresAt    *int64  `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	FailedAt     *int64  `json:"failedAt,omitempty" xml:"failedAt,omitempty"`
	Id           *string `json:"id,omitempty" xml:"id,omitempty"`
	LastErrorMsg *string `json:"lastErrorMsg,omitempty" xml:"lastErrorMsg,omitempty"`
	StartedAt    *int64  `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	ThreadId     *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s RetrieveRunResponseBodyRun) String() string {
	return dara.Prettify(s)
}

func (s RetrieveRunResponseBodyRun) GoString() string {
	return s.String()
}

func (s *RetrieveRunResponseBodyRun) GetCancelledAt() *int64 {
	return s.CancelledAt
}

func (s *RetrieveRunResponseBodyRun) GetCompletedAt() *int64 {
	return s.CompletedAt
}

func (s *RetrieveRunResponseBodyRun) GetCreateAt() *int64 {
	return s.CreateAt
}

func (s *RetrieveRunResponseBodyRun) GetExpiresAt() *int64 {
	return s.ExpiresAt
}

func (s *RetrieveRunResponseBodyRun) GetFailedAt() *int64 {
	return s.FailedAt
}

func (s *RetrieveRunResponseBodyRun) GetId() *string {
	return s.Id
}

func (s *RetrieveRunResponseBodyRun) GetLastErrorMsg() *string {
	return s.LastErrorMsg
}

func (s *RetrieveRunResponseBodyRun) GetStartedAt() *int64 {
	return s.StartedAt
}

func (s *RetrieveRunResponseBodyRun) GetStatus() *string {
	return s.Status
}

func (s *RetrieveRunResponseBodyRun) GetThreadId() *string {
	return s.ThreadId
}

func (s *RetrieveRunResponseBodyRun) SetCancelledAt(v int64) *RetrieveRunResponseBodyRun {
	s.CancelledAt = &v
	return s
}

func (s *RetrieveRunResponseBodyRun) SetCompletedAt(v int64) *RetrieveRunResponseBodyRun {
	s.CompletedAt = &v
	return s
}

func (s *RetrieveRunResponseBodyRun) SetCreateAt(v int64) *RetrieveRunResponseBodyRun {
	s.CreateAt = &v
	return s
}

func (s *RetrieveRunResponseBodyRun) SetExpiresAt(v int64) *RetrieveRunResponseBodyRun {
	s.ExpiresAt = &v
	return s
}

func (s *RetrieveRunResponseBodyRun) SetFailedAt(v int64) *RetrieveRunResponseBodyRun {
	s.FailedAt = &v
	return s
}

func (s *RetrieveRunResponseBodyRun) SetId(v string) *RetrieveRunResponseBodyRun {
	s.Id = &v
	return s
}

func (s *RetrieveRunResponseBodyRun) SetLastErrorMsg(v string) *RetrieveRunResponseBodyRun {
	s.LastErrorMsg = &v
	return s
}

func (s *RetrieveRunResponseBodyRun) SetStartedAt(v int64) *RetrieveRunResponseBodyRun {
	s.StartedAt = &v
	return s
}

func (s *RetrieveRunResponseBodyRun) SetStatus(v string) *RetrieveRunResponseBodyRun {
	s.Status = &v
	return s
}

func (s *RetrieveRunResponseBodyRun) SetThreadId(v string) *RetrieveRunResponseBodyRun {
	s.ThreadId = &v
	return s
}

func (s *RetrieveRunResponseBodyRun) Validate() error {
	return dara.Validate(s)
}
