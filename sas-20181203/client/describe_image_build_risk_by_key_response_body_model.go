// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBuildRiskByKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeImageBuildRiskByKeyResponseBody
	GetCode() *string
	SetData(v *DescribeImageBuildRiskByKeyResponseBodyData) *DescribeImageBuildRiskByKeyResponseBody
	GetData() *DescribeImageBuildRiskByKeyResponseBodyData
	SetMessage(v string) *DescribeImageBuildRiskByKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeImageBuildRiskByKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeImageBuildRiskByKeyResponseBody
	GetSuccess() *bool
}

type DescribeImageBuildRiskByKeyResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeImageBuildRiskByKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 52870893-48A7-5A9E-9E05-6253E5B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeImageBuildRiskByKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBuildRiskByKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageBuildRiskByKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeImageBuildRiskByKeyResponseBody) GetData() *DescribeImageBuildRiskByKeyResponseBodyData {
	return s.Data
}

func (s *DescribeImageBuildRiskByKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeImageBuildRiskByKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageBuildRiskByKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeImageBuildRiskByKeyResponseBody) SetCode(v string) *DescribeImageBuildRiskByKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBody) SetData(v *DescribeImageBuildRiskByKeyResponseBodyData) *DescribeImageBuildRiskByKeyResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBody) SetMessage(v string) *DescribeImageBuildRiskByKeyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBody) SetRequestId(v string) *DescribeImageBuildRiskByKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBody) SetSuccess(v bool) *DescribeImageBuildRiskByKeyResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageBuildRiskByKeyResponseBodyData struct {
	// The risks.
	List []*DescribeImageBuildRiskByKeyResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s DescribeImageBuildRiskByKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBuildRiskByKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageBuildRiskByKeyResponseBodyData) GetList() []*DescribeImageBuildRiskByKeyResponseBodyDataList {
	return s.List
}

func (s *DescribeImageBuildRiskByKeyResponseBodyData) GetPageInfo() *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *DescribeImageBuildRiskByKeyResponseBodyData) SetList(v []*DescribeImageBuildRiskByKeyResponseBodyDataList) *DescribeImageBuildRiskByKeyResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyData) SetPageInfo(v *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo) *DescribeImageBuildRiskByKeyResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageBuildRiskByKeyResponseBodyDataList struct {
	// The suggestion on how to handle the risk.
	//
	// example:
	//
	// do not use root user
	Advice *string `json:"Advice,omitempty" xml:"Advice,omitempty"`
	// The description of the suggestion on how to handle the risk.
	//
	// example:
	//
	// the root user has excessive permissions
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image build command.
	//
	// example:
	//
	// user root
	LayerCmd *string `json:"LayerCmd,omitempty" xml:"LayerCmd,omitempty"`
	// The digest of the image.
	//
	// example:
	//
	// 6ec898e6274f942e0e4a053eff1c3119026a6704e56cff206b2cec71f636****
	LayerDigest *string `json:"LayerDigest,omitempty" xml:"LayerDigest,omitempty"`
	// The prompt message on the risk.
	//
	// example:
	//
	// the root user has excessive permissions
	Promt *string `json:"Promt,omitempty" xml:"Promt,omitempty"`
	// The type key of the risk rule.
	//
	// example:
	//
	// other
	RiskClass *string `json:"RiskClass,omitempty" xml:"RiskClass,omitempty"`
	// The type name of the risk rule.
	//
	// example:
	//
	// other
	RiskClassName *string `json:"RiskClassName,omitempty" xml:"RiskClassName,omitempty"`
	// The key of the risk rule.
	//
	// example:
	//
	// no_user
	RiskKey *string `json:"RiskKey,omitempty" xml:"RiskKey,omitempty"`
	// The name of the risk rule.
	//
	// example:
	//
	// no_user
	RiskKeyName *string `json:"RiskKeyName,omitempty" xml:"RiskKeyName,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// medium
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeImageBuildRiskByKeyResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBuildRiskByKeyResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) GetAdvice() *string {
	return s.Advice
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) GetLayerCmd() *string {
	return s.LayerCmd
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) GetLayerDigest() *string {
	return s.LayerDigest
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) GetPromt() *string {
	return s.Promt
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) GetRiskClass() *string {
	return s.RiskClass
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) GetRiskClassName() *string {
	return s.RiskClassName
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) GetRiskKey() *string {
	return s.RiskKey
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) GetRiskKeyName() *string {
	return s.RiskKeyName
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) SetAdvice(v string) *DescribeImageBuildRiskByKeyResponseBodyDataList {
	s.Advice = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) SetDescription(v string) *DescribeImageBuildRiskByKeyResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) SetLayerCmd(v string) *DescribeImageBuildRiskByKeyResponseBodyDataList {
	s.LayerCmd = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) SetLayerDigest(v string) *DescribeImageBuildRiskByKeyResponseBodyDataList {
	s.LayerDigest = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) SetPromt(v string) *DescribeImageBuildRiskByKeyResponseBodyDataList {
	s.Promt = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) SetRiskClass(v string) *DescribeImageBuildRiskByKeyResponseBodyDataList {
	s.RiskClass = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) SetRiskClassName(v string) *DescribeImageBuildRiskByKeyResponseBodyDataList {
	s.RiskClassName = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) SetRiskKey(v string) *DescribeImageBuildRiskByKeyResponseBodyDataList {
	s.RiskKey = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) SetRiskKeyName(v string) *DescribeImageBuildRiskByKeyResponseBodyDataList {
	s.RiskKeyName = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) SetRiskLevel(v string) *DescribeImageBuildRiskByKeyResponseBodyDataList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeImageBuildRiskByKeyResponseBodyDataPageInfo struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 109
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageBuildRiskByKeyResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBuildRiskByKeyResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo) SetTotalCount(v int32) *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageBuildRiskByKeyResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}
