// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleTagListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityRuleTagListResponseBody
	GetCode() *string
	SetData(v []*GetQualityRuleTagListResponseBodyData) *GetQualityRuleTagListResponseBody
	GetData() []*GetQualityRuleTagListResponseBodyData
	SetMessage(v string) *GetQualityRuleTagListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityRuleTagListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityRuleTagListResponseBody
	GetSuccess() *bool
}

type GetQualityRuleTagListResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetQualityRuleTagListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityRuleTagListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleTagListResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTagListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityRuleTagListResponseBody) GetData() []*GetQualityRuleTagListResponseBodyData {
	return s.Data
}

func (s *GetQualityRuleTagListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityRuleTagListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityRuleTagListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityRuleTagListResponseBody) SetCode(v string) *GetQualityRuleTagListResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityRuleTagListResponseBody) SetData(v []*GetQualityRuleTagListResponseBodyData) *GetQualityRuleTagListResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityRuleTagListResponseBody) SetMessage(v string) *GetQualityRuleTagListResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityRuleTagListResponseBody) SetRequestId(v string) *GetQualityRuleTagListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityRuleTagListResponseBody) SetSuccess(v bool) *GetQualityRuleTagListResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityRuleTagListResponseBody) Validate() error {
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

type GetQualityRuleTagListResponseBodyData struct {
	RuleTagId   *int64  `json:"RuleTagId,omitempty" xml:"RuleTagId,omitempty"`
	RuleTagName *string `json:"RuleTagName,omitempty" xml:"RuleTagName,omitempty"`
}

func (s GetQualityRuleTagListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleTagListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTagListResponseBodyData) GetRuleTagId() *int64 {
	return s.RuleTagId
}

func (s *GetQualityRuleTagListResponseBodyData) GetRuleTagName() *string {
	return s.RuleTagName
}

func (s *GetQualityRuleTagListResponseBodyData) SetRuleTagId(v int64) *GetQualityRuleTagListResponseBodyData {
	s.RuleTagId = &v
	return s
}

func (s *GetQualityRuleTagListResponseBodyData) SetRuleTagName(v string) *GetQualityRuleTagListResponseBodyData {
	s.RuleTagName = &v
	return s
}

func (s *GetQualityRuleTagListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
