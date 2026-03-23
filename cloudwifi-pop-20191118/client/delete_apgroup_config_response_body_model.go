// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApgroupConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeleteApgroupConfigResponseBody
	GetData() *string
	SetErrorCode(v int32) *DeleteApgroupConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *DeleteApgroupConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *DeleteApgroupConfigResponseBody
	GetIsSuccess() *bool
}

type DeleteApgroupConfigResponseBody struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s DeleteApgroupConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApgroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApgroupConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteApgroupConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *DeleteApgroupConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteApgroupConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteApgroupConfigResponseBody) SetData(v string) *DeleteApgroupConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteApgroupConfigResponseBody) SetErrorCode(v int32) *DeleteApgroupConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteApgroupConfigResponseBody) SetErrorMessage(v string) *DeleteApgroupConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteApgroupConfigResponseBody) SetIsSuccess(v bool) *DeleteApgroupConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteApgroupConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
