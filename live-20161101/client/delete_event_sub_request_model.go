// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventSubRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteEventSubRequest
	GetAppId() *string
	SetSubscribeId(v string) *DeleteEventSubRequest
	GetSubscribeId() *string
}

type DeleteEventSubRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9qb1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The subscription ID. You can obtain the ID from the response to the [CreateEventSub](https://help.aliyun.com/document_detail/2848209.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ad53276431c****
	SubscribeId *string `json:"SubscribeId,omitempty" xml:"SubscribeId,omitempty"`
}

func (s DeleteEventSubRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventSubRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventSubRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteEventSubRequest) GetSubscribeId() *string {
	return s.SubscribeId
}

func (s *DeleteEventSubRequest) SetAppId(v string) *DeleteEventSubRequest {
	s.AppId = &v
	return s
}

func (s *DeleteEventSubRequest) SetSubscribeId(v string) *DeleteEventSubRequest {
	s.SubscribeId = &v
	return s
}

func (s *DeleteEventSubRequest) Validate() error {
	return dara.Validate(s)
}
