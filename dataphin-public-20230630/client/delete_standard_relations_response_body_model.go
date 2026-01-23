// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteStandardRelationsResponseBody
	GetCode() *string
	SetData(v *DeleteStandardRelationsResponseBodyData) *DeleteStandardRelationsResponseBody
	GetData() *DeleteStandardRelationsResponseBodyData
	SetHttpStatusCode(v int32) *DeleteStandardRelationsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteStandardRelationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteStandardRelationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteStandardRelationsResponseBody
	GetSuccess() *bool
}

type DeleteStandardRelationsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteStandardRelationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DeleteStandardRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStandardRelationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteStandardRelationsResponseBody) GetData() *DeleteStandardRelationsResponseBodyData {
	return s.Data
}

func (s *DeleteStandardRelationsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteStandardRelationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteStandardRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStandardRelationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteStandardRelationsResponseBody) SetCode(v string) *DeleteStandardRelationsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStandardRelationsResponseBody) SetData(v *DeleteStandardRelationsResponseBodyData) *DeleteStandardRelationsResponseBody {
	s.Data = v
	return s
}

func (s *DeleteStandardRelationsResponseBody) SetHttpStatusCode(v int32) *DeleteStandardRelationsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteStandardRelationsResponseBody) SetMessage(v string) *DeleteStandardRelationsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStandardRelationsResponseBody) SetRequestId(v string) *DeleteStandardRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStandardRelationsResponseBody) SetSuccess(v bool) *DeleteStandardRelationsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteStandardRelationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteStandardRelationsResponseBodyData struct {
	NotExistStandardIdList []*int64 `json:"NotExistStandardIdList,omitempty" xml:"NotExistStandardIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s DeleteStandardRelationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardRelationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteStandardRelationsResponseBodyData) GetNotExistStandardIdList() []*int64 {
	return s.NotExistStandardIdList
}

func (s *DeleteStandardRelationsResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *DeleteStandardRelationsResponseBodyData) SetNotExistStandardIdList(v []*int64) *DeleteStandardRelationsResponseBodyData {
	s.NotExistStandardIdList = v
	return s
}

func (s *DeleteStandardRelationsResponseBodyData) SetSuccessCount(v int32) *DeleteStandardRelationsResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *DeleteStandardRelationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
