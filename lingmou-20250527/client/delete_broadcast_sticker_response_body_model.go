// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBroadcastStickerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteBroadcastStickerResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteBroadcastStickerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteBroadcastStickerResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBroadcastStickerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBroadcastStickerResponseBody
	GetSuccess() *bool
}

type DeleteBroadcastStickerResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 90C68329-A75C-5449-A928-4D0BAD7AA0FA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteBroadcastStickerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBroadcastStickerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBroadcastStickerResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBroadcastStickerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBroadcastStickerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBroadcastStickerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBroadcastStickerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBroadcastStickerResponseBody) SetCode(v string) *DeleteBroadcastStickerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBroadcastStickerResponseBody) SetHttpStatusCode(v int32) *DeleteBroadcastStickerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBroadcastStickerResponseBody) SetMessage(v string) *DeleteBroadcastStickerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBroadcastStickerResponseBody) SetRequestId(v string) *DeleteBroadcastStickerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBroadcastStickerResponseBody) SetSuccess(v bool) *DeleteBroadcastStickerResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBroadcastStickerResponseBody) Validate() error {
	return dara.Validate(s)
}
