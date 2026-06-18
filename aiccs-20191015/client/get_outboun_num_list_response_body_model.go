// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOutbounNumListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOutbounNumListResponseBody
	GetCode() *string
	SetData(v *GetOutbounNumListResponseBodyData) *GetOutbounNumListResponseBody
	GetData() *GetOutbounNumListResponseBodyData
	SetHttpStatusCode(v int64) *GetOutbounNumListResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetOutbounNumListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOutbounNumListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOutbounNumListResponseBody
	GetSuccess() *bool
}

type GetOutbounNumListResponseBody struct {
	// Status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Information about the number list.
	Data *GetOutbounNumListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOutbounNumListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOutbounNumListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOutbounNumListResponseBody) GetData() *GetOutbounNumListResponseBodyData {
	return s.Data
}

func (s *GetOutbounNumListResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetOutbounNumListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOutbounNumListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOutbounNumListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOutbounNumListResponseBody) SetCode(v string) *GetOutbounNumListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetData(v *GetOutbounNumListResponseBodyData) *GetOutbounNumListResponseBody {
	s.Data = v
	return s
}

func (s *GetOutbounNumListResponseBody) SetHttpStatusCode(v int64) *GetOutbounNumListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetMessage(v string) *GetOutbounNumListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetRequestId(v string) *GetOutbounNumListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetSuccess(v bool) *GetOutbounNumListResponseBody {
	s.Success = &v
	return s
}

func (s *GetOutbounNumListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOutbounNumListResponseBodyData struct {
	// Caller number information.
	Num []*GetOutbounNumListResponseBodyDataNum `json:"Num,omitempty" xml:"Num,omitempty" type:"Repeated"`
	// Number group information.
	NumGroup []*GetOutbounNumListResponseBodyDataNumGroup `json:"NumGroup,omitempty" xml:"NumGroup,omitempty" type:"Repeated"`
}

func (s GetOutbounNumListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOutbounNumListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBodyData) GetNum() []*GetOutbounNumListResponseBodyDataNum {
	return s.Num
}

func (s *GetOutbounNumListResponseBodyData) GetNumGroup() []*GetOutbounNumListResponseBodyDataNumGroup {
	return s.NumGroup
}

func (s *GetOutbounNumListResponseBodyData) SetNum(v []*GetOutbounNumListResponseBodyDataNum) *GetOutbounNumListResponseBodyData {
	s.Num = v
	return s
}

func (s *GetOutbounNumListResponseBodyData) SetNumGroup(v []*GetOutbounNumListResponseBodyDataNumGroup) *GetOutbounNumListResponseBodyData {
	s.NumGroup = v
	return s
}

func (s *GetOutbounNumListResponseBodyData) Validate() error {
	if s.Num != nil {
		for _, item := range s.Num {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NumGroup != nil {
		for _, item := range s.NumGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOutbounNumListResponseBodyDataNum struct {
	// Description of the caller number (geographic location information).
	//
	// example:
	//
	// 浙江省杭州市
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Caller number type. Valid values:
	//
	// - **1**: Number.
	//
	// - **2**: Number group.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// Caller number.
	//
	// example:
	//
	// 07512234****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetOutbounNumListResponseBodyDataNum) String() string {
	return dara.Prettify(s)
}

func (s GetOutbounNumListResponseBodyDataNum) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBodyDataNum) GetDescription() *string {
	return s.Description
}

func (s *GetOutbounNumListResponseBodyDataNum) GetType() *int32 {
	return s.Type
}

func (s *GetOutbounNumListResponseBodyDataNum) GetValue() *string {
	return s.Value
}

func (s *GetOutbounNumListResponseBodyDataNum) SetDescription(v string) *GetOutbounNumListResponseBodyDataNum {
	s.Description = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNum) SetType(v int32) *GetOutbounNumListResponseBodyDataNum {
	s.Type = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNum) SetValue(v string) *GetOutbounNumListResponseBodyDataNum {
	s.Value = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNum) Validate() error {
	return dara.Validate(s)
}

type GetOutbounNumListResponseBodyDataNumGroup struct {
	// Number group description (number group name).
	//
	// example:
	//
	// Jella
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Number group type. Valid values:
	//
	// - **1**: Number.
	//
	// - **2**: Number group.
	//
	// example:
	//
	// 2
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// Number group ID.
	//
	// example:
	//
	// 7688****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetOutbounNumListResponseBodyDataNumGroup) String() string {
	return dara.Prettify(s)
}

func (s GetOutbounNumListResponseBodyDataNumGroup) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) GetDescription() *string {
	return s.Description
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) GetType() *int32 {
	return s.Type
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) GetValue() *string {
	return s.Value
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) SetDescription(v string) *GetOutbounNumListResponseBodyDataNumGroup {
	s.Description = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) SetType(v int32) *GetOutbounNumListResponseBodyDataNumGroup {
	s.Type = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) SetValue(v string) *GetOutbounNumListResponseBodyDataNumGroup {
	s.Value = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) Validate() error {
	return dara.Validate(s)
}
