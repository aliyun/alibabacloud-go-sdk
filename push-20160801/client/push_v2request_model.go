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
	// This parameter is required.
	//
	// example:
	//
	// ****530646
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// c8016d13-6e76-****-9bda-769383d11787
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
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
