// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *ListHotelInfoResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *ListHotelInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotelInfoResponseBody
	GetRequestId() *string
	SetResult(v []*ListHotelInfoResponseBodyResult) *ListHotelInfoResponseBody
	GetResult() []*ListHotelInfoResponseBodyResult
	SetStatusCode(v int32) *ListHotelInfoResponseBody
	GetStatusCode() *int32
}

type ListHotelInfoResponseBody struct {
	Extentions map[string]interface{}             `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*ListHotelInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListHotelInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotelInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelInfoResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *ListHotelInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotelInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotelInfoResponseBody) GetResult() []*ListHotelInfoResponseBodyResult {
	return s.Result
}

func (s *ListHotelInfoResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelInfoResponseBody) SetExtentions(v map[string]interface{}) *ListHotelInfoResponseBody {
	s.Extentions = v
	return s
}

func (s *ListHotelInfoResponseBody) SetMessage(v string) *ListHotelInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelInfoResponseBody) SetRequestId(v string) *ListHotelInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelInfoResponseBody) SetResult(v []*ListHotelInfoResponseBodyResult) *ListHotelInfoResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelInfoResponseBody) SetStatusCode(v int32) *ListHotelInfoResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ListHotelInfoResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotelInfoResponseBodyResult struct {
	AuthAccount  []*ListHotelInfoResponseBodyResultAuthAccount `json:"AuthAccount,omitempty" xml:"AuthAccount,omitempty" type:"Repeated"`
	HotelAddress *string                                       `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	// example:
	//
	// cf2446fc9d144c85aaee4f9ae20a96e7
	HotelId   *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelName *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
}

func (s ListHotelInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListHotelInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelInfoResponseBodyResult) GetAuthAccount() []*ListHotelInfoResponseBodyResultAuthAccount {
	return s.AuthAccount
}

func (s *ListHotelInfoResponseBodyResult) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *ListHotelInfoResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *ListHotelInfoResponseBodyResult) GetHotelName() *string {
	return s.HotelName
}

func (s *ListHotelInfoResponseBodyResult) SetAuthAccount(v []*ListHotelInfoResponseBodyResultAuthAccount) *ListHotelInfoResponseBodyResult {
	s.AuthAccount = v
	return s
}

func (s *ListHotelInfoResponseBodyResult) SetHotelAddress(v string) *ListHotelInfoResponseBodyResult {
	s.HotelAddress = &v
	return s
}

func (s *ListHotelInfoResponseBodyResult) SetHotelId(v string) *ListHotelInfoResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *ListHotelInfoResponseBodyResult) SetHotelName(v string) *ListHotelInfoResponseBodyResult {
	s.HotelName = &v
	return s
}

func (s *ListHotelInfoResponseBodyResult) Validate() error {
	if s.AuthAccount != nil {
		for _, item := range s.AuthAccount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotelInfoResponseBodyResultAuthAccount struct {
	// example:
	//
	// leetest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListHotelInfoResponseBodyResultAuthAccount) String() string {
	return dara.Prettify(s)
}

func (s ListHotelInfoResponseBodyResultAuthAccount) GoString() string {
	return s.String()
}

func (s *ListHotelInfoResponseBodyResultAuthAccount) GetUserName() *string {
	return s.UserName
}

func (s *ListHotelInfoResponseBodyResultAuthAccount) SetUserName(v string) *ListHotelInfoResponseBodyResultAuthAccount {
	s.UserName = &v
	return s
}

func (s *ListHotelInfoResponseBodyResultAuthAccount) Validate() error {
	return dara.Validate(s)
}
