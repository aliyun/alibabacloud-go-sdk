// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnionMediaIndustryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListUnionMediaIndustryResponseBody
	GetCode() *int32
	SetData(v []*ListUnionMediaIndustryResponseBodyData) *ListUnionMediaIndustryResponseBody
	GetData() []*ListUnionMediaIndustryResponseBodyData
	SetErrorMsg(v string) *ListUnionMediaIndustryResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *ListUnionMediaIndustryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUnionMediaIndustryResponseBody
	GetSuccess() *bool
}

type ListUnionMediaIndustryResponseBody struct {
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListUnionMediaIndustryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg  *string                                   `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUnionMediaIndustryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUnionMediaIndustryResponseBody) GoString() string {
	return s.String()
}

func (s *ListUnionMediaIndustryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListUnionMediaIndustryResponseBody) GetData() []*ListUnionMediaIndustryResponseBodyData {
	return s.Data
}

func (s *ListUnionMediaIndustryResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListUnionMediaIndustryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUnionMediaIndustryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUnionMediaIndustryResponseBody) SetCode(v int32) *ListUnionMediaIndustryResponseBody {
	s.Code = &v
	return s
}

func (s *ListUnionMediaIndustryResponseBody) SetData(v []*ListUnionMediaIndustryResponseBodyData) *ListUnionMediaIndustryResponseBody {
	s.Data = v
	return s
}

func (s *ListUnionMediaIndustryResponseBody) SetErrorMsg(v string) *ListUnionMediaIndustryResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListUnionMediaIndustryResponseBody) SetRequestId(v string) *ListUnionMediaIndustryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUnionMediaIndustryResponseBody) SetSuccess(v bool) *ListUnionMediaIndustryResponseBody {
	s.Success = &v
	return s
}

func (s *ListUnionMediaIndustryResponseBody) Validate() error {
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

type ListUnionMediaIndustryResponseBodyData struct {
	IndustryCode       *string `json:"IndustryCode,omitempty" xml:"IndustryCode,omitempty"`
	IndustryName       *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	Level              *string `json:"Level,omitempty" xml:"Level,omitempty"`
	ParentIndustryCode *string `json:"ParentIndustryCode,omitempty" xml:"ParentIndustryCode,omitempty"`
	UsedFlag           *bool   `json:"UsedFlag,omitempty" xml:"UsedFlag,omitempty"`
}

func (s ListUnionMediaIndustryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUnionMediaIndustryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUnionMediaIndustryResponseBodyData) GetIndustryCode() *string {
	return s.IndustryCode
}

func (s *ListUnionMediaIndustryResponseBodyData) GetIndustryName() *string {
	return s.IndustryName
}

func (s *ListUnionMediaIndustryResponseBodyData) GetLevel() *string {
	return s.Level
}

func (s *ListUnionMediaIndustryResponseBodyData) GetParentIndustryCode() *string {
	return s.ParentIndustryCode
}

func (s *ListUnionMediaIndustryResponseBodyData) GetUsedFlag() *bool {
	return s.UsedFlag
}

func (s *ListUnionMediaIndustryResponseBodyData) SetIndustryCode(v string) *ListUnionMediaIndustryResponseBodyData {
	s.IndustryCode = &v
	return s
}

func (s *ListUnionMediaIndustryResponseBodyData) SetIndustryName(v string) *ListUnionMediaIndustryResponseBodyData {
	s.IndustryName = &v
	return s
}

func (s *ListUnionMediaIndustryResponseBodyData) SetLevel(v string) *ListUnionMediaIndustryResponseBodyData {
	s.Level = &v
	return s
}

func (s *ListUnionMediaIndustryResponseBodyData) SetParentIndustryCode(v string) *ListUnionMediaIndustryResponseBodyData {
	s.ParentIndustryCode = &v
	return s
}

func (s *ListUnionMediaIndustryResponseBodyData) SetUsedFlag(v bool) *ListUnionMediaIndustryResponseBodyData {
	s.UsedFlag = &v
	return s
}

func (s *ListUnionMediaIndustryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
