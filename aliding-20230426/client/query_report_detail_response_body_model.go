// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReportDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*QueryReportDetailResponseBodyContent) *QueryReportDetailResponseBody
	GetContent() []*QueryReportDetailResponseBodyContent
	SetCreateTime(v int64) *QueryReportDetailResponseBody
	GetCreateTime() *int64
	SetCreatorId(v string) *QueryReportDetailResponseBody
	GetCreatorId() *string
	SetCreatorName(v string) *QueryReportDetailResponseBody
	GetCreatorName() *string
	SetDeptName(v string) *QueryReportDetailResponseBody
	GetDeptName() *string
	SetModifiedTime(v int64) *QueryReportDetailResponseBody
	GetModifiedTime() *int64
	SetRemark(v string) *QueryReportDetailResponseBody
	GetRemark() *string
	SetReportId(v string) *QueryReportDetailResponseBody
	GetReportId() *string
	SetRequestId(v string) *QueryReportDetailResponseBody
	GetRequestId() *string
	SetTemplateName(v string) *QueryReportDetailResponseBody
	GetTemplateName() *string
	SetVendorRequestId(v string) *QueryReportDetailResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *QueryReportDetailResponseBody
	GetVendorType() *string
}

type QueryReportDetailResponseBody struct {
	Content []*QueryReportDetailResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// example:
	//
	// 1691980997000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 012345
	CreatorId   *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	DeptName    *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// example:
	//
	// 1691980997000
	ModifiedTime *int64  `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Remark       *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 1231232134
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s QueryReportDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryReportDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReportDetailResponseBody) GetContent() []*QueryReportDetailResponseBodyContent {
	return s.Content
}

func (s *QueryReportDetailResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryReportDetailResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *QueryReportDetailResponseBody) GetCreatorName() *string {
	return s.CreatorName
}

func (s *QueryReportDetailResponseBody) GetDeptName() *string {
	return s.DeptName
}

func (s *QueryReportDetailResponseBody) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *QueryReportDetailResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *QueryReportDetailResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *QueryReportDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryReportDetailResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *QueryReportDetailResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *QueryReportDetailResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *QueryReportDetailResponseBody) SetContent(v []*QueryReportDetailResponseBodyContent) *QueryReportDetailResponseBody {
	s.Content = v
	return s
}

func (s *QueryReportDetailResponseBody) SetCreateTime(v int64) *QueryReportDetailResponseBody {
	s.CreateTime = &v
	return s
}

func (s *QueryReportDetailResponseBody) SetCreatorId(v string) *QueryReportDetailResponseBody {
	s.CreatorId = &v
	return s
}

func (s *QueryReportDetailResponseBody) SetCreatorName(v string) *QueryReportDetailResponseBody {
	s.CreatorName = &v
	return s
}

func (s *QueryReportDetailResponseBody) SetDeptName(v string) *QueryReportDetailResponseBody {
	s.DeptName = &v
	return s
}

func (s *QueryReportDetailResponseBody) SetModifiedTime(v int64) *QueryReportDetailResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *QueryReportDetailResponseBody) SetRemark(v string) *QueryReportDetailResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryReportDetailResponseBody) SetReportId(v string) *QueryReportDetailResponseBody {
	s.ReportId = &v
	return s
}

func (s *QueryReportDetailResponseBody) SetRequestId(v string) *QueryReportDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReportDetailResponseBody) SetTemplateName(v string) *QueryReportDetailResponseBody {
	s.TemplateName = &v
	return s
}

func (s *QueryReportDetailResponseBody) SetVendorRequestId(v string) *QueryReportDetailResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *QueryReportDetailResponseBody) SetVendorType(v string) *QueryReportDetailResponseBody {
	s.VendorType = &v
	return s
}

func (s *QueryReportDetailResponseBody) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryReportDetailResponseBodyContent struct {
	Images []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 0
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 1
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryReportDetailResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s QueryReportDetailResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryReportDetailResponseBodyContent) GetImages() []*string {
	return s.Images
}

func (s *QueryReportDetailResponseBodyContent) GetKey() *string {
	return s.Key
}

func (s *QueryReportDetailResponseBodyContent) GetSort() *string {
	return s.Sort
}

func (s *QueryReportDetailResponseBodyContent) GetType() *string {
	return s.Type
}

func (s *QueryReportDetailResponseBodyContent) GetValue() *string {
	return s.Value
}

func (s *QueryReportDetailResponseBodyContent) SetImages(v []*string) *QueryReportDetailResponseBodyContent {
	s.Images = v
	return s
}

func (s *QueryReportDetailResponseBodyContent) SetKey(v string) *QueryReportDetailResponseBodyContent {
	s.Key = &v
	return s
}

func (s *QueryReportDetailResponseBodyContent) SetSort(v string) *QueryReportDetailResponseBodyContent {
	s.Sort = &v
	return s
}

func (s *QueryReportDetailResponseBodyContent) SetType(v string) *QueryReportDetailResponseBodyContent {
	s.Type = &v
	return s
}

func (s *QueryReportDetailResponseBodyContent) SetValue(v string) *QueryReportDetailResponseBodyContent {
	s.Value = &v
	return s
}

func (s *QueryReportDetailResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
