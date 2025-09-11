// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryShortUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryShortUrlResponseBody
	GetCode() *string
	SetData(v *QueryShortUrlResponseBodyData) *QueryShortUrlResponseBody
	GetData() *QueryShortUrlResponseBodyData
	SetMessage(v string) *QueryShortUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryShortUrlResponseBody
	GetRequestId() *string
}

type QueryShortUrlResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the short URL.
	Data *QueryShortUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryShortUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryShortUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryShortUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryShortUrlResponseBody) GetData() *QueryShortUrlResponseBodyData {
	return s.Data
}

func (s *QueryShortUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryShortUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryShortUrlResponseBody) SetCode(v string) *QueryShortUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryShortUrlResponseBody) SetData(v *QueryShortUrlResponseBodyData) *QueryShortUrlResponseBody {
	s.Data = v
	return s
}

func (s *QueryShortUrlResponseBody) SetMessage(v string) *QueryShortUrlResponseBody {
	s.Message = &v
	return s
}

func (s *QueryShortUrlResponseBody) SetRequestId(v string) *QueryShortUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryShortUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryShortUrlResponseBodyData struct {
	// The time when the short URL was created.
	//
	// example:
	//
	// 2019-01-08 16:44:13
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the short URL expires.
	//
	// example:
	//
	// 2019-01-22 11:21:11
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The PV.
	//
	// example:
	//
	// 300
	PageViewCount *string `json:"PageViewCount,omitempty" xml:"PageViewCount,omitempty"`
	// The short URL.
	//
	// example:
	//
	// http://****.cn/6y8uy7
	ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
	// The service name of the short URL.
	//
	// example:
	//
	// The Alibaba Cloud Short Link service.
	ShortUrlName *string `json:"ShortUrlName,omitempty" xml:"ShortUrlName,omitempty"`
	// The status of the short URL. Valid values:
	//
	// 	- **expired**
	//
	// 	- **effective**
	//
	// 	- **audit**
	//
	// 	- **reject**
	//
	// example:
	//
	// expired
	ShortUrlStatus *string `json:"ShortUrlStatus,omitempty" xml:"ShortUrlStatus,omitempty"`
	// The source address.
	//
	// example:
	//
	// https://www.****.com/product/sms
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	// The UV.
	//
	// example:
	//
	// 23
	UniqueVisitorCount *string `json:"UniqueVisitorCount,omitempty" xml:"UniqueVisitorCount,omitempty"`
}

func (s QueryShortUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryShortUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryShortUrlResponseBodyData) GetCreateDate() *string {
	return s.CreateDate
}

func (s *QueryShortUrlResponseBodyData) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *QueryShortUrlResponseBodyData) GetPageViewCount() *string {
	return s.PageViewCount
}

func (s *QueryShortUrlResponseBodyData) GetShortUrl() *string {
	return s.ShortUrl
}

func (s *QueryShortUrlResponseBodyData) GetShortUrlName() *string {
	return s.ShortUrlName
}

func (s *QueryShortUrlResponseBodyData) GetShortUrlStatus() *string {
	return s.ShortUrlStatus
}

func (s *QueryShortUrlResponseBodyData) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *QueryShortUrlResponseBodyData) GetUniqueVisitorCount() *string {
	return s.UniqueVisitorCount
}

func (s *QueryShortUrlResponseBodyData) SetCreateDate(v string) *QueryShortUrlResponseBodyData {
	s.CreateDate = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetExpireDate(v string) *QueryShortUrlResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetPageViewCount(v string) *QueryShortUrlResponseBodyData {
	s.PageViewCount = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetShortUrl(v string) *QueryShortUrlResponseBodyData {
	s.ShortUrl = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetShortUrlName(v string) *QueryShortUrlResponseBodyData {
	s.ShortUrlName = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetShortUrlStatus(v string) *QueryShortUrlResponseBodyData {
	s.ShortUrlStatus = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetSourceUrl(v string) *QueryShortUrlResponseBodyData {
	s.SourceUrl = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetUniqueVisitorCount(v string) *QueryShortUrlResponseBodyData {
	s.UniqueVisitorCount = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
