// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageCodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodes(v []*Codes) *DetectImageCodesResponseBody
	GetCodes() []*Codes
	SetRequestId(v string) *DetectImageCodesResponseBody
	GetRequestId() *string
}

type DetectImageCodesResponseBody struct {
	// The barcodes or QR codes.
	//
	// This parameter is required.
	Codes []*Codes `json:"Codes,omitempty" xml:"Codes,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6E93D6C9-5AC0-49F9-914D-E02678D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageCodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectImageCodesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageCodesResponseBody) GetCodes() []*Codes {
	return s.Codes
}

func (s *DetectImageCodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectImageCodesResponseBody) SetCodes(v []*Codes) *DetectImageCodesResponseBody {
	s.Codes = v
	return s
}

func (s *DetectImageCodesResponseBody) SetRequestId(v string) *DetectImageCodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageCodesResponseBody) Validate() error {
	return dara.Validate(s)
}
