// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateStandardMappingResponseBody
	GetCode() *string
	SetData(v *CreateStandardMappingResponseBodyData) *CreateStandardMappingResponseBody
	GetData() *CreateStandardMappingResponseBodyData
	SetHttpStatusCode(v int32) *CreateStandardMappingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateStandardMappingResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateStandardMappingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateStandardMappingResponseBody
	GetSuccess() *bool
}

type CreateStandardMappingResponseBody struct {
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateStandardMappingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStandardMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardMappingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStandardMappingResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateStandardMappingResponseBody) GetData() *CreateStandardMappingResponseBodyData {
	return s.Data
}

func (s *CreateStandardMappingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateStandardMappingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateStandardMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStandardMappingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStandardMappingResponseBody) SetCode(v string) *CreateStandardMappingResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStandardMappingResponseBody) SetData(v *CreateStandardMappingResponseBodyData) *CreateStandardMappingResponseBody {
	s.Data = v
	return s
}

func (s *CreateStandardMappingResponseBody) SetHttpStatusCode(v int32) *CreateStandardMappingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateStandardMappingResponseBody) SetMessage(v string) *CreateStandardMappingResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStandardMappingResponseBody) SetRequestId(v string) *CreateStandardMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStandardMappingResponseBody) SetSuccess(v bool) *CreateStandardMappingResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStandardMappingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardMappingResponseBodyData struct {
	FailedGuidList []*string `json:"FailedGuidList,omitempty" xml:"FailedGuidList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s CreateStandardMappingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardMappingResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateStandardMappingResponseBodyData) GetFailedGuidList() []*string {
	return s.FailedGuidList
}

func (s *CreateStandardMappingResponseBodyData) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *CreateStandardMappingResponseBodyData) SetFailedGuidList(v []*string) *CreateStandardMappingResponseBodyData {
	s.FailedGuidList = v
	return s
}

func (s *CreateStandardMappingResponseBodyData) SetSuccessCount(v int64) *CreateStandardMappingResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *CreateStandardMappingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
