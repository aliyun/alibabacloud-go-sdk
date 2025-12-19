// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiatePptCreationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitiatePptCreationResponseBody
	GetCode() *string
	SetData(v *InitiatePptCreationResponseBodyData) *InitiatePptCreationResponseBody
	GetData() *InitiatePptCreationResponseBodyData
	SetHttpStatusCode(v int32) *InitiatePptCreationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *InitiatePptCreationResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitiatePptCreationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InitiatePptCreationResponseBody
	GetSuccess() *bool
}

type InitiatePptCreationResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *InitiatePptCreationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 11AC01F1-88FB-5C4D-B6F5-E8BB136CD5A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InitiatePptCreationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitiatePptCreationResponseBody) GoString() string {
	return s.String()
}

func (s *InitiatePptCreationResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitiatePptCreationResponseBody) GetData() *InitiatePptCreationResponseBodyData {
	return s.Data
}

func (s *InitiatePptCreationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *InitiatePptCreationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitiatePptCreationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitiatePptCreationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InitiatePptCreationResponseBody) SetCode(v string) *InitiatePptCreationResponseBody {
	s.Code = &v
	return s
}

func (s *InitiatePptCreationResponseBody) SetData(v *InitiatePptCreationResponseBodyData) *InitiatePptCreationResponseBody {
	s.Data = v
	return s
}

func (s *InitiatePptCreationResponseBody) SetHttpStatusCode(v int32) *InitiatePptCreationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InitiatePptCreationResponseBody) SetMessage(v string) *InitiatePptCreationResponseBody {
	s.Message = &v
	return s
}

func (s *InitiatePptCreationResponseBody) SetRequestId(v string) *InitiatePptCreationResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitiatePptCreationResponseBody) SetSuccess(v bool) *InitiatePptCreationResponseBody {
	s.Success = &v
	return s
}

func (s *InitiatePptCreationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitiatePptCreationResponseBodyData struct {
	// AppKey
	//
	// example:
	//
	// S1X5ecouBztZelaQ
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// Code
	//
	// example:
	//
	// 7dhqd2
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s InitiatePptCreationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InitiatePptCreationResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitiatePptCreationResponseBodyData) GetAppKey() *string {
	return s.AppKey
}

func (s *InitiatePptCreationResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *InitiatePptCreationResponseBodyData) SetAppKey(v string) *InitiatePptCreationResponseBodyData {
	s.AppKey = &v
	return s
}

func (s *InitiatePptCreationResponseBodyData) SetCode(v string) *InitiatePptCreationResponseBodyData {
	s.Code = &v
	return s
}

func (s *InitiatePptCreationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
