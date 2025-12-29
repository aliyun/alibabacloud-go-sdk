// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUAIDApplyTokenSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetUAIDApplyTokenSignResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetUAIDApplyTokenSignResponseBody
	GetCode() *string
	SetData(v *GetUAIDApplyTokenSignResponseBodyData) *GetUAIDApplyTokenSignResponseBody
	GetData() *GetUAIDApplyTokenSignResponseBodyData
	SetMessage(v string) *GetUAIDApplyTokenSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUAIDApplyTokenSignResponseBody
	GetRequestId() *string
}

type GetUAIDApplyTokenSignResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetUAIDApplyTokenSignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD******177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUAIDApplyTokenSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUAIDApplyTokenSignResponseBody) GoString() string {
	return s.String()
}

func (s *GetUAIDApplyTokenSignResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetUAIDApplyTokenSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUAIDApplyTokenSignResponseBody) GetData() *GetUAIDApplyTokenSignResponseBodyData {
	return s.Data
}

func (s *GetUAIDApplyTokenSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUAIDApplyTokenSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUAIDApplyTokenSignResponseBody) SetAccessDeniedDetail(v string) *GetUAIDApplyTokenSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBody) SetCode(v string) *GetUAIDApplyTokenSignResponseBody {
	s.Code = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBody) SetData(v *GetUAIDApplyTokenSignResponseBodyData) *GetUAIDApplyTokenSignResponseBody {
	s.Data = v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBody) SetMessage(v string) *GetUAIDApplyTokenSignResponseBody {
	s.Message = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBody) SetRequestId(v string) *GetUAIDApplyTokenSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUAIDApplyTokenSignResponseBodyData struct {
	// example:
	//
	// CM
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// b8b5b3a*******0b9893484fdf412c99
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// example:
	//
	// D2E204D74EEB373E468632********23F592C4C9
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s GetUAIDApplyTokenSignResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUAIDApplyTokenSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUAIDApplyTokenSignResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *GetUAIDApplyTokenSignResponseBodyData) GetOutId() *string {
	return s.OutId
}

func (s *GetUAIDApplyTokenSignResponseBodyData) GetSign() *string {
	return s.Sign
}

func (s *GetUAIDApplyTokenSignResponseBodyData) SetCarrier(v string) *GetUAIDApplyTokenSignResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBodyData) SetOutId(v string) *GetUAIDApplyTokenSignResponseBodyData {
	s.OutId = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBodyData) SetSign(v string) *GetUAIDApplyTokenSignResponseBodyData {
	s.Sign = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponseBodyData) Validate() error {
	return dara.Validate(s)
}
