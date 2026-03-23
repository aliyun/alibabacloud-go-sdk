// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApSsidConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeleteApSsidConfigResponseBody
	GetData() *string
	SetErrorCode(v int32) *DeleteApSsidConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *DeleteApSsidConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *DeleteApSsidConfigResponseBody
	GetIsSuccess() *bool
}

type DeleteApSsidConfigResponseBody struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s DeleteApSsidConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApSsidConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApSsidConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteApSsidConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *DeleteApSsidConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteApSsidConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteApSsidConfigResponseBody) SetData(v string) *DeleteApSsidConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteApSsidConfigResponseBody) SetErrorCode(v int32) *DeleteApSsidConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteApSsidConfigResponseBody) SetErrorMessage(v string) *DeleteApSsidConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteApSsidConfigResponseBody) SetIsSuccess(v bool) *DeleteApSsidConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteApSsidConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
