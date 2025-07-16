// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchKeywordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArguments(v []interface{}) *CreateSearchKeywordResponseBody
	GetArguments() []interface{}
	SetRequestId(v string) *CreateSearchKeywordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSearchKeywordResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *CreateSearchKeywordResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CreateSearchKeywordResponseBody
	GetVendorType() *string
}

type CreateSearchKeywordResponseBody struct {
	// example:
	//
	// []
	Arguments []interface{} `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s CreateSearchKeywordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchKeywordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSearchKeywordResponseBody) GetArguments() []interface{} {
	return s.Arguments
}

func (s *CreateSearchKeywordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSearchKeywordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSearchKeywordResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CreateSearchKeywordResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CreateSearchKeywordResponseBody) SetArguments(v []interface{}) *CreateSearchKeywordResponseBody {
	s.Arguments = v
	return s
}

func (s *CreateSearchKeywordResponseBody) SetRequestId(v string) *CreateSearchKeywordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSearchKeywordResponseBody) SetSuccess(v bool) *CreateSearchKeywordResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSearchKeywordResponseBody) SetVendorRequestId(v string) *CreateSearchKeywordResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CreateSearchKeywordResponseBody) SetVendorType(v string) *CreateSearchKeywordResponseBody {
	s.VendorType = &v
	return s
}

func (s *CreateSearchKeywordResponseBody) Validate() error {
	return dara.Validate(s)
}
