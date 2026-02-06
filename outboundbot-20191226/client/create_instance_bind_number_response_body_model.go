// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceBindNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstanceBindNumberResponseBody
	GetCode() *string
	SetData(v *CreateInstanceBindNumberResponseBodyData) *CreateInstanceBindNumberResponseBody
	GetData() *CreateInstanceBindNumberResponseBodyData
	SetHttpStatusCode(v int32) *CreateInstanceBindNumberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateInstanceBindNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstanceBindNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceBindNumberResponseBody
	GetSuccess() *bool
}

type CreateInstanceBindNumberResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateInstanceBindNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceBindNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceBindNumberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceBindNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceBindNumberResponseBody) GetData() *CreateInstanceBindNumberResponseBodyData {
	return s.Data
}

func (s *CreateInstanceBindNumberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateInstanceBindNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstanceBindNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceBindNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceBindNumberResponseBody) SetCode(v string) *CreateInstanceBindNumberResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceBindNumberResponseBody) SetData(v *CreateInstanceBindNumberResponseBodyData) *CreateInstanceBindNumberResponseBody {
	s.Data = v
	return s
}

func (s *CreateInstanceBindNumberResponseBody) SetHttpStatusCode(v int32) *CreateInstanceBindNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceBindNumberResponseBody) SetMessage(v string) *CreateInstanceBindNumberResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceBindNumberResponseBody) SetRequestId(v string) *CreateInstanceBindNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceBindNumberResponseBody) SetSuccess(v bool) *CreateInstanceBindNumberResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceBindNumberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceBindNumberResponseBodyData struct {
	List []*CreateInstanceBindNumberResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CreateInstanceBindNumberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceBindNumberResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateInstanceBindNumberResponseBodyData) GetList() []*CreateInstanceBindNumberResponseBodyDataList {
	return s.List
}

func (s *CreateInstanceBindNumberResponseBodyData) SetList(v []*CreateInstanceBindNumberResponseBodyDataList) *CreateInstanceBindNumberResponseBodyData {
	s.List = v
	return s
}

func (s *CreateInstanceBindNumberResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateInstanceBindNumberResponseBodyDataList struct {
	// example:
	//
	// 96b847ad-2683-4794-b7b4-7ef094fb81f6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceBindNumberResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceBindNumberResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CreateInstanceBindNumberResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceBindNumberResponseBodyDataList) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceBindNumberResponseBodyDataList) SetInstanceId(v string) *CreateInstanceBindNumberResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceBindNumberResponseBodyDataList) SetSuccess(v bool) *CreateInstanceBindNumberResponseBodyDataList {
	s.Success = &v
	return s
}

func (s *CreateInstanceBindNumberResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
