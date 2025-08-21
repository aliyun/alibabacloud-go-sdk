// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushStatByMsgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *QueryPushStatByMsgRequest
	GetAppKey() *int64
	SetMessageId(v int64) *QueryPushStatByMsgRequest
	GetMessageId() *int64
}

type QueryPushStatByMsgRequest struct {
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
	// 510427
	MessageId *int64 `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s QueryPushStatByMsgRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPushStatByMsgRequest) GoString() string {
	return s.String()
}

func (s *QueryPushStatByMsgRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *QueryPushStatByMsgRequest) GetMessageId() *int64 {
	return s.MessageId
}

func (s *QueryPushStatByMsgRequest) SetAppKey(v int64) *QueryPushStatByMsgRequest {
	s.AppKey = &v
	return s
}

func (s *QueryPushStatByMsgRequest) SetMessageId(v int64) *QueryPushStatByMsgRequest {
	s.MessageId = &v
	return s
}

func (s *QueryPushStatByMsgRequest) Validate() error {
	return dara.Validate(s)
}
