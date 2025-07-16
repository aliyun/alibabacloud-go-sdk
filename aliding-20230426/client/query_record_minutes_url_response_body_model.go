// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordMinutesUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordMinutesUrls(v []*QueryRecordMinutesUrlResponseBodyRecordMinutesUrls) *QueryRecordMinutesUrlResponseBody
	GetRecordMinutesUrls() []*QueryRecordMinutesUrlResponseBodyRecordMinutesUrls
	SetRequestId(v string) *QueryRecordMinutesUrlResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *QueryRecordMinutesUrlResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *QueryRecordMinutesUrlResponseBody
	GetVendorType() *string
}

type QueryRecordMinutesUrlResponseBody struct {
	RecordMinutesUrls []*QueryRecordMinutesUrlResponseBodyRecordMinutesUrls `json:"recordMinutesUrls,omitempty" xml:"recordMinutesUrls,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s QueryRecordMinutesUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordMinutesUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlResponseBody) GetRecordMinutesUrls() []*QueryRecordMinutesUrlResponseBodyRecordMinutesUrls {
	return s.RecordMinutesUrls
}

func (s *QueryRecordMinutesUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRecordMinutesUrlResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *QueryRecordMinutesUrlResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *QueryRecordMinutesUrlResponseBody) SetRecordMinutesUrls(v []*QueryRecordMinutesUrlResponseBodyRecordMinutesUrls) *QueryRecordMinutesUrlResponseBody {
	s.RecordMinutesUrls = v
	return s
}

func (s *QueryRecordMinutesUrlResponseBody) SetRequestId(v string) *QueryRecordMinutesUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordMinutesUrlResponseBody) SetVendorRequestId(v string) *QueryRecordMinutesUrlResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *QueryRecordMinutesUrlResponseBody) SetVendorType(v string) *QueryRecordMinutesUrlResponseBody {
	s.VendorType = &v
	return s
}

func (s *QueryRecordMinutesUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryRecordMinutesUrlResponseBodyRecordMinutesUrls struct {
	// example:
	//
	// url
	RecordMinutesUrl *string `json:"RecordMinutesUrl,omitempty" xml:"RecordMinutesUrl,omitempty"`
}

func (s QueryRecordMinutesUrlResponseBodyRecordMinutesUrls) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordMinutesUrlResponseBodyRecordMinutesUrls) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlResponseBodyRecordMinutesUrls) GetRecordMinutesUrl() *string {
	return s.RecordMinutesUrl
}

func (s *QueryRecordMinutesUrlResponseBodyRecordMinutesUrls) SetRecordMinutesUrl(v string) *QueryRecordMinutesUrlResponseBodyRecordMinutesUrls {
	s.RecordMinutesUrl = &v
	return s
}

func (s *QueryRecordMinutesUrlResponseBodyRecordMinutesUrls) Validate() error {
	return dara.Validate(s)
}
