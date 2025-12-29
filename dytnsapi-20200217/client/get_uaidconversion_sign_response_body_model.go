// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUAIDConversionSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetUAIDConversionSignResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetUAIDConversionSignResponseBody
	GetCode() *string
	SetData(v *GetUAIDConversionSignResponseBodyData) *GetUAIDConversionSignResponseBody
	GetData() *GetUAIDConversionSignResponseBodyData
	SetMessage(v string) *GetUAIDConversionSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUAIDConversionSignResponseBody
	GetRequestId() *string
}

type GetUAIDConversionSignResponseBody struct {
	AccessDeniedDetail *string                                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *GetUAIDConversionSignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUAIDConversionSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUAIDConversionSignResponseBody) GoString() string {
	return s.String()
}

func (s *GetUAIDConversionSignResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetUAIDConversionSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUAIDConversionSignResponseBody) GetData() *GetUAIDConversionSignResponseBodyData {
	return s.Data
}

func (s *GetUAIDConversionSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUAIDConversionSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUAIDConversionSignResponseBody) SetAccessDeniedDetail(v string) *GetUAIDConversionSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetUAIDConversionSignResponseBody) SetCode(v string) *GetUAIDConversionSignResponseBody {
	s.Code = &v
	return s
}

func (s *GetUAIDConversionSignResponseBody) SetData(v *GetUAIDConversionSignResponseBodyData) *GetUAIDConversionSignResponseBody {
	s.Data = v
	return s
}

func (s *GetUAIDConversionSignResponseBody) SetMessage(v string) *GetUAIDConversionSignResponseBody {
	s.Message = &v
	return s
}

func (s *GetUAIDConversionSignResponseBody) SetRequestId(v string) *GetUAIDConversionSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUAIDConversionSignResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUAIDConversionSignResponseBodyData struct {
	// example:
	//
	// 示例值示例值示例值
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值示例值
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// example:
	//
	// 示例值
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s GetUAIDConversionSignResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUAIDConversionSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUAIDConversionSignResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *GetUAIDConversionSignResponseBodyData) GetOutId() *string {
	return s.OutId
}

func (s *GetUAIDConversionSignResponseBodyData) GetSign() *string {
	return s.Sign
}

func (s *GetUAIDConversionSignResponseBodyData) SetCarrier(v string) *GetUAIDConversionSignResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *GetUAIDConversionSignResponseBodyData) SetOutId(v string) *GetUAIDConversionSignResponseBodyData {
	s.OutId = &v
	return s
}

func (s *GetUAIDConversionSignResponseBodyData) SetSign(v string) *GetUAIDConversionSignResponseBodyData {
	s.Sign = &v
	return s
}

func (s *GetUAIDConversionSignResponseBodyData) Validate() error {
	return dara.Validate(s)
}
