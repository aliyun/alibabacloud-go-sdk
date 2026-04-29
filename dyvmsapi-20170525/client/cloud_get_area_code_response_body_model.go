// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetAreaCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudGetAreaCodeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudGetAreaCodeResponseBody
	GetCode() *string
	SetData(v *CloudGetAreaCodeResponseBodyData) *CloudGetAreaCodeResponseBody
	GetData() *CloudGetAreaCodeResponseBodyData
	SetMessage(v string) *CloudGetAreaCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudGetAreaCodeResponseBody
	GetRequestId() *string
}

type CloudGetAreaCodeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudGetAreaCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 44D40F36-32D9-555E-89A6-26D39C3044D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudGetAreaCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAreaCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CloudGetAreaCodeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudGetAreaCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudGetAreaCodeResponseBody) GetData() *CloudGetAreaCodeResponseBodyData {
	return s.Data
}

func (s *CloudGetAreaCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudGetAreaCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudGetAreaCodeResponseBody) SetAccessDeniedDetail(v string) *CloudGetAreaCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudGetAreaCodeResponseBody) SetCode(v string) *CloudGetAreaCodeResponseBody {
	s.Code = &v
	return s
}

func (s *CloudGetAreaCodeResponseBody) SetData(v *CloudGetAreaCodeResponseBodyData) *CloudGetAreaCodeResponseBody {
	s.Data = v
	return s
}

func (s *CloudGetAreaCodeResponseBody) SetMessage(v string) *CloudGetAreaCodeResponseBody {
	s.Message = &v
	return s
}

func (s *CloudGetAreaCodeResponseBody) SetRequestId(v string) *CloudGetAreaCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudGetAreaCodeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudGetAreaCodeResponseBodyData struct {
	// 区号
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 号码所属城市
	//
	// example:
	//
	// 示例值示例值
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 号码所属省份
	//
	// example:
	//
	// 示例值示例值示例值
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 手机号
	//
	// example:
	//
	// 183xxxx7093
	Tel *string `json:"Tel,omitempty" xml:"Tel,omitempty"`
}

func (s CloudGetAreaCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAreaCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudGetAreaCodeResponseBodyData) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudGetAreaCodeResponseBodyData) GetCity() *string {
	return s.City
}

func (s *CloudGetAreaCodeResponseBodyData) GetProvince() *string {
	return s.Province
}

func (s *CloudGetAreaCodeResponseBodyData) GetTel() *string {
	return s.Tel
}

func (s *CloudGetAreaCodeResponseBodyData) SetAreaCode(v string) *CloudGetAreaCodeResponseBodyData {
	s.AreaCode = &v
	return s
}

func (s *CloudGetAreaCodeResponseBodyData) SetCity(v string) *CloudGetAreaCodeResponseBodyData {
	s.City = &v
	return s
}

func (s *CloudGetAreaCodeResponseBodyData) SetProvince(v string) *CloudGetAreaCodeResponseBodyData {
	s.Province = &v
	return s
}

func (s *CloudGetAreaCodeResponseBodyData) SetTel(v string) *CloudGetAreaCodeResponseBodyData {
	s.Tel = &v
	return s
}

func (s *CloudGetAreaCodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
