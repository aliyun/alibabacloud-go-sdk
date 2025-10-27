// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMassPushV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *MassPushV2Request
	GetAppKey() *int64
	SetIdempotentToken(v string) *MassPushV2Request
	GetIdempotentToken() *string
	SetPushTasks(v []*PushTask) *MassPushV2Request
	GetPushTasks() []*PushTask
}

type MassPushV2Request struct {
	// This parameter is required.
	AppKey          *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	IdempotentToken *string `json:"IdempotentToken,omitempty" xml:"IdempotentToken,omitempty"`
	// This parameter is required.
	PushTasks []*PushTask `json:"PushTasks,omitempty" xml:"PushTasks,omitempty" type:"Repeated"`
}

func (s MassPushV2Request) String() string {
	return dara.Prettify(s)
}

func (s MassPushV2Request) GoString() string {
	return s.String()
}

func (s *MassPushV2Request) GetAppKey() *int64 {
	return s.AppKey
}

func (s *MassPushV2Request) GetIdempotentToken() *string {
	return s.IdempotentToken
}

func (s *MassPushV2Request) GetPushTasks() []*PushTask {
	return s.PushTasks
}

func (s *MassPushV2Request) SetAppKey(v int64) *MassPushV2Request {
	s.AppKey = &v
	return s
}

func (s *MassPushV2Request) SetIdempotentToken(v string) *MassPushV2Request {
	s.IdempotentToken = &v
	return s
}

func (s *MassPushV2Request) SetPushTasks(v []*PushTask) *MassPushV2Request {
	s.PushTasks = v
	return s
}

func (s *MassPushV2Request) Validate() error {
	if s.PushTasks != nil {
		for _, item := range s.PushTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
