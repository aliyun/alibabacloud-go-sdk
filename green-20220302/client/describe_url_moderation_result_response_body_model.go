// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUrlModerationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeUrlModerationResultResponseBody
	GetCode() *int32
	SetData(v *DescribeUrlModerationResultResponseBodyData) *DescribeUrlModerationResultResponseBody
	GetData() *DescribeUrlModerationResultResponseBodyData
	SetMsg(v string) *DescribeUrlModerationResultResponseBody
	GetMsg() *string
	SetRequestId(v string) *DescribeUrlModerationResultResponseBody
	GetRequestId() *string
}

type DescribeUrlModerationResultResponseBody struct {
	// The returned HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeUrlModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01F9144A-2088-5D87-935B-2DB865284B1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUrlModerationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUrlModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUrlModerationResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeUrlModerationResultResponseBody) GetData() *DescribeUrlModerationResultResponseBodyData {
	return s.Data
}

func (s *DescribeUrlModerationResultResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DescribeUrlModerationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUrlModerationResultResponseBody) SetCode(v int32) *DescribeUrlModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBody) SetData(v *DescribeUrlModerationResultResponseBodyData) *DescribeUrlModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUrlModerationResultResponseBody) SetMsg(v string) *DescribeUrlModerationResultResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBody) SetRequestId(v string) *DescribeUrlModerationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUrlModerationResultResponseBodyData struct {
	// The value of dataId that is specified in the API request. If this parameter is not specified in the API request, this field is not available in the response.
	//
	// example:
	//
	// 26769ada6e264e7ba9aa048241e12be9
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The supplementary information.
	ExtraInfo *DescribeUrlModerationResultResponseBodyDataExtraInfo `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty" type:"Struct"`
	// The ReqId field returned by an asynchronous URL moderation operation.
	//
	// example:
	//
	// B0963D30-BAB4-562F-9ED0-7A23AEC51C7C
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	// The returned results.
	Result []*DescribeUrlModerationResultResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeUrlModerationResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUrlModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUrlModerationResultResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *DescribeUrlModerationResultResponseBodyData) GetExtraInfo() *DescribeUrlModerationResultResponseBodyDataExtraInfo {
	return s.ExtraInfo
}

func (s *DescribeUrlModerationResultResponseBodyData) GetReqId() *string {
	return s.ReqId
}

func (s *DescribeUrlModerationResultResponseBodyData) GetResult() []*DescribeUrlModerationResultResponseBodyDataResult {
	return s.Result
}

func (s *DescribeUrlModerationResultResponseBodyData) SetDataId(v string) *DescribeUrlModerationResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyData) SetExtraInfo(v *DescribeUrlModerationResultResponseBodyDataExtraInfo) *DescribeUrlModerationResultResponseBodyData {
	s.ExtraInfo = v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyData) SetReqId(v string) *DescribeUrlModerationResultResponseBodyData {
	s.ReqId = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyData) SetResult(v []*DescribeUrlModerationResultResponseBodyDataResult) *DescribeUrlModerationResultResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyData) Validate() error {
	if s.ExtraInfo != nil {
		if err := s.ExtraInfo.Validate(); err != nil {
			return err
		}
	}
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

type DescribeUrlModerationResultResponseBodyDataExtraInfo struct {
	// The ICP number.
	//
	// example:
	//
	// xx
	IcpNo *string `json:"IcpNo,omitempty" xml:"IcpNo,omitempty"`
	// The type of the ICP filing.
	//
	// example:
	//
	// xx
	IcpType *string `json:"IcpType,omitempty" xml:"IcpType,omitempty"`
	// The type of site
	//
	// example:
	//
	// game
	SiteType *string `json:"SiteType,omitempty" xml:"SiteType,omitempty"`
}

func (s DescribeUrlModerationResultResponseBodyDataExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeUrlModerationResultResponseBodyDataExtraInfo) GoString() string {
	return s.String()
}

func (s *DescribeUrlModerationResultResponseBodyDataExtraInfo) GetIcpNo() *string {
	return s.IcpNo
}

func (s *DescribeUrlModerationResultResponseBodyDataExtraInfo) GetIcpType() *string {
	return s.IcpType
}

func (s *DescribeUrlModerationResultResponseBodyDataExtraInfo) GetSiteType() *string {
	return s.SiteType
}

func (s *DescribeUrlModerationResultResponseBodyDataExtraInfo) SetIcpNo(v string) *DescribeUrlModerationResultResponseBodyDataExtraInfo {
	s.IcpNo = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyDataExtraInfo) SetIcpType(v string) *DescribeUrlModerationResultResponseBodyDataExtraInfo {
	s.IcpType = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyDataExtraInfo) SetSiteType(v string) *DescribeUrlModerationResultResponseBodyDataExtraInfo {
	s.SiteType = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyDataExtraInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeUrlModerationResultResponseBodyDataResult struct {
	// The score of the confidence level. Valid values: 0 to 100. The value is accurate to two decimal places.
	//
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The labels returned after the asynchronous URL moderation.
	//
	// example:
	//
	// sexual_url
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DescribeUrlModerationResultResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeUrlModerationResultResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeUrlModerationResultResponseBodyDataResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *DescribeUrlModerationResultResponseBodyDataResult) GetLabel() *string {
	return s.Label
}

func (s *DescribeUrlModerationResultResponseBodyDataResult) SetConfidence(v float32) *DescribeUrlModerationResultResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyDataResult) SetLabel(v string) *DescribeUrlModerationResultResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
