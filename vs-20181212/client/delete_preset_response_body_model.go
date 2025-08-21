// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePresetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeletePresetResponseBody
	GetId() *string
	SetRequestId(v string) *DeletePresetResponseBody
	GetRequestId() *string
}

type DeletePresetResponseBody struct {
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// F18FD685-B194-4489-9609-F80A9490A258
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePresetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePresetResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePresetResponseBody) GetId() *string {
	return s.Id
}

func (s *DeletePresetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePresetResponseBody) SetId(v string) *DeletePresetResponseBody {
	s.Id = &v
	return s
}

func (s *DeletePresetResponseBody) SetRequestId(v string) *DeletePresetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePresetResponseBody) Validate() error {
	return dara.Validate(s)
}
