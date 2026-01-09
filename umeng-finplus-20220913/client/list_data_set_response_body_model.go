// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataSetResponseBody
	GetCode() *string
	SetData(v []*ListDataSetResponseBodyData) *ListDataSetResponseBody
	GetData() []*ListDataSetResponseBodyData
	SetMsg(v string) *ListDataSetResponseBody
	GetMsg() *string
	SetRequestId(v string) *ListDataSetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataSetResponseBody
	GetSuccess() *bool
}

type ListDataSetResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListDataSetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Msg       *string                        `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSetResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataSetResponseBody) GetData() []*ListDataSetResponseBodyData {
	return s.Data
}

func (s *ListDataSetResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ListDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataSetResponseBody) SetCode(v string) *ListDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataSetResponseBody) SetData(v []*ListDataSetResponseBodyData) *ListDataSetResponseBody {
	s.Data = v
	return s
}

func (s *ListDataSetResponseBody) SetMsg(v string) *ListDataSetResponseBody {
	s.Msg = &v
	return s
}

func (s *ListDataSetResponseBody) SetRequestId(v string) *ListDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSetResponseBody) SetSuccess(v bool) *ListDataSetResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataSetResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSetResponseBodyData struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DatasetId  *int64  `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	LineNum    *int64  `json:"lineNum,omitempty" xml:"lineNum,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataSetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataSetResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDataSetResponseBodyData) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *ListDataSetResponseBodyData) GetLineNum() *int64 {
	return s.LineNum
}

func (s *ListDataSetResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListDataSetResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListDataSetResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListDataSetResponseBodyData) SetCreateTime(v string) *ListDataSetResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListDataSetResponseBodyData) SetDatasetId(v int64) *ListDataSetResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *ListDataSetResponseBodyData) SetLineNum(v int64) *ListDataSetResponseBodyData {
	s.LineNum = &v
	return s
}

func (s *ListDataSetResponseBodyData) SetName(v string) *ListDataSetResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListDataSetResponseBodyData) SetStatus(v string) *ListDataSetResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDataSetResponseBodyData) SetType(v string) *ListDataSetResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListDataSetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
