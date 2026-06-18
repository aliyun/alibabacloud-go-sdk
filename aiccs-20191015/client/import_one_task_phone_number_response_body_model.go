// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportOneTaskPhoneNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ImportOneTaskPhoneNumberResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ImportOneTaskPhoneNumberResponseBody
	GetCode() *string
	SetData(v *ImportOneTaskPhoneNumberResponseBodyData) *ImportOneTaskPhoneNumberResponseBody
	GetData() *ImportOneTaskPhoneNumberResponseBodyData
	SetMessage(v string) *ImportOneTaskPhoneNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportOneTaskPhoneNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportOneTaskPhoneNumberResponseBody
	GetSuccess() *bool
}

type ImportOneTaskPhoneNumberResponseBody struct {
	// The detailed reason why the access is denied.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ImportOneTaskPhoneNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the status code.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D6A51251-F7C4-596A-9F45-3C3219A5450D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - **true**: The API call was successful.
	//
	// - **false**: The API call failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportOneTaskPhoneNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportOneTaskPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ImportOneTaskPhoneNumberResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ImportOneTaskPhoneNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportOneTaskPhoneNumberResponseBody) GetData() *ImportOneTaskPhoneNumberResponseBodyData {
	return s.Data
}

func (s *ImportOneTaskPhoneNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportOneTaskPhoneNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportOneTaskPhoneNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportOneTaskPhoneNumberResponseBody) SetAccessDeniedDetail(v string) *ImportOneTaskPhoneNumberResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ImportOneTaskPhoneNumberResponseBody) SetCode(v string) *ImportOneTaskPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ImportOneTaskPhoneNumberResponseBody) SetData(v *ImportOneTaskPhoneNumberResponseBodyData) *ImportOneTaskPhoneNumberResponseBody {
	s.Data = v
	return s
}

func (s *ImportOneTaskPhoneNumberResponseBody) SetMessage(v string) *ImportOneTaskPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ImportOneTaskPhoneNumberResponseBody) SetRequestId(v string) *ImportOneTaskPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportOneTaskPhoneNumberResponseBody) SetSuccess(v bool) *ImportOneTaskPhoneNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ImportOneTaskPhoneNumberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportOneTaskPhoneNumberResponseBodyData struct {
	// The task detail ID.
	//
	// example:
	//
	// 92304322323*****
	DetailId *int64 `json:"DetailId,omitempty" xml:"DetailId,omitempty"`
}

func (s ImportOneTaskPhoneNumberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImportOneTaskPhoneNumberResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportOneTaskPhoneNumberResponseBodyData) GetDetailId() *int64 {
	return s.DetailId
}

func (s *ImportOneTaskPhoneNumberResponseBodyData) SetDetailId(v int64) *ImportOneTaskPhoneNumberResponseBodyData {
	s.DetailId = &v
	return s
}

func (s *ImportOneTaskPhoneNumberResponseBodyData) Validate() error {
	return dara.Validate(s)
}
