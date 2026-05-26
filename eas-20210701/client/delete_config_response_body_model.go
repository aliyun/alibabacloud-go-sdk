// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleted(v int32) *DeleteConfigResponseBody
	GetDeleted() *int32
	SetMessage(v string) *DeleteConfigResponseBody
	GetMessage() *string
}

type DeleteConfigResponseBody struct {
	// 删除的配置数量
	//
	// example:
	//
	// 1
	Deleted *int32 `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 操作结果消息
	//
	// example:
	//
	// Successfully deleted 1 configs
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeleteConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigResponseBody) GetDeleted() *int32 {
	return s.Deleted
}

func (s *DeleteConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteConfigResponseBody) SetDeleted(v int32) *DeleteConfigResponseBody {
	s.Deleted = &v
	return s
}

func (s *DeleteConfigResponseBody) SetMessage(v string) *DeleteConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
