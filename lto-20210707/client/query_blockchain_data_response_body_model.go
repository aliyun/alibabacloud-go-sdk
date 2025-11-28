// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBlockchainDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryBlockchainDataResponseBody
	GetCode() *string
	SetData(v *QueryBlockchainDataResponseBodyData) *QueryBlockchainDataResponseBody
	GetData() *QueryBlockchainDataResponseBodyData
	SetMessage(v string) *QueryBlockchainDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryBlockchainDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryBlockchainDataResponseBody
	GetSuccess() *bool
}

type QueryBlockchainDataResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryBlockchainDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryBlockchainDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBlockchainDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBlockchainDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryBlockchainDataResponseBody) GetData() *QueryBlockchainDataResponseBodyData {
	return s.Data
}

func (s *QueryBlockchainDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryBlockchainDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBlockchainDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryBlockchainDataResponseBody) SetCode(v string) *QueryBlockchainDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBlockchainDataResponseBody) SetData(v *QueryBlockchainDataResponseBodyData) *QueryBlockchainDataResponseBody {
	s.Data = v
	return s
}

func (s *QueryBlockchainDataResponseBody) SetMessage(v string) *QueryBlockchainDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBlockchainDataResponseBody) SetRequestId(v string) *QueryBlockchainDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBlockchainDataResponseBody) SetSuccess(v bool) *QueryBlockchainDataResponseBody {
	s.Success = &v
	return s
}

func (s *QueryBlockchainDataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryBlockchainDataResponseBodyData struct {
	AlgType       *string `json:"AlgType,omitempty" xml:"AlgType,omitempty"`
	PlainData     *string `json:"PlainData,omitempty" xml:"PlainData,omitempty"`
	PrivacyData   *string `json:"PrivacyData,omitempty" xml:"PrivacyData,omitempty"`
	PrivacyRuleId *string `json:"PrivacyRuleId,omitempty" xml:"PrivacyRuleId,omitempty"`
}

func (s QueryBlockchainDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryBlockchainDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBlockchainDataResponseBodyData) GetAlgType() *string {
	return s.AlgType
}

func (s *QueryBlockchainDataResponseBodyData) GetPlainData() *string {
	return s.PlainData
}

func (s *QueryBlockchainDataResponseBodyData) GetPrivacyData() *string {
	return s.PrivacyData
}

func (s *QueryBlockchainDataResponseBodyData) GetPrivacyRuleId() *string {
	return s.PrivacyRuleId
}

func (s *QueryBlockchainDataResponseBodyData) SetAlgType(v string) *QueryBlockchainDataResponseBodyData {
	s.AlgType = &v
	return s
}

func (s *QueryBlockchainDataResponseBodyData) SetPlainData(v string) *QueryBlockchainDataResponseBodyData {
	s.PlainData = &v
	return s
}

func (s *QueryBlockchainDataResponseBodyData) SetPrivacyData(v string) *QueryBlockchainDataResponseBodyData {
	s.PrivacyData = &v
	return s
}

func (s *QueryBlockchainDataResponseBodyData) SetPrivacyRuleId(v string) *QueryBlockchainDataResponseBodyData {
	s.PrivacyRuleId = &v
	return s
}

func (s *QueryBlockchainDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}
