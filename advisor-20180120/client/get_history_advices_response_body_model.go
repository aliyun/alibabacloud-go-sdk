// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoryAdvicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetHistoryAdvicesResponseBodyData) *GetHistoryAdvicesResponseBody
	GetData() *GetHistoryAdvicesResponseBodyData
	SetRequestId(v string) *GetHistoryAdvicesResponseBody
	GetRequestId() *string
}

type GetHistoryAdvicesResponseBody struct {
	Data *GetHistoryAdvicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHistoryAdvicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryAdvicesResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoryAdvicesResponseBody) GetData() *GetHistoryAdvicesResponseBodyData {
	return s.Data
}

func (s *GetHistoryAdvicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHistoryAdvicesResponseBody) SetData(v *GetHistoryAdvicesResponseBodyData) *GetHistoryAdvicesResponseBody {
	s.Data = v
	return s
}

func (s *GetHistoryAdvicesResponseBody) SetRequestId(v string) *GetHistoryAdvicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHistoryAdvicesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHistoryAdvicesResponseBodyData struct {
	// example:
	//
	// 1
	PageNo *int32                                     `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	Result []*GetHistoryAdvicesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetHistoryAdvicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryAdvicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistoryAdvicesResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetHistoryAdvicesResponseBodyData) GetResult() []*GetHistoryAdvicesResponseBodyDataResult {
	return s.Result
}

func (s *GetHistoryAdvicesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *GetHistoryAdvicesResponseBodyData) SetPageNo(v int32) *GetHistoryAdvicesResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyData) SetResult(v []*GetHistoryAdvicesResponseBodyDataResult) *GetHistoryAdvicesResponseBodyData {
	s.Result = v
	return s
}

func (s *GetHistoryAdvicesResponseBodyData) SetTotal(v int32) *GetHistoryAdvicesResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyData) Validate() error {
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

type GetHistoryAdvicesResponseBodyDataResult struct {
	// example:
	//
	// EcsHighCpuUtilization
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName   *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 1
	Severity *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetHistoryAdvicesResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryAdvicesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetHistoryAdvicesResponseBodyDataResult) GetCheckId() *string {
	return s.CheckId
}

func (s *GetHistoryAdvicesResponseBodyDataResult) GetCheckName() *string {
	return s.CheckName
}

func (s *GetHistoryAdvicesResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *GetHistoryAdvicesResponseBodyDataResult) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *GetHistoryAdvicesResponseBodyDataResult) GetProduct() *string {
	return s.Product
}

func (s *GetHistoryAdvicesResponseBodyDataResult) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetHistoryAdvicesResponseBodyDataResult) GetSeverity() *int32 {
	return s.Severity
}

func (s *GetHistoryAdvicesResponseBodyDataResult) GetUrl() *string {
	return s.Url
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetCheckId(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.CheckId = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetCheckName(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.CheckName = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetDescription(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetGmtCreated(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.GmtCreated = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetProduct(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.Product = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetResourceId(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.ResourceId = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetSeverity(v int32) *GetHistoryAdvicesResponseBodyDataResult {
	s.Severity = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) SetUrl(v string) *GetHistoryAdvicesResponseBodyDataResult {
	s.Url = &v
	return s
}

func (s *GetHistoryAdvicesResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
