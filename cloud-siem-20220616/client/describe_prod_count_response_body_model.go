// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProdCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeProdCountResponseBodyData) *DescribeProdCountResponseBody
	GetData() *DescribeProdCountResponseBodyData
	SetRequestId(v string) *DescribeProdCountResponseBody
	GetRequestId() *string
}

type DescribeProdCountResponseBody struct {
	// The data returned.
	Data *DescribeProdCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProdCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProdCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProdCountResponseBody) GetData() *DescribeProdCountResponseBodyData {
	return s.Data
}

func (s *DescribeProdCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProdCountResponseBody) SetData(v *DescribeProdCountResponseBodyData) *DescribeProdCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeProdCountResponseBody) SetRequestId(v string) *DescribeProdCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProdCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProdCountResponseBodyData struct {
	AliyunImportedCount *int32 `json:"AliyunImportedCount,omitempty" xml:"AliyunImportedCount,omitempty"`
	// The number of Alibaba Cloud services.
	//
	// example:
	//
	// 19
	AliyunProdCount     *int32 `json:"AliyunProdCount,omitempty" xml:"AliyunProdCount,omitempty"`
	HcloudImportedCount *int32 `json:"HcloudImportedCount,omitempty" xml:"HcloudImportedCount,omitempty"`
	// The number of Huawei Cloud services.
	//
	// example:
	//
	// 2
	HcloudProdCount  *int32 `json:"HcloudProdCount,omitempty" xml:"HcloudProdCount,omitempty"`
	IdcImportedCount *int32 `json:"IdcImportedCount,omitempty" xml:"IdcImportedCount,omitempty"`
	// example:
	//
	// 2
	IdcProdCount        *int32 `json:"IdcProdCount,omitempty" xml:"IdcProdCount,omitempty"`
	QcloudImportedCount *int32 `json:"QcloudImportedCount,omitempty" xml:"QcloudImportedCount,omitempty"`
	// The number of Tencent Cloud services.
	//
	// example:
	//
	// 2
	QcloudProdCount *int32 `json:"QcloudProdCount,omitempty" xml:"QcloudProdCount,omitempty"`
}

func (s DescribeProdCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeProdCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeProdCountResponseBodyData) GetAliyunImportedCount() *int32 {
	return s.AliyunImportedCount
}

func (s *DescribeProdCountResponseBodyData) GetAliyunProdCount() *int32 {
	return s.AliyunProdCount
}

func (s *DescribeProdCountResponseBodyData) GetHcloudImportedCount() *int32 {
	return s.HcloudImportedCount
}

func (s *DescribeProdCountResponseBodyData) GetHcloudProdCount() *int32 {
	return s.HcloudProdCount
}

func (s *DescribeProdCountResponseBodyData) GetIdcImportedCount() *int32 {
	return s.IdcImportedCount
}

func (s *DescribeProdCountResponseBodyData) GetIdcProdCount() *int32 {
	return s.IdcProdCount
}

func (s *DescribeProdCountResponseBodyData) GetQcloudImportedCount() *int32 {
	return s.QcloudImportedCount
}

func (s *DescribeProdCountResponseBodyData) GetQcloudProdCount() *int32 {
	return s.QcloudProdCount
}

func (s *DescribeProdCountResponseBodyData) SetAliyunImportedCount(v int32) *DescribeProdCountResponseBodyData {
	s.AliyunImportedCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetAliyunProdCount(v int32) *DescribeProdCountResponseBodyData {
	s.AliyunProdCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetHcloudImportedCount(v int32) *DescribeProdCountResponseBodyData {
	s.HcloudImportedCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetHcloudProdCount(v int32) *DescribeProdCountResponseBodyData {
	s.HcloudProdCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetIdcImportedCount(v int32) *DescribeProdCountResponseBodyData {
	s.IdcImportedCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetIdcProdCount(v int32) *DescribeProdCountResponseBodyData {
	s.IdcProdCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetQcloudImportedCount(v int32) *DescribeProdCountResponseBodyData {
	s.QcloudImportedCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetQcloudProdCount(v int32) *DescribeProdCountResponseBodyData {
	s.QcloudProdCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
