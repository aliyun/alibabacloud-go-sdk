// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSet2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataSet2ResponseBody
	GetCode() *string
	SetData(v *ListDataSet2ResponseBodyData) *ListDataSet2ResponseBody
	GetData() *ListDataSet2ResponseBodyData
	SetMsg(v string) *ListDataSet2ResponseBody
	GetMsg() *string
	SetRequestId(v string) *ListDataSet2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataSet2ResponseBody
	GetSuccess() *bool
}

type ListDataSet2ResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListDataSet2ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                       `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataSet2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSet2ResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSet2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataSet2ResponseBody) GetData() *ListDataSet2ResponseBodyData {
	return s.Data
}

func (s *ListDataSet2ResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ListDataSet2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSet2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataSet2ResponseBody) SetCode(v string) *ListDataSet2ResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataSet2ResponseBody) SetData(v *ListDataSet2ResponseBodyData) *ListDataSet2ResponseBody {
	s.Data = v
	return s
}

func (s *ListDataSet2ResponseBody) SetMsg(v string) *ListDataSet2ResponseBody {
	s.Msg = &v
	return s
}

func (s *ListDataSet2ResponseBody) SetRequestId(v string) *ListDataSet2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSet2ResponseBody) SetSuccess(v bool) *ListDataSet2ResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataSet2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataSet2ResponseBodyData struct {
	Data     []*ListDataSet2ResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	TotalNum *int32                              `json:"totalNum,omitempty" xml:"totalNum,omitempty"`
}

func (s ListDataSet2ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataSet2ResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataSet2ResponseBodyData) GetData() []*ListDataSet2ResponseBodyDataData {
	return s.Data
}

func (s *ListDataSet2ResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListDataSet2ResponseBodyData) SetData(v []*ListDataSet2ResponseBodyDataData) *ListDataSet2ResponseBodyData {
	s.Data = v
	return s
}

func (s *ListDataSet2ResponseBodyData) SetTotalNum(v int32) *ListDataSet2ResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *ListDataSet2ResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDataSet2ResponseBodyDataData struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DatasetId  *int64  `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	LineNum    *int64  `json:"lineNum,omitempty" xml:"lineNum,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataSet2ResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s ListDataSet2ResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListDataSet2ResponseBodyDataData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDataSet2ResponseBodyDataData) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *ListDataSet2ResponseBodyDataData) GetLineNum() *int64 {
	return s.LineNum
}

func (s *ListDataSet2ResponseBodyDataData) GetName() *string {
	return s.Name
}

func (s *ListDataSet2ResponseBodyDataData) GetStatus() *string {
	return s.Status
}

func (s *ListDataSet2ResponseBodyDataData) GetType() *string {
	return s.Type
}

func (s *ListDataSet2ResponseBodyDataData) SetCreateTime(v string) *ListDataSet2ResponseBodyDataData {
	s.CreateTime = &v
	return s
}

func (s *ListDataSet2ResponseBodyDataData) SetDatasetId(v int64) *ListDataSet2ResponseBodyDataData {
	s.DatasetId = &v
	return s
}

func (s *ListDataSet2ResponseBodyDataData) SetLineNum(v int64) *ListDataSet2ResponseBodyDataData {
	s.LineNum = &v
	return s
}

func (s *ListDataSet2ResponseBodyDataData) SetName(v string) *ListDataSet2ResponseBodyDataData {
	s.Name = &v
	return s
}

func (s *ListDataSet2ResponseBodyDataData) SetStatus(v string) *ListDataSet2ResponseBodyDataData {
	s.Status = &v
	return s
}

func (s *ListDataSet2ResponseBodyDataData) SetType(v string) *ListDataSet2ResponseBodyDataData {
	s.Type = &v
	return s
}

func (s *ListDataSet2ResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}
