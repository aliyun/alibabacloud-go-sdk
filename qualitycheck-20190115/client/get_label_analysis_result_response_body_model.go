// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLabelAnalysisResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLabelAnalysisResultResponseBody
	GetCode() *string
	SetData(v *GetLabelAnalysisResultResponseBodyData) *GetLabelAnalysisResultResponseBody
	GetData() *GetLabelAnalysisResultResponseBodyData
	SetMessage(v string) *GetLabelAnalysisResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLabelAnalysisResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLabelAnalysisResultResponseBody
	GetSuccess() *bool
}

type GetLabelAnalysisResultResponseBody struct {
	// The response code. **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *GetLabelAnalysisResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLabelAnalysisResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLabelAnalysisResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetLabelAnalysisResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLabelAnalysisResultResponseBody) GetData() *GetLabelAnalysisResultResponseBodyData {
	return s.Data
}

func (s *GetLabelAnalysisResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLabelAnalysisResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLabelAnalysisResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLabelAnalysisResultResponseBody) SetCode(v string) *GetLabelAnalysisResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBody) SetData(v *GetLabelAnalysisResultResponseBodyData) *GetLabelAnalysisResultResponseBody {
	s.Data = v
	return s
}

func (s *GetLabelAnalysisResultResponseBody) SetMessage(v string) *GetLabelAnalysisResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBody) SetRequestId(v string) *GetLabelAnalysisResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBody) SetSuccess(v bool) *GetLabelAnalysisResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLabelAnalysisResultResponseBodyData struct {
	// The total number of input tokens accumulated during this task.
	//
	// example:
	//
	// 7371
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// The total number of output tokens accumulated during this task.
	//
	// example:
	//
	// 355
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// The tree-structured tag results.
	TagList []*GetLabelAnalysisResultResponseBodyDataTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The total number of tokens accumulated during this task.
	//
	// example:
	//
	// 7726
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
	// The total number of Qwen-Plus model calls accumulated during this task.
	//
	// example:
	//
	// 4
	TyxmPlusCount *int64 `json:"TyxmPlusCount,omitempty" xml:"TyxmPlusCount,omitempty"`
	// The total number of Qwen-Turbo model calls accumulated during this task.
	//
	// example:
	//
	// 0
	TyxmTurboCount *int64 `json:"TyxmTurboCount,omitempty" xml:"TyxmTurboCount,omitempty"`
}

func (s GetLabelAnalysisResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLabelAnalysisResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLabelAnalysisResultResponseBodyData) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetLabelAnalysisResultResponseBodyData) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetLabelAnalysisResultResponseBodyData) GetTagList() []*GetLabelAnalysisResultResponseBodyDataTagList {
	return s.TagList
}

func (s *GetLabelAnalysisResultResponseBodyData) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetLabelAnalysisResultResponseBodyData) GetTyxmPlusCount() *int64 {
	return s.TyxmPlusCount
}

func (s *GetLabelAnalysisResultResponseBodyData) GetTyxmTurboCount() *int64 {
	return s.TyxmTurboCount
}

