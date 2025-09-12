// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SelectDataSetResponseBody
	GetCode() *string
	SetData(v *SelectDataSetResponseBodyData) *SelectDataSetResponseBody
	GetData() *SelectDataSetResponseBodyData
	SetMsg(v string) *SelectDataSetResponseBody
	GetMsg() *string
	SetRequestId(v string) *SelectDataSetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SelectDataSetResponseBody
	GetSuccess() *bool
}

type SelectDataSetResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SelectDataSetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                        `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SelectDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SelectDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *SelectDataSetResponseBody) GetCode() *string {
	return s.Code
}

func (s *SelectDataSetResponseBody) GetData() *SelectDataSetResponseBodyData {
	return s.Data
}

func (s *SelectDataSetResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SelectDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SelectDataSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SelectDataSetResponseBody) SetCode(v string) *SelectDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *SelectDataSetResponseBody) SetData(v *SelectDataSetResponseBodyData) *SelectDataSetResponseBody {
	s.Data = v
	return s
}

func (s *SelectDataSetResponseBody) SetMsg(v string) *SelectDataSetResponseBody {
	s.Msg = &v
	return s
}

func (s *SelectDataSetResponseBody) SetRequestId(v string) *SelectDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *SelectDataSetResponseBody) SetSuccess(v bool) *SelectDataSetResponseBody {
	s.Success = &v
	return s
}

func (s *SelectDataSetResponseBody) Validate() error {
	return dara.Validate(s)
}

type SelectDataSetResponseBodyData struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DatasetId  *int64  `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	LineNum    *int64  `json:"lineNum,omitempty" xml:"lineNum,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SelectDataSetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SelectDataSetResponseBodyData) GoString() string {
	return s.String()
}

func (s *SelectDataSetResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SelectDataSetResponseBodyData) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *SelectDataSetResponseBodyData) GetLineNum() *int64 {
	return s.LineNum
}

func (s *SelectDataSetResponseBodyData) GetName() *string {
	return s.Name
}

func (s *SelectDataSetResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SelectDataSetResponseBodyData) GetType() *string {
	return s.Type
}

func (s *SelectDataSetResponseBodyData) SetCreateTime(v string) *SelectDataSetResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *SelectDataSetResponseBodyData) SetDatasetId(v int64) *SelectDataSetResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *SelectDataSetResponseBodyData) SetLineNum(v int64) *SelectDataSetResponseBodyData {
	s.LineNum = &v
	return s
}

func (s *SelectDataSetResponseBodyData) SetName(v string) *SelectDataSetResponseBodyData {
	s.Name = &v
	return s
}

func (s *SelectDataSetResponseBodyData) SetStatus(v string) *SelectDataSetResponseBodyData {
	s.Status = &v
	return s
}

func (s *SelectDataSetResponseBodyData) SetType(v string) *SelectDataSetResponseBodyData {
	s.Type = &v
	return s
}

func (s *SelectDataSetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
