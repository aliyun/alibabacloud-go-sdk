// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDynamicDictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeywordList(v []*string) *GenerateDynamicDictResponseBody
	GetKeywordList() []*string
	SetRequestId(v string) *GenerateDynamicDictResponseBody
	GetRequestId() *string
}

type GenerateDynamicDictResponseBody struct {
	// The custom weak passwords.
	KeywordList []*string `json:"KeywordList,omitempty" xml:"KeywordList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 00E9B912-6066-5E4E-9F24-35EA09F2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateDynamicDictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateDynamicDictResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDynamicDictResponseBody) GetKeywordList() []*string {
	return s.KeywordList
}

func (s *GenerateDynamicDictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateDynamicDictResponseBody) SetKeywordList(v []*string) *GenerateDynamicDictResponseBody {
	s.KeywordList = v
	return s
}

func (s *GenerateDynamicDictResponseBody) SetRequestId(v string) *GenerateDynamicDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDynamicDictResponseBody) Validate() error {
	return dara.Validate(s)
}
