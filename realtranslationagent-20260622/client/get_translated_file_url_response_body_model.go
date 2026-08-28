// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslatedFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTranslatedFileUrlResponseBody
	GetCode() *string
	SetData(v *GetTranslatedFileUrlResponseBodyData) *GetTranslatedFileUrlResponseBody
	GetData() *GetTranslatedFileUrlResponseBodyData
	SetMessage(v string) *GetTranslatedFileUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTranslatedFileUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTranslatedFileUrlResponseBody
	GetSuccess() *bool
}

type GetTranslatedFileUrlResponseBody struct {
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *GetTranslatedFileUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 18D108E8-9625-5A26-BF0C-23EA0A2646B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTranslatedFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTranslatedFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranslatedFileUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTranslatedFileUrlResponseBody) GetData() *GetTranslatedFileUrlResponseBodyData {
	return s.Data
}

func (s *GetTranslatedFileUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTranslatedFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranslatedFileUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTranslatedFileUrlResponseBody) SetCode(v string) *GetTranslatedFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBody) SetData(v *GetTranslatedFileUrlResponseBodyData) *GetTranslatedFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetTranslatedFileUrlResponseBody) SetMessage(v string) *GetTranslatedFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBody) SetRequestId(v string) *GetTranslatedFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBody) SetSuccess(v bool) *GetTranslatedFileUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranslatedFileUrlResponseBodyData struct {
	// The download URL of the file.
	//
	// example:
	//
	// https://gtm-translate-service-prod.oss-cn-zhangjiakou.aliyuncs.com/translations/replaced/a_file_992736840.pptx?Expires=1782738716&OSSAccessKeyId=STS.NZm6TgFWU2sgpBxs2UD154B9w&Signature=uPORIIsYjiwRFzoyNUN8Htnwdfc%3D&security-token=CAIS3gJ1q6Ft5B2yfSjIr5nYfe7Tq4h0hbGMcmTJlzIASL4Z24eSlTz2IHhMf3FtAuwft%2FU0mWFW7foSlrp6SJtIXleCZtF94oxN9h2gb4fb4woBanWd08%2FLI3OaLjKm9u2wCryLYbGwU%2FOpbE%2B%2B5U0X6LDmdDKkckW4OJmS8%2FBOZcgWWQ%2FKBlgvRq0hRG1YpdQdKGHaONu0LxfumRCwNkdzvRdmgm4NgsbWgO%2Fks0OC1ACnmrdM%2FdupesL0MPMBZskvD42Hu8VtbbfE3SJq7BxHybx7lqQs%2B02c5onGWwQKv0zfYrGJo4M0cF9jLqcmHutYtvH6jvlxpuGWjInt1RdGMKRHXj7YAZy63dDYCHRtm2ect12R0R3spTPvXvGd22tMCfkrqw7Ahz2PACvRGM5dh0AbW042tZHwaHNHYcJrPu9YH1QLobvGc7TkCSYBIdG7lRJ8EPtayyu0U3F2gASJGhqAATkIHCzHZWyr%2F8WJFcinsIxuI8iYoH3pOTZ2HJgDW38Zbu0NqMZmH%2BbRahVPiN7s4ckYmF50hCbgXt%2BdD9R6jB%2BkNufrVtN%2FAVgRMLOeEn2FZF0CCLlyjaIqq8QIxaPpDSLxpV6wZSh5enenoKm%2B1wunWRdd0gwG03k%2F4RCcDbb%2FIAA%3D
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The expiration time. Unit: milliseconds.
	//
	// example:
	//
	// 1774147442
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// translated_a_file.pptx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetTranslatedFileUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTranslatedFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTranslatedFileUrlResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetTranslatedFileUrlResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetTranslatedFileUrlResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetTranslatedFileUrlResponseBodyData) SetDownloadUrl(v string) *GetTranslatedFileUrlResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBodyData) SetExpireTime(v int64) *GetTranslatedFileUrlResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBodyData) SetFileName(v string) *GetTranslatedFileUrlResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
