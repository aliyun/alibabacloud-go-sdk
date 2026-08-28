// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginalFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOriginalFileUrlResponseBody
	GetCode() *string
	SetData(v *GetOriginalFileUrlResponseBodyData) *GetOriginalFileUrlResponseBody
	GetData() *GetOriginalFileUrlResponseBodyData
	SetMessage(v string) *GetOriginalFileUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOriginalFileUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOriginalFileUrlResponseBody
	GetSuccess() *bool
}

type GetOriginalFileUrlResponseBody struct {
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *GetOriginalFileUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 9F80A58B-DFBA-55A1-B9D2-819B32904127
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOriginalFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOriginalFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetOriginalFileUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOriginalFileUrlResponseBody) GetData() *GetOriginalFileUrlResponseBodyData {
	return s.Data
}

func (s *GetOriginalFileUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOriginalFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOriginalFileUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOriginalFileUrlResponseBody) SetCode(v string) *GetOriginalFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetOriginalFileUrlResponseBody) SetData(v *GetOriginalFileUrlResponseBodyData) *GetOriginalFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetOriginalFileUrlResponseBody) SetMessage(v string) *GetOriginalFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetOriginalFileUrlResponseBody) SetRequestId(v string) *GetOriginalFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOriginalFileUrlResponseBody) SetSuccess(v bool) *GetOriginalFileUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetOriginalFileUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOriginalFileUrlResponseBodyData struct {
	// The file download URL.
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
	// The file name.
	//
	// example:
	//
	// translated_a_file.pptx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetOriginalFileUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOriginalFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOriginalFileUrlResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetOriginalFileUrlResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetOriginalFileUrlResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetOriginalFileUrlResponseBodyData) SetDownloadUrl(v string) *GetOriginalFileUrlResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetOriginalFileUrlResponseBodyData) SetExpireTime(v int64) *GetOriginalFileUrlResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetOriginalFileUrlResponseBodyData) SetFileName(v string) *GetOriginalFileUrlResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetOriginalFileUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
