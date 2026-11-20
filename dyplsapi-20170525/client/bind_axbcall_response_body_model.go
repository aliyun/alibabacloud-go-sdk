// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAXBCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindAXBCallResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *BindAXBCallResponseBody
	GetCode() *string
	SetData(v *BindAXBCallResponseBodyData) *BindAXBCallResponseBody
	GetData() *BindAXBCallResponseBodyData
	SetMessage(v string) *BindAXBCallResponseBody
	GetMessage() *string
	SetSuccess(v bool) *BindAXBCallResponseBody
	GetSuccess() *bool
}

type BindAXBCallResponseBody struct {
	// The details of the access denied error.
	//
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code. Valid values:
	//
	// - 0000: Success.
	//
	// - Other values: Failure.
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The AXB call binding result.
	//
	// example:
	//
	// {"BindId":"476567566"}
	Data *BindAXBCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s BindAXBCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAXBCallResponseBody) GoString() string {
	return s.String()
}

func (s *BindAXBCallResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindAXBCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindAXBCallResponseBody) GetData() *BindAXBCallResponseBodyData {
	return s.Data
}

func (s *BindAXBCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindAXBCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindAXBCallResponseBody) SetAccessDeniedDetail(v string) *BindAXBCallResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindAXBCallResponseBody) SetCode(v string) *BindAXBCallResponseBody {
	s.Code = &v
	return s
}

func (s *BindAXBCallResponseBody) SetData(v *BindAXBCallResponseBodyData) *BindAXBCallResponseBody {
	s.Data = v
	return s
}

func (s *BindAXBCallResponseBody) SetMessage(v string) *BindAXBCallResponseBody {
	s.Message = &v
	return s
}

func (s *BindAXBCallResponseBody) SetSuccess(v bool) *BindAXBCallResponseBody {
	s.Success = &v
	return s
}

func (s *BindAXBCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAXBCallResponseBodyData struct {
	// The ID of the binding relationship.
	//
	// example:
	//
	// 476567566
	BindId *string `json:"BindId,omitempty" xml:"BindId,omitempty"`
}

func (s BindAXBCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindAXBCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAXBCallResponseBodyData) GetBindId() *string {
	return s.BindId
}

func (s *BindAXBCallResponseBodyData) SetBindId(v string) *BindAXBCallResponseBodyData {
	s.BindId = &v
	return s
}

func (s *BindAXBCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}
