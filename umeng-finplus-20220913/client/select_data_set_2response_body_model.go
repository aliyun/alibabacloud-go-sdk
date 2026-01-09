// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectDataSet2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SelectDataSet2ResponseBody
	GetCode() *string
	SetData(v *SelectDataSet2ResponseBodyData) *SelectDataSet2ResponseBody
	GetData() *SelectDataSet2ResponseBodyData
	SetMsg(v string) *SelectDataSet2ResponseBody
	GetMsg() *string
	SetRequestId(v string) *SelectDataSet2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SelectDataSet2ResponseBody
	GetSuccess() *bool
}

type SelectDataSet2ResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SelectDataSet2ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                         `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SelectDataSet2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SelectDataSet2ResponseBody) GoString() string {
	return s.String()
}

func (s *SelectDataSet2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *SelectDataSet2ResponseBody) GetData() *SelectDataSet2ResponseBodyData {
	return s.Data
}

func (s *SelectDataSet2ResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SelectDataSet2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SelectDataSet2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SelectDataSet2ResponseBody) SetCode(v string) *SelectDataSet2ResponseBody {
	s.Code = &v
	return s
}

func (s *SelectDataSet2ResponseBody) SetData(v *SelectDataSet2ResponseBodyData) *SelectDataSet2ResponseBody {
	s.Data = v
	return s
}

func (s *SelectDataSet2ResponseBody) SetMsg(v string) *SelectDataSet2ResponseBody {
	s.Msg = &v
	return s
}

func (s *SelectDataSet2ResponseBody) SetRequestId(v string) *SelectDataSet2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *SelectDataSet2ResponseBody) SetSuccess(v bool) *SelectDataSet2ResponseBody {
	s.Success = &v
	return s
}

func (s *SelectDataSet2ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SelectDataSet2ResponseBodyData struct {
	CreateTime   *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DatasetId    *int64  `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	LineNum      *int64  `json:"lineNum,omitempty" xml:"lineNum,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	OssFileCount *int64  `json:"ossFileCount,omitempty" xml:"ossFileCount,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusMsg    *string `json:"statusMsg,omitempty" xml:"statusMsg,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SelectDataSet2ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SelectDataSet2ResponseBodyData) GoString() string {
	return s.String()
}

func (s *SelectDataSet2ResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SelectDataSet2ResponseBodyData) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *SelectDataSet2ResponseBodyData) GetLineNum() *int64 {
	return s.LineNum
}

func (s *SelectDataSet2ResponseBodyData) GetName() *string {
	return s.Name
}

func (s *SelectDataSet2ResponseBodyData) GetOssFileCount() *int64 {
	return s.OssFileCount
}

func (s *SelectDataSet2ResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SelectDataSet2ResponseBodyData) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *SelectDataSet2ResponseBodyData) GetType() *string {
	return s.Type
}

func (s *SelectDataSet2ResponseBodyData) SetCreateTime(v string) *SelectDataSet2ResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *SelectDataSet2ResponseBodyData) SetDatasetId(v int64) *SelectDataSet2ResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *SelectDataSet2ResponseBodyData) SetLineNum(v int64) *SelectDataSet2ResponseBodyData {
	s.LineNum = &v
	return s
}

func (s *SelectDataSet2ResponseBodyData) SetName(v string) *SelectDataSet2ResponseBodyData {
	s.Name = &v
	return s
}

func (s *SelectDataSet2ResponseBodyData) SetOssFileCount(v int64) *SelectDataSet2ResponseBodyData {
	s.OssFileCount = &v
	return s
}

func (s *SelectDataSet2ResponseBodyData) SetStatus(v string) *SelectDataSet2ResponseBodyData {
	s.Status = &v
	return s
}

func (s *SelectDataSet2ResponseBodyData) SetStatusMsg(v string) *SelectDataSet2ResponseBodyData {
	s.StatusMsg = &v
	return s
}

func (s *SelectDataSet2ResponseBodyData) SetType(v string) *SelectDataSet2ResponseBodyData {
	s.Type = &v
	return s
}

func (s *SelectDataSet2ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
