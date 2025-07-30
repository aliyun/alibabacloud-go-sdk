// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *DescribeChannelAccountResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeChannelAccountResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeChannelAccountResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeChannelAccountResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeChannelAccountResponseBody
	GetHttpStatusCode() *int32
	SetPassword(v string) *DescribeChannelAccountResponseBody
	GetPassword() *string
	SetRequestId(v string) *DescribeChannelAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeChannelAccountResponseBody
	GetSuccess() *bool
	SetUsername(v string) *DescribeChannelAccountResponseBody
	GetUsername() *string
}

type DescribeChannelAccountResponseBody struct {
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeChannelAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelAccountResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeChannelAccountResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeChannelAccountResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeChannelAccountResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeChannelAccountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeChannelAccountResponseBody) GetPassword() *string {
	return s.Password
}

func (s *DescribeChannelAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeChannelAccountResponseBody) GetUsername() *string {
	return s.Username
}

func (s *DescribeChannelAccountResponseBody) SetDynamicCode(v string) *DescribeChannelAccountResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeChannelAccountResponseBody) SetDynamicMessage(v string) *DescribeChannelAccountResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeChannelAccountResponseBody) SetErrCode(v string) *DescribeChannelAccountResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeChannelAccountResponseBody) SetErrMessage(v string) *DescribeChannelAccountResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeChannelAccountResponseBody) SetHttpStatusCode(v int32) *DescribeChannelAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeChannelAccountResponseBody) SetPassword(v string) *DescribeChannelAccountResponseBody {
	s.Password = &v
	return s
}

func (s *DescribeChannelAccountResponseBody) SetRequestId(v string) *DescribeChannelAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelAccountResponseBody) SetSuccess(v bool) *DescribeChannelAccountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeChannelAccountResponseBody) SetUsername(v string) *DescribeChannelAccountResponseBody {
	s.Username = &v
	return s
}

func (s *DescribeChannelAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
