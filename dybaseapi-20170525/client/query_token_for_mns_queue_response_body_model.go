// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTokenForMnsQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTokenForMnsQueueResponseBody
	GetCode() *string
	SetMessage(v string) *QueryTokenForMnsQueueResponseBody
	GetMessage() *string
	SetMessageTokenDTO(v *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) *QueryTokenForMnsQueueResponseBody
	GetMessageTokenDTO() *QueryTokenForMnsQueueResponseBodyMessageTokenDTO
	SetRequestId(v string) *QueryTokenForMnsQueueResponseBody
	GetRequestId() *string
}

type QueryTokenForMnsQueueResponseBody struct {
	Code            *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	MessageTokenDTO *QueryTokenForMnsQueueResponseBodyMessageTokenDTO `json:"MessageTokenDTO,omitempty" xml:"MessageTokenDTO,omitempty" type:"Struct"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTokenForMnsQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTokenForMnsQueueResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTokenForMnsQueueResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTokenForMnsQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTokenForMnsQueueResponseBody) GetMessageTokenDTO() *QueryTokenForMnsQueueResponseBodyMessageTokenDTO {
	return s.MessageTokenDTO
}

func (s *QueryTokenForMnsQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTokenForMnsQueueResponseBody) SetCode(v string) *QueryTokenForMnsQueueResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBody) SetMessage(v string) *QueryTokenForMnsQueueResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBody) SetMessageTokenDTO(v *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) *QueryTokenForMnsQueueResponseBody {
	s.MessageTokenDTO = v
	return s
}

func (s *QueryTokenForMnsQueueResponseBody) SetRequestId(v string) *QueryTokenForMnsQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBody) Validate() error {
	if s.MessageTokenDTO != nil {
		if err := s.MessageTokenDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTokenForMnsQueueResponseBodyMessageTokenDTO struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime      *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s QueryTokenForMnsQueueResponseBodyMessageTokenDTO) String() string {
	return dara.Prettify(s)
}

func (s QueryTokenForMnsQueueResponseBodyMessageTokenDTO) GoString() string {
	return s.String()
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) SetAccessKeyId(v string) *QueryTokenForMnsQueueResponseBodyMessageTokenDTO {
	s.AccessKeyId = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) SetAccessKeySecret(v string) *QueryTokenForMnsQueueResponseBodyMessageTokenDTO {
	s.AccessKeySecret = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) SetCreateTime(v string) *QueryTokenForMnsQueueResponseBodyMessageTokenDTO {
	s.CreateTime = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) SetExpireTime(v string) *QueryTokenForMnsQueueResponseBodyMessageTokenDTO {
	s.ExpireTime = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) SetSecurityToken(v string) *QueryTokenForMnsQueueResponseBodyMessageTokenDTO {
	s.SecurityToken = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) Validate() error {
	return dara.Validate(s)
}
