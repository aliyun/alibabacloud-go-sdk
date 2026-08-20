// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *PushV2Request
	GetAppKey() *int64
	SetIdempotentToken(v string) *PushV2Request
	GetIdempotentToken() *string
	SetPushTask(v *PushTask) *PushV2Request
	GetPushTask() *PushTask
}

type PushV2Request struct {
	// The AppKey information.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****530646
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// An idempotence parameter that prevents duplicate push notifications caused by retries from the API caller. If you use the same IdempotentToken within 15 minutes, only one push notification is sent, and subsequent calls return the result of the first successful push.
	//
	// >
	//
	// > - The parameter format is a standard 36-character UUID (8-4-4-4-12). Each valid character is a hexadecimal digit in the range 0-9 or a-f, case-insensitive.
	//
	// > - This parameter only prevents duplicate push notifications caused by retries. It cannot prevent duplicate push notifications caused by concurrent calls.
	//
	// example:
	//
	// c8016d13-6e76-****-9bda-769383d11787
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// The push task.
	//
	// This parameter is required.
	PushTask *PushTask `json:"PushTask,omitempty" xml:"PushTask,omitempty"`
}

func (s PushV2Request) String() string {
	return dara.Prettify(s)
}

func (s PushV2Request) GoString() string {
	return s.String()
}

func (s *PushV2Request) GetAppKey() *int64 {
	return s.AppKey
}

func (s *PushV2Request) GetIdempotentToken() *string {
	return s.IdempotentToken
}

func (s *PushV2Request) GetPushTask() *PushTask {
	return s.PushTask
}

func (s *PushV2Request) SetAppKey(v int64) *PushV2Request {
	s.AppKey = &v
	return s
}

func (s *PushV2Request) SetIdempotentToken(v string) *PushV2Request {
	s.IdempotentToken = &v
	return s
}

func (s *PushV2Request) SetPushTask(v *PushTask) *PushV2Request {
	s.PushTask = v
	return s
}

func (s *PushV2Request) Validate() error {
	if s.PushTask != nil {
		if err := s.PushTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}
