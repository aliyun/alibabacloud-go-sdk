// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteContinuouslyPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *CompleteContinuouslyPushRequest
	GetAppKey() *int64
	SetMessageId(v string) *CompleteContinuouslyPushRequest
	GetMessageId() *string
}

type CompleteContinuouslyPushRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4010290149170430
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s CompleteContinuouslyPushRequest) String() string {
	return dara.Prettify(s)
}

func (s CompleteContinuouslyPushRequest) GoString() string {
	return s.String()
}

func (s *CompleteContinuouslyPushRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *CompleteContinuouslyPushRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *CompleteContinuouslyPushRequest) SetAppKey(v int64) *CompleteContinuouslyPushRequest {
	s.AppKey = &v
	return s
}

func (s *CompleteContinuouslyPushRequest) SetMessageId(v string) *CompleteContinuouslyPushRequest {
	s.MessageId = &v
	return s
}

func (s *CompleteContinuouslyPushRequest) Validate() error {
	return dara.Validate(s)
}