func (s *GetLabelAnalysisResultResponseBodyData) SetInputTokens(v int64) *GetLabelAnalysisResultResponseBodyData {
	s.InputTokens = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyData) SetOutputTokens(v int64) *GetLabelAnalysisResultResponseBodyData {
	s.OutputTokens = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyData) SetTagList(v []*GetLabelAnalysisResultResponseBodyDataTagList) *GetLabelAnalysisResultResponseBodyData {
	s.TagList = v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyData) SetTotalTokens(v int64) *GetLabelAnalysisResultResponseBodyData {
	s.TotalTokens = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyData) SetTyxmPlusCount(v int64) *GetLabelAnalysisResultResponseBodyData {
	s.TyxmPlusCount = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyData) SetTyxmTurboCount(v int64) *GetLabelAnalysisResultResponseBodyData {
	s.TyxmTurboCount = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyData) Validate() error {
	if s.TagList != nil {
		for _, item := range s.TagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLabelAnalysisResultResponseBodyDataTagList struct {
	// The list of child nodes.
	Children []*GetLabelAnalysisResultResponseBodyDataTagListChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// The description of the tag analysis.
	//
	// example:
	//
	// 用户在千问内通过高德打车支付17元失败，转支付宝后变原价；客服围绕支付失败与余额、实名认证、授权绑定进行排查。问题发生在千问调用高德打车场景，属千问×高德。
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The tag name.
	//
	// example:
	//
	// 千问×高德
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetLabelAnalysisResultResponseBodyDataTagList) String() string {
	return dara.Prettify(s)
}

func (s GetLabelAnalysisResultResponseBodyDataTagList) GoString() string {
	return s.String()
}

func (s *GetLabelAnalysisResultResponseBodyDataTagList) GetChildren() []*GetLabelAnalysisResultResponseBodyDataTagListChildren {
	return s.Children
}

func (s *GetLabelAnalysisResultResponseBodyDataTagList) GetRemarks() *string {
	return s.Remarks
}

func (s *GetLabelAnalysisResultResponseBodyDataTagList) GetTagName() *string {
	return s.TagName
}

func (s *GetLabelAnalysisResultResponseBodyDataTagList) SetChildren(v []*GetLabelAnalysisResultResponseBodyDataTagListChildren) *GetLabelAnalysisResultResponseBodyDataTagList {
	s.Children = v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagList) SetRemarks(v string) *GetLabelAnalysisResultResponseBodyDataTagList {
	s.Remarks = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagList) SetTagName(v string) *GetLabelAnalysisResultResponseBodyDataTagList {
	s.TagName = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagList) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLabelAnalysisResultResponseBodyDataTagListChildren struct {
	// The list of child nodes.
	Children []*GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// The description of the tag analysis.
	//
	// example:
	//
	// 用户在千问内通过高德打车支付17元失败，转支付宝后变原价；客服围绕支付失败与余额、实名认证、授权绑定进行排查。问题发生在千问调用高德打车场景，属千问×高德。
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The tag name.
	//
	// example:
	//
	// 千问×高德
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetLabelAnalysisResultResponseBodyDataTagListChildren) String() string {
	return dara.Prettify(s)
}

func (s GetLabelAnalysisResultResponseBodyDataTagListChildren) GoString() string {
	return s.String()
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildren) GetChildren() []*GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren {
	return s.Children
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildren) GetRemarks() *string {
	return s.Remarks
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildren) GetTagName() *string {
	return s.TagName
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildren) SetChildren(v []*GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren) *GetLabelAnalysisResultResponseBodyDataTagListChildren {
	s.Children = v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildren) SetRemarks(v string) *GetLabelAnalysisResultResponseBodyDataTagListChildren {
	s.Remarks = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildren) SetTagName(v string) *GetLabelAnalysisResultResponseBodyDataTagListChildren {
	s.TagName = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildren) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren struct {
	// The list of child nodes.
	Children []*GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// The description of the tag analysis.
	//
	// example:
	//
	// 用户在千问内通过高德打车支付17元失败，转支付宝后变原价；客服围绕支付失败与余额、实名认证、授权绑定进行排查。问题发生在千问调用高德打车场景，属千问×高德。
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The tag name.
	//
	// example:
	//
	// 千问×高德
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren) String() string {
	return dara.Prettify(s)
}

func (s GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren) GoString() string {
	return s.String()
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren) GetChildren() []*GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren {
	return s.Children
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren) GetRemarks() *string {
	return s.Remarks
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren) GetTagName() *string {
	return s.TagName
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren) SetChildren(v []*GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren) *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren {
	s.Children = v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren) SetRemarks(v string) *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren {
	s.Remarks = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren) SetTagName(v string) *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren {
	s.TagName = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildren) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren struct {
	// The list of child nodes.
	Children []*GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// The description of the tag analysis.
	//
	// example:
	//
	// 用户在千问内通过高德打车支付17元失败，转支付宝后变原价；客服围绕支付失败与余额、实名认证、授权绑定进行排查。问题发生在千问调用高德打车场景，属千问×高德。
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The tag name.
	//
	// example:
	//
	// 千问×高德
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren) String() string {
	return dara.Prettify(s)
}

func (s GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren) GoString() string {
	return s.String()
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren) GetChildren() []*GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren {
	return s.Children
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren) GetRemarks() *string {
	return s.Remarks
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren) GetTagName() *string {
	return s.TagName
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren) SetChildren(v []*GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren) *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren {
	s.Children = v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren) SetRemarks(v string) *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren {
	s.Remarks = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren) SetTagName(v string) *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren {
	s.TagName = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildren) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren struct {
	// The description of the tag analysis.
	//
	// example:
	//
	// 当前层级未命中有效标签
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The tag name.
	//
	// example:
	//
	// 无效会话
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren) String() string {
	return dara.Prettify(s)
}

func (s GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren) GoString() string {
	return s.String()
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren) GetRemarks() *string {
	return s.Remarks
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren) GetTagName() *string {
	return s.TagName
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren) SetRemarks(v string) *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren {
	s.Remarks = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren) SetTagName(v string) *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren {
	s.TagName = &v
	return s
}

func (s *GetLabelAnalysisResultResponseBodyDataTagListChildrenChildrenChildrenChildren) Validate() error {
	return dara.Validate(s)
}
