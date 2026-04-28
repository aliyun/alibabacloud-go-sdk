// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeAdvicesResponseBodyData) *DescribeAdvicesResponseBody
	GetData() *DescribeAdvicesResponseBodyData
	SetRequestId(v string) *DescribeAdvicesResponseBody
	GetRequestId() *string
}

type DescribeAdvicesResponseBody struct {
	Data *DescribeAdvicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAdvicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesResponseBody) GetData() *DescribeAdvicesResponseBodyData {
	return s.Data
}

func (s *DescribeAdvicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdvicesResponseBody) SetData(v *DescribeAdvicesResponseBodyData) *DescribeAdvicesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdvicesResponseBody) SetRequestId(v string) *DescribeAdvicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvicesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAdvicesResponseBodyData struct {
	Advice []*DescribeAdvicesResponseBodyDataAdvice `json:"Advice,omitempty" xml:"Advice,omitempty" type:"Repeated"`
}

func (s DescribeAdvicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesResponseBodyData) GetAdvice() []*DescribeAdvicesResponseBodyDataAdvice {
	return s.Advice
}

func (s *DescribeAdvicesResponseBodyData) SetAdvice(v []*DescribeAdvicesResponseBodyDataAdvice) *DescribeAdvicesResponseBodyData {
	s.Advice = v
	return s
}

func (s *DescribeAdvicesResponseBodyData) Validate() error {
	if s.Advice != nil {
		for _, item := range s.Advice {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAdvicesResponseBodyDataAdvice struct {
	AliyunId    *int64  `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName   *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreated  *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsExpired   *bool   `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Resource    *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceId  *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Severity    *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s DescribeAdvicesResponseBodyDataAdvice) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesResponseBodyDataAdvice) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetAliyunId() *int64 {
	return s.AliyunId
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetCheckName() *string {
	return s.CheckName
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetContent() *string {
	return s.Content
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetDescription() *string {
	return s.Description
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetId() *int64 {
	return s.Id
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetIsExpired() *bool {
	return s.IsExpired
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetResource() *string {
	return s.Resource
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeAdvicesResponseBodyDataAdvice) GetSeverity() *int32 {
	return s.Severity
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetAliyunId(v int64) *DescribeAdvicesResponseBodyDataAdvice {
	s.AliyunId = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetCheckId(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetCheckName(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.CheckName = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetCheckPlanId(v int64) *DescribeAdvicesResponseBodyDataAdvice {
	s.CheckPlanId = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetContent(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Content = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetDescription(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Description = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetGmtCreated(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.GmtCreated = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetGmtModified(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.GmtModified = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetId(v int64) *DescribeAdvicesResponseBodyDataAdvice {
	s.Id = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetIsExpired(v bool) *DescribeAdvicesResponseBodyDataAdvice {
	s.IsExpired = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetProduct(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetResource(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.Resource = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetResourceId(v string) *DescribeAdvicesResponseBodyDataAdvice {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) SetSeverity(v int32) *DescribeAdvicesResponseBodyDataAdvice {
	s.Severity = &v
	return s
}

func (s *DescribeAdvicesResponseBodyDataAdvice) Validate() error {
	return dara.Validate(s)
}
