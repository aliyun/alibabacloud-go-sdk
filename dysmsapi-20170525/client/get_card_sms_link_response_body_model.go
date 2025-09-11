// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardSmsLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCardSmsLinkResponseBody
	GetCode() *string
	SetData(v *GetCardSmsLinkResponseBodyData) *GetCardSmsLinkResponseBody
	GetData() *GetCardSmsLinkResponseBodyData
	SetRequestId(v string) *GetCardSmsLinkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCardSmsLinkResponseBody
	GetSuccess() *bool
}

type GetCardSmsLinkResponseBody struct {
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetCardSmsLinkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CC89A90C-978F-46AC-B80D-54738371E7CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCardSmsLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCardSmsLinkResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardSmsLinkResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCardSmsLinkResponseBody) GetData() *GetCardSmsLinkResponseBodyData {
	return s.Data
}

func (s *GetCardSmsLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCardSmsLinkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCardSmsLinkResponseBody) SetCode(v string) *GetCardSmsLinkResponseBody {
	s.Code = &v
	return s
}

func (s *GetCardSmsLinkResponseBody) SetData(v *GetCardSmsLinkResponseBodyData) *GetCardSmsLinkResponseBody {
	s.Data = v
	return s
}

func (s *GetCardSmsLinkResponseBody) SetRequestId(v string) *GetCardSmsLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCardSmsLinkResponseBody) SetSuccess(v bool) *GetCardSmsLinkResponseBody {
	s.Success = &v
	return s
}

func (s *GetCardSmsLinkResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCardSmsLinkResponseBodyData struct {
	// The mobile phone numbers that support card messages.
	//
	// example:
	//
	// [\\"1390000****\\",\\"1370000****\\"]
	CardPhoneNumbers *string `json:"CardPhoneNumbers,omitempty" xml:"CardPhoneNumbers,omitempty"`
	// The signatures must correspond to the mobile numbers and short URLs in sequence.
	//
	// example:
	//
	// ["aliyun","aliyun2"]
	CardSignNames *string `json:"CardSignNames,omitempty" xml:"CardSignNames,omitempty"`
	// The short URLs.
	//
	// example:
	//
	// [\\"mw2m.cn/LAaGGa\\",\\"mw2m.cn/LAAaes\\"]
	CardSmsLinks *string `json:"CardSmsLinks,omitempty" xml:"CardSmsLinks,omitempty"`
	// The review status of the card message template.
	//
	// 	- **0**: pending approval
	//
	// 	- **1**: approved
	//
	// 	- **2**: rejected
	//
	// > Unapproved card messages are rolled back.
	//
	// example:
	//
	// 0
	CardTmpState *int32 `json:"CardTmpState,omitempty" xml:"CardTmpState,omitempty"`
	// The mobile phone numbers that do not support card messages.
	//
	// example:
	//
	// 1390000****
	NotMediaMobiles *string `json:"NotMediaMobiles,omitempty" xml:"NotMediaMobiles,omitempty"`
}

func (s GetCardSmsLinkResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCardSmsLinkResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCardSmsLinkResponseBodyData) GetCardPhoneNumbers() *string {
	return s.CardPhoneNumbers
}

func (s *GetCardSmsLinkResponseBodyData) GetCardSignNames() *string {
	return s.CardSignNames
}

func (s *GetCardSmsLinkResponseBodyData) GetCardSmsLinks() *string {
	return s.CardSmsLinks
}

func (s *GetCardSmsLinkResponseBodyData) GetCardTmpState() *int32 {
	return s.CardTmpState
}

func (s *GetCardSmsLinkResponseBodyData) GetNotMediaMobiles() *string {
	return s.NotMediaMobiles
}

func (s *GetCardSmsLinkResponseBodyData) SetCardPhoneNumbers(v string) *GetCardSmsLinkResponseBodyData {
	s.CardPhoneNumbers = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) SetCardSignNames(v string) *GetCardSmsLinkResponseBodyData {
	s.CardSignNames = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) SetCardSmsLinks(v string) *GetCardSmsLinkResponseBodyData {
	s.CardSmsLinks = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) SetCardTmpState(v int32) *GetCardSmsLinkResponseBodyData {
	s.CardTmpState = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) SetNotMediaMobiles(v string) *GetCardSmsLinkResponseBodyData {
	s.NotMediaMobiles = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) Validate() error {
	return dara.Validate(s)
}
