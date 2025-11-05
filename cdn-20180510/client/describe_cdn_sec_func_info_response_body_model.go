// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSecFuncInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*DescribeCdnSecFuncInfoResponseBodyContent) *DescribeCdnSecFuncInfoResponseBody
	GetContent() []*DescribeCdnSecFuncInfoResponseBodyContent
	SetDescription(v string) *DescribeCdnSecFuncInfoResponseBody
	GetDescription() *string
	SetHttpStatus(v string) *DescribeCdnSecFuncInfoResponseBody
	GetHttpStatus() *string
	SetRequestId(v string) *DescribeCdnSecFuncInfoResponseBody
	GetRequestId() *string
	SetRetCode(v string) *DescribeCdnSecFuncInfoResponseBody
	GetRetCode() *string
}

type DescribeCdnSecFuncInfoResponseBody struct {
	// Queried data.
	Content []*DescribeCdnSecFuncInfoResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatus *string `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BCD7D917-76F1-442F-BB75-C810DE34C761
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP request response code.
	//
	// 	- 0: OK.
	//
	// 	- Values other than 0: an error.
	//
	// example:
	//
	// 0
	RetCode *string `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s DescribeCdnSecFuncInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSecFuncInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnSecFuncInfoResponseBody) GetContent() []*DescribeCdnSecFuncInfoResponseBodyContent {
	return s.Content
}

func (s *DescribeCdnSecFuncInfoResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeCdnSecFuncInfoResponseBody) GetHttpStatus() *string {
	return s.HttpStatus
}

func (s *DescribeCdnSecFuncInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnSecFuncInfoResponseBody) GetRetCode() *string {
	return s.RetCode
}

func (s *DescribeCdnSecFuncInfoResponseBody) SetContent(v []*DescribeCdnSecFuncInfoResponseBodyContent) *DescribeCdnSecFuncInfoResponseBody {
	s.Content = v
	return s
}

func (s *DescribeCdnSecFuncInfoResponseBody) SetDescription(v string) *DescribeCdnSecFuncInfoResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeCdnSecFuncInfoResponseBody) SetHttpStatus(v string) *DescribeCdnSecFuncInfoResponseBody {
	s.HttpStatus = &v
	return s
}

func (s *DescribeCdnSecFuncInfoResponseBody) SetRequestId(v string) *DescribeCdnSecFuncInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnSecFuncInfoResponseBody) SetRetCode(v string) *DescribeCdnSecFuncInfoResponseBody {
	s.RetCode = &v
	return s
}

func (s *DescribeCdnSecFuncInfoResponseBody) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnSecFuncInfoResponseBodyContent struct {
	// The tag.
	//
	// example:
	//
	// TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The value.
	//
	// example:
	//
	// TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCdnSecFuncInfoResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSecFuncInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeCdnSecFuncInfoResponseBodyContent) GetLabel() *string {
	return s.Label
}

func (s *DescribeCdnSecFuncInfoResponseBodyContent) GetValue() *string {
	return s.Value
}

func (s *DescribeCdnSecFuncInfoResponseBodyContent) SetLabel(v string) *DescribeCdnSecFuncInfoResponseBodyContent {
	s.Label = &v
	return s
}

func (s *DescribeCdnSecFuncInfoResponseBodyContent) SetValue(v string) *DescribeCdnSecFuncInfoResponseBodyContent {
	s.Value = &v
	return s
}

func (s *DescribeCdnSecFuncInfoResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
