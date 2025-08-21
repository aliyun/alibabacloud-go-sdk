// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *CancelPushRequest
	GetAppKey() *int64
	SetMessageId(v int64) *CancelPushRequest
	GetMessageId() *int64
}

type CancelPushRequest struct {
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
	// 501029
	MessageId *int64 `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s CancelPushRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelPushRequest) GoString() string {
	return s.String()
}

func (s *CancelPushRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *CancelPushRequest) GetMessageId() *int64 {
	return s.MessageId
}

func (s *CancelPushRequest) SetAppKey(v int64) *CancelPushRequest {
	s.AppKey = &v
	return s
}

func (s *CancelPushRequest) SetMessageId(v int64) *CancelPushRequest {
	s.MessageId = &v
	return s
}

func (s *CancelPushRequest) Validate() error {
	return dara.Validate(s)
}
