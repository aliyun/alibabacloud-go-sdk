// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupWmInfoMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *LookupWmInfoMappingResponseBodyData) *LookupWmInfoMappingResponseBody
	GetData() *LookupWmInfoMappingResponseBodyData
	SetRequestId(v string) *LookupWmInfoMappingResponseBody
	GetRequestId() *string
}

type LookupWmInfoMappingResponseBody struct {
	Data *LookupWmInfoMappingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 7E9D7ACD-53D5-56EF-A913-79D148D06299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LookupWmInfoMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LookupWmInfoMappingResponseBody) GoString() string {
	return s.String()
}

func (s *LookupWmInfoMappingResponseBody) GetData() *LookupWmInfoMappingResponseBodyData {
	return s.Data
}

func (s *LookupWmInfoMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LookupWmInfoMappingResponseBody) SetData(v *LookupWmInfoMappingResponseBodyData) *LookupWmInfoMappingResponseBody {
	s.Data = v
	return s
}

func (s *LookupWmInfoMappingResponseBody) SetRequestId(v string) *LookupWmInfoMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *LookupWmInfoMappingResponseBody) Validate() error {
	return dara.Validate(s)
}

type LookupWmInfoMappingResponseBodyData struct {
	// example:
	//
	// aGVsbG8gc2F*****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
}

func (s LookupWmInfoMappingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s LookupWmInfoMappingResponseBodyData) GoString() string {
	return s.String()
}

func (s *LookupWmInfoMappingResponseBodyData) GetWmInfoBytesB64() *string {
	return s.WmInfoBytesB64
}

func (s *LookupWmInfoMappingResponseBodyData) SetWmInfoBytesB64(v string) *LookupWmInfoMappingResponseBodyData {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *LookupWmInfoMappingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
