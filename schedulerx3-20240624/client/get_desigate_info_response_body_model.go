// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDesigateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetDesigateInfoResponseBody
	GetCode() *int32
	SetData(v *GetDesigateInfoResponseBodyData) *GetDesigateInfoResponseBody
	GetData() *GetDesigateInfoResponseBodyData
	SetMessage(v string) *GetDesigateInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDesigateInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDesigateInfoResponseBody
	GetSuccess() *bool
}

type GetDesigateInfoResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetDesigateInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1DF6732E-15D8-5E1F-95E3-C10077F556B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDesigateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDesigateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDesigateInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetDesigateInfoResponseBody) GetData() *GetDesigateInfoResponseBodyData {
	return s.Data
}

func (s *GetDesigateInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDesigateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDesigateInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDesigateInfoResponseBody) SetCode(v int32) *GetDesigateInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDesigateInfoResponseBody) SetData(v *GetDesigateInfoResponseBodyData) *GetDesigateInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDesigateInfoResponseBody) SetMessage(v string) *GetDesigateInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDesigateInfoResponseBody) SetRequestId(v string) *GetDesigateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDesigateInfoResponseBody) SetSuccess(v bool) *GetDesigateInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetDesigateInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDesigateInfoResponseBodyData struct {
	// example:
	//
	// 2
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	// example:
	//
	// true
	Transferable *bool `json:"Transferable,omitempty" xml:"Transferable,omitempty"`
}

func (s GetDesigateInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDesigateInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDesigateInfoResponseBodyData) GetDesignateType() *int32 {
	return s.DesignateType
}

func (s *GetDesigateInfoResponseBodyData) GetTransferable() *bool {
	return s.Transferable
}

func (s *GetDesigateInfoResponseBodyData) SetDesignateType(v int32) *GetDesigateInfoResponseBodyData {
	s.DesignateType = &v
	return s
}

func (s *GetDesigateInfoResponseBodyData) SetTransferable(v bool) *GetDesigateInfoResponseBodyData {
	s.Transferable = &v
	return s
}

func (s *GetDesigateInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
