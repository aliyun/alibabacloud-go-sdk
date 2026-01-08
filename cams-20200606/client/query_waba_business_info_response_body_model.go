// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWabaBusinessInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryWabaBusinessInfoResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryWabaBusinessInfoResponseBody
	GetCode() *string
	SetData(v *QueryWabaBusinessInfoResponseBodyData) *QueryWabaBusinessInfoResponseBody
	GetData() *QueryWabaBusinessInfoResponseBodyData
	SetMessage(v string) *QueryWabaBusinessInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryWabaBusinessInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryWabaBusinessInfoResponseBody
	GetSuccess() *bool
}

type QueryWabaBusinessInfoResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business information about the WABA.
	Data *QueryWabaBusinessInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// None.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryWabaBusinessInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryWabaBusinessInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWabaBusinessInfoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryWabaBusinessInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryWabaBusinessInfoResponseBody) GetData() *QueryWabaBusinessInfoResponseBodyData {
	return s.Data
}

func (s *QueryWabaBusinessInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryWabaBusinessInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryWabaBusinessInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryWabaBusinessInfoResponseBody) SetAccessDeniedDetail(v string) *QueryWabaBusinessInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBody) SetCode(v string) *QueryWabaBusinessInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBody) SetData(v *QueryWabaBusinessInfoResponseBodyData) *QueryWabaBusinessInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryWabaBusinessInfoResponseBody) SetMessage(v string) *QueryWabaBusinessInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBody) SetRequestId(v string) *QueryWabaBusinessInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBody) SetSuccess(v bool) *QueryWabaBusinessInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryWabaBusinessInfoResponseBodyData struct {
	// The Business Manager ID.
	//
	// example:
	//
	// 192882828733
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The Business Manager name.
	//
	// example:
	//
	// Alibaba
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// The verification status.
	//
	// example:
	//
	// verified
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
	// Deprecated
	//
	// The industry.
	//
	// example:
	//
	// Retail
	Vertical *string `json:"Vertical,omitempty" xml:"Vertical,omitempty"`
}

func (s QueryWabaBusinessInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryWabaBusinessInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryWabaBusinessInfoResponseBodyData) GetBusinessId() *string {
	return s.BusinessId
}

func (s *QueryWabaBusinessInfoResponseBodyData) GetBusinessName() *string {
	return s.BusinessName
}

func (s *QueryWabaBusinessInfoResponseBodyData) GetVerificationStatus() *string {
	return s.VerificationStatus
}

func (s *QueryWabaBusinessInfoResponseBodyData) GetVertical() *string {
	return s.Vertical
}

func (s *QueryWabaBusinessInfoResponseBodyData) SetBusinessId(v string) *QueryWabaBusinessInfoResponseBodyData {
	s.BusinessId = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBodyData) SetBusinessName(v string) *QueryWabaBusinessInfoResponseBodyData {
	s.BusinessName = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBodyData) SetVerificationStatus(v string) *QueryWabaBusinessInfoResponseBodyData {
	s.VerificationStatus = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBodyData) SetVertical(v string) *QueryWabaBusinessInfoResponseBodyData {
	s.Vertical = &v
	return s
}

func (s *QueryWabaBusinessInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
