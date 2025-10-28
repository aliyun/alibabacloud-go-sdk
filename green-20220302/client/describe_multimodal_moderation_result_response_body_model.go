// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultimodalModerationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeMultimodalModerationResultResponseBody
	GetCode() *int64
	SetData(v *DescribeMultimodalModerationResultResponseBodyData) *DescribeMultimodalModerationResultResponseBody
	GetData() *DescribeMultimodalModerationResultResponseBodyData
	SetMsg(v string) *DescribeMultimodalModerationResultResponseBody
	GetMsg() *string
	SetRequestId(v string) *DescribeMultimodalModerationResultResponseBody
	GetRequestId() *string
}

type DescribeMultimodalModerationResultResponseBody struct {
	Code      *int64                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeMultimodalModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                                             `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMultimodalModerationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultimodalModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultimodalModerationResultResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeMultimodalModerationResultResponseBody) GetData() *DescribeMultimodalModerationResultResponseBodyData {
	return s.Data
}

func (s *DescribeMultimodalModerationResultResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DescribeMultimodalModerationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMultimodalModerationResultResponseBody) SetCode(v int64) *DescribeMultimodalModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBody) SetData(v *DescribeMultimodalModerationResultResponseBodyData) *DescribeMultimodalModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBody) SetMsg(v string) *DescribeMultimodalModerationResultResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBody) SetRequestId(v string) *DescribeMultimodalModerationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultimodalModerationResultResponseBodyData struct {
	CommentDatas []*DescribeMultimodalModerationResultResponseBodyDataCommentDatas `json:"CommentDatas,omitempty" xml:"CommentDatas,omitempty" type:"Repeated"`
	DataId       *string                                                           `json:"DataId,omitempty" xml:"DataId,omitempty"`
	MainData     *DescribeMultimodalModerationResultResponseBodyDataMainData       `json:"MainData,omitempty" xml:"MainData,omitempty" type:"Struct"`
	ReqId        *string                                                           `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	RiskLevel    *string                                                           `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeMultimodalModerationResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultimodalModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMultimodalModerationResultResponseBodyData) GetCommentDatas() []*DescribeMultimodalModerationResultResponseBodyDataCommentDatas {
	return s.CommentDatas
}

func (s *DescribeMultimodalModerationResultResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *DescribeMultimodalModerationResultResponseBodyData) GetMainData() *DescribeMultimodalModerationResultResponseBodyDataMainData {
	return s.MainData
}

func (s *DescribeMultimodalModerationResultResponseBodyData) GetReqId() *string {
	return s.ReqId
}

func (s *DescribeMultimodalModerationResultResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeMultimodalModerationResultResponseBodyData) SetCommentDatas(v []*DescribeMultimodalModerationResultResponseBodyDataCommentDatas) *DescribeMultimodalModerationResultResponseBodyData {
	s.CommentDatas = v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyData) SetDataId(v string) *DescribeMultimodalModerationResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyData) SetMainData(v *DescribeMultimodalModerationResultResponseBodyDataMainData) *DescribeMultimodalModerationResultResponseBodyData {
	s.MainData = v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyData) SetReqId(v string) *DescribeMultimodalModerationResultResponseBodyData {
	s.ReqId = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyData) SetRiskLevel(v string) *DescribeMultimodalModerationResultResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyData) Validate() error {
	if s.CommentDatas != nil {
		for _, item := range s.CommentDatas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MainData != nil {
		if err := s.MainData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultimodalModerationResultResponseBodyDataCommentDatas struct {
	CommentDatas []*DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatas `json:"CommentDatas,omitempty" xml:"CommentDatas,omitempty" type:"Repeated"`
	Results      []*DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults      `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DescribeMultimodalModerationResultResponseBodyDataCommentDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultimodalModerationResultResponseBodyDataCommentDatas) GoString() string {
	return s.String()
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatas) GetCommentDatas() []*DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatas {
	return s.CommentDatas
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatas) GetResults() []*DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults {
	return s.Results
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatas) SetCommentDatas(v []*DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatas) *DescribeMultimodalModerationResultResponseBodyDataCommentDatas {
	s.CommentDatas = v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatas) SetResults(v []*DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults) *DescribeMultimodalModerationResultResponseBodyDataCommentDatas {
	s.Results = v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatas) Validate() error {
	if s.CommentDatas != nil {
		for _, item := range s.CommentDatas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatas struct {
	Results []*DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatas) GoString() string {
	return s.String()
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatas) GetResults() []*DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults {
	return s.Results
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatas) SetResults(v []*DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults) *DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatas {
	s.Results = v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatas) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults) GoString() string {
	return s.String()
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults) GetDescription() *string {
	return s.Description
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults) GetLabel() *string {
	return s.Label
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults) SetDescription(v string) *DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults {
	s.Description = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults) SetLabel(v string) *DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults {
	s.Label = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasCommentDatasResults) Validate() error {
	return dara.Validate(s)
}

type DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults) GoString() string {
	return s.String()
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults) GetDescription() *string {
	return s.Description
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults) GetLabel() *string {
	return s.Label
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults) SetDescription(v string) *DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults {
	s.Description = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults) SetLabel(v string) *DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults {
	s.Label = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyDataCommentDatasResults) Validate() error {
	return dara.Validate(s)
}

type DescribeMultimodalModerationResultResponseBodyDataMainData struct {
	Results []*DescribeMultimodalModerationResultResponseBodyDataMainDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DescribeMultimodalModerationResultResponseBodyDataMainData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultimodalModerationResultResponseBodyDataMainData) GoString() string {
	return s.String()
}

func (s *DescribeMultimodalModerationResultResponseBodyDataMainData) GetResults() []*DescribeMultimodalModerationResultResponseBodyDataMainDataResults {
	return s.Results
}

func (s *DescribeMultimodalModerationResultResponseBodyDataMainData) SetResults(v []*DescribeMultimodalModerationResultResponseBodyDataMainDataResults) *DescribeMultimodalModerationResultResponseBodyDataMainData {
	s.Results = v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyDataMainData) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultimodalModerationResultResponseBodyDataMainDataResults struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DescribeMultimodalModerationResultResponseBodyDataMainDataResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultimodalModerationResultResponseBodyDataMainDataResults) GoString() string {
	return s.String()
}

func (s *DescribeMultimodalModerationResultResponseBodyDataMainDataResults) GetDescription() *string {
	return s.Description
}

func (s *DescribeMultimodalModerationResultResponseBodyDataMainDataResults) GetLabel() *string {
	return s.Label
}

func (s *DescribeMultimodalModerationResultResponseBodyDataMainDataResults) SetDescription(v string) *DescribeMultimodalModerationResultResponseBodyDataMainDataResults {
	s.Description = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyDataMainDataResults) SetLabel(v string) *DescribeMultimodalModerationResultResponseBodyDataMainDataResults {
	s.Label = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponseBodyDataMainDataResults) Validate() error {
	return dara.Validate(s)
}
