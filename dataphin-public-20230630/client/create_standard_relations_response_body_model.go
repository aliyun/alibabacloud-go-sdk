// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateStandardRelationsResponseBody
	GetCode() *string
	SetData(v *CreateStandardRelationsResponseBodyData) *CreateStandardRelationsResponseBody
	GetData() *CreateStandardRelationsResponseBodyData
	SetHttpStatusCode(v int32) *CreateStandardRelationsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateStandardRelationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateStandardRelationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateStandardRelationsResponseBody
	GetSuccess() *bool
}

type CreateStandardRelationsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateStandardRelationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateStandardRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStandardRelationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateStandardRelationsResponseBody) GetData() *CreateStandardRelationsResponseBodyData {
	return s.Data
}

func (s *CreateStandardRelationsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateStandardRelationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateStandardRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStandardRelationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStandardRelationsResponseBody) SetCode(v string) *CreateStandardRelationsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStandardRelationsResponseBody) SetData(v *CreateStandardRelationsResponseBodyData) *CreateStandardRelationsResponseBody {
	s.Data = v
	return s
}

func (s *CreateStandardRelationsResponseBody) SetHttpStatusCode(v int32) *CreateStandardRelationsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateStandardRelationsResponseBody) SetMessage(v string) *CreateStandardRelationsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStandardRelationsResponseBody) SetRequestId(v string) *CreateStandardRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStandardRelationsResponseBody) SetSuccess(v bool) *CreateStandardRelationsResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStandardRelationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardRelationsResponseBodyData struct {
	NotExistStandardIdList []*int64 `json:"NotExistStandardIdList,omitempty" xml:"NotExistStandardIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s CreateStandardRelationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRelationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateStandardRelationsResponseBodyData) GetNotExistStandardIdList() []*int64 {
	return s.NotExistStandardIdList
}

func (s *CreateStandardRelationsResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *CreateStandardRelationsResponseBodyData) SetNotExistStandardIdList(v []*int64) *CreateStandardRelationsResponseBodyData {
	s.NotExistStandardIdList = v
	return s
}

func (s *CreateStandardRelationsResponseBodyData) SetSuccessCount(v int32) *CreateStandardRelationsResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *CreateStandardRelationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
