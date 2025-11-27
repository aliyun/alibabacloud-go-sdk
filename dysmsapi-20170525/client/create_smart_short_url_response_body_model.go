// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmartShortUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSmartShortUrlResponseBody
	GetCode() *string
	SetMessage(v string) *CreateSmartShortUrlResponseBody
	GetMessage() *string
	SetModel(v []*CreateSmartShortUrlResponseBodyModel) *CreateSmartShortUrlResponseBody
	GetModel() []*CreateSmartShortUrlResponseBodyModel
	SetRequestId(v string) *CreateSmartShortUrlResponseBody
	GetRequestId() *string
}

type CreateSmartShortUrlResponseBody struct {
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   []*CreateSmartShortUrlResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSmartShortUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartShortUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmartShortUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSmartShortUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSmartShortUrlResponseBody) GetModel() []*CreateSmartShortUrlResponseBodyModel {
	return s.Model
}

func (s *CreateSmartShortUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSmartShortUrlResponseBody) SetCode(v string) *CreateSmartShortUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmartShortUrlResponseBody) SetMessage(v string) *CreateSmartShortUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmartShortUrlResponseBody) SetModel(v []*CreateSmartShortUrlResponseBodyModel) *CreateSmartShortUrlResponseBody {
	s.Model = v
	return s
}

func (s *CreateSmartShortUrlResponseBody) SetRequestId(v string) *CreateSmartShortUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSmartShortUrlResponseBody) Validate() error {
	if s.Model != nil {
		for _, item := range s.Model {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSmartShortUrlResponseBodyModel struct {
	// example:
	//
	// 示例值
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 11
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// example:
	//
	// 示例值
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// 示例值
	ShortName *string `json:"ShortName,omitempty" xml:"ShortName,omitempty"`
	// example:
	//
	// 示例值示例值
	ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
}

func (s CreateSmartShortUrlResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartShortUrlResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CreateSmartShortUrlResponseBodyModel) GetDomain() *string {
	return s.Domain
}

func (s *CreateSmartShortUrlResponseBodyModel) GetExpiration() *int64 {
	return s.Expiration
}

func (s *CreateSmartShortUrlResponseBodyModel) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CreateSmartShortUrlResponseBodyModel) GetShortName() *string {
	return s.ShortName
}

func (s *CreateSmartShortUrlResponseBodyModel) GetShortUrl() *string {
	return s.ShortUrl
}

func (s *CreateSmartShortUrlResponseBodyModel) SetDomain(v string) *CreateSmartShortUrlResponseBodyModel {
	s.Domain = &v
	return s
}

func (s *CreateSmartShortUrlResponseBodyModel) SetExpiration(v int64) *CreateSmartShortUrlResponseBodyModel {
	s.Expiration = &v
	return s
}

func (s *CreateSmartShortUrlResponseBodyModel) SetPhoneNumber(v string) *CreateSmartShortUrlResponseBodyModel {
	s.PhoneNumber = &v
	return s
}

func (s *CreateSmartShortUrlResponseBodyModel) SetShortName(v string) *CreateSmartShortUrlResponseBodyModel {
	s.ShortName = &v
	return s
}

func (s *CreateSmartShortUrlResponseBodyModel) SetShortUrl(v string) *CreateSmartShortUrlResponseBodyModel {
	s.ShortUrl = &v
	return s
}

func (s *CreateSmartShortUrlResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
