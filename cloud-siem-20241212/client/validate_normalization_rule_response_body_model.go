// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateNormalizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidateNormalizationRuleResponseBody
	GetRequestId() *string
	SetValidateResult(v []*ValidateNormalizationRuleResponseBodyValidateResult) *ValidateNormalizationRuleResponseBody
	GetValidateResult() []*ValidateNormalizationRuleResponseBodyValidateResult
}

type ValidateNormalizationRuleResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ValidateResult []*ValidateNormalizationRuleResponseBodyValidateResult `json:"ValidateResult,omitempty" xml:"ValidateResult,omitempty" type:"Repeated"`
}

func (s ValidateNormalizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateNormalizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateNormalizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateNormalizationRuleResponseBody) GetValidateResult() []*ValidateNormalizationRuleResponseBodyValidateResult {
	return s.ValidateResult
}

func (s *ValidateNormalizationRuleResponseBody) SetRequestId(v string) *ValidateNormalizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBody) SetValidateResult(v []*ValidateNormalizationRuleResponseBodyValidateResult) *ValidateNormalizationRuleResponseBody {
	s.ValidateResult = v
	return s
}

func (s *ValidateNormalizationRuleResponseBody) Validate() error {
	if s.ValidateResult != nil {
		for _, item := range s.ValidateResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ValidateNormalizationRuleResponseBodyValidateResult struct {
	// example:
	//
	// host。
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// example:
	//
	// ze。
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// example:
	//
	// success。
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// host。
	NormalizationFieldName *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	// example:
	//
	// true。
	NormalizationFieldRequired *bool `json:"NormalizationFieldRequired,omitempty" xml:"NormalizationFieldRequired,omitempty"`
	// example:
	//
	// 1。
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateNormalizationRuleResponseBodyValidateResult) String() string {
	return dara.Prettify(s)
}

func (s ValidateNormalizationRuleResponseBodyValidateResult) GoString() string {
	return s.String()
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetFieldName() *string {
	return s.FieldName
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetFieldValue() *string {
	return s.FieldValue
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetMessage() *string {
	return s.Message
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetNormalizationFieldRequired() *bool {
	return s.NormalizationFieldRequired
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetResult() *int32 {
	return s.Result
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetFieldName(v string) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.FieldName = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetFieldValue(v string) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.FieldValue = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetMessage(v string) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.Message = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetNormalizationFieldName(v string) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.NormalizationFieldName = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetNormalizationFieldRequired(v bool) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.NormalizationFieldRequired = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetResult(v int32) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.Result = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) Validate() error {
	return dara.Validate(s)
}
