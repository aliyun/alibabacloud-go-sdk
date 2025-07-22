// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventSubscribeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteEventSubscribeRequest
	GetAppId() *string
	SetOwnerId(v int64) *DeleteEventSubscribeRequest
	GetOwnerId() *int64
	SetSubscribeId(v string) *DeleteEventSubscribeRequest
	GetSubscribeId() *string
}

type DeleteEventSubscribeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 9qb1****
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ad53276431c****
	SubscribeId *string `json:"SubscribeId,omitempty" xml:"SubscribeId,omitempty"`
}

func (s DeleteEventSubscribeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventSubscribeRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventSubscribeRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteEventSubscribeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteEventSubscribeRequest) GetSubscribeId() *string {
	return s.SubscribeId
}

func (s *DeleteEventSubscribeRequest) SetAppId(v string) *DeleteEventSubscribeRequest {
	s.AppId = &v
	return s
}

func (s *DeleteEventSubscribeRequest) SetOwnerId(v int64) *DeleteEventSubscribeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteEventSubscribeRequest) SetSubscribeId(v string) *DeleteEventSubscribeRequest {
	s.SubscribeId = &v
	return s
}

func (s *DeleteEventSubscribeRequest) Validate() error {
	return dara.Validate(s)
}
