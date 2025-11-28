// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllPrivacyRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllPrivacyRuleResponseBody
	GetCode() *string
	SetData(v []*ListAllPrivacyRuleResponseBodyData) *ListAllPrivacyRuleResponseBody
	GetData() []*ListAllPrivacyRuleResponseBodyData
	SetHttpStatusCode(v int32) *ListAllPrivacyRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAllPrivacyRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllPrivacyRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllPrivacyRuleResponseBody
	GetSuccess() *bool
}

type ListAllPrivacyRuleResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListAllPrivacyRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAllPrivacyRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllPrivacyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllPrivacyRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllPrivacyRuleResponseBody) GetData() []*ListAllPrivacyRuleResponseBodyData {
	return s.Data
}

func (s *ListAllPrivacyRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAllPrivacyRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllPrivacyRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllPrivacyRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllPrivacyRuleResponseBody) SetCode(v string) *ListAllPrivacyRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllPrivacyRuleResponseBody) SetData(v []*ListAllPrivacyRuleResponseBodyData) *ListAllPrivacyRuleResponseBody {
	s.Data = v
	return s
}

func (s *ListAllPrivacyRuleResponseBody) SetHttpStatusCode(v int32) *ListAllPrivacyRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAllPrivacyRuleResponseBody) SetMessage(v string) *ListAllPrivacyRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllPrivacyRuleResponseBody) SetRequestId(v string) *ListAllPrivacyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllPrivacyRuleResponseBody) SetSuccess(v bool) *ListAllPrivacyRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllPrivacyRuleResponseBody) Validate() error {
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

type ListAllPrivacyRuleResponseBodyData struct {
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PrivacyRuleId *string `json:"PrivacyRuleId,omitempty" xml:"PrivacyRuleId,omitempty"`
}

func (s ListAllPrivacyRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllPrivacyRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllPrivacyRuleResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAllPrivacyRuleResponseBodyData) GetPrivacyRuleId() *string {
	return s.PrivacyRuleId
}

func (s *ListAllPrivacyRuleResponseBodyData) SetName(v string) *ListAllPrivacyRuleResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAllPrivacyRuleResponseBodyData) SetPrivacyRuleId(v string) *ListAllPrivacyRuleResponseBodyData {
	s.PrivacyRuleId = &v
	return s
}

func (s *ListAllPrivacyRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
