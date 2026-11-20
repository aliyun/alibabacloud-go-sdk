// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindAXBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UnBindAXBResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UnBindAXBResponseBody
	GetCode() *string
	SetData(v *UnBindAXBResponseBodyData) *UnBindAXBResponseBody
	GetData() *UnBindAXBResponseBodyData
	SetMessage(v string) *UnBindAXBResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnBindAXBResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnBindAXBResponseBody
	GetSuccess() *bool
}

type UnBindAXBResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned status code. Valid values:
	//
	// - 0000: Success.
	//
	// - Other values: Failure.
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of dissociating the AXB call binding.
	//
	// example:
	//
	// {"Code":"0000","Message":"Success","Success":true}
	Data *UnBindAXBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4A2568E1-7D83-5491-8DEF-600D8AFF97E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: Success.
	//
	// - false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnBindAXBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnBindAXBResponseBody) GoString() string {
	return s.String()
}

func (s *UnBindAXBResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UnBindAXBResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnBindAXBResponseBody) GetData() *UnBindAXBResponseBodyData {
	return s.Data
}

func (s *UnBindAXBResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnBindAXBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnBindAXBResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnBindAXBResponseBody) SetAccessDeniedDetail(v string) *UnBindAXBResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnBindAXBResponseBody) SetCode(v string) *UnBindAXBResponseBody {
	s.Code = &v
	return s
}

func (s *UnBindAXBResponseBody) SetData(v *UnBindAXBResponseBodyData) *UnBindAXBResponseBody {
	s.Data = v
	return s
}

func (s *UnBindAXBResponseBody) SetMessage(v string) *UnBindAXBResponseBody {
	s.Message = &v
	return s
}

func (s *UnBindAXBResponseBody) SetRequestId(v string) *UnBindAXBResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnBindAXBResponseBody) SetSuccess(v bool) *UnBindAXBResponseBody {
	s.Success = &v
	return s
}

func (s *UnBindAXBResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnBindAXBResponseBodyData struct {
	// The returned status code. A value of 0000 indicates success. Other values indicate failure.
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request is successful. A value of true indicates success, and a value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnBindAXBResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UnBindAXBResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnBindAXBResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *UnBindAXBResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *UnBindAXBResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UnBindAXBResponseBodyData) SetCode(v string) *UnBindAXBResponseBodyData {
	s.Code = &v
	return s
}

func (s *UnBindAXBResponseBodyData) SetMessage(v string) *UnBindAXBResponseBodyData {
	s.Message = &v
	return s
}

func (s *UnBindAXBResponseBodyData) SetSuccess(v bool) *UnBindAXBResponseBodyData {
	s.Success = &v
	return s
}

func (s *UnBindAXBResponseBodyData) Validate() error {
	return dara.Validate(s)
}
