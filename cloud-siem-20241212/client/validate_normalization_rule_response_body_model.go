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
	// The ID of the request.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of validation results.
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
	// The name of the field.
	//
	// example:
	//
	// host
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The field information.
	//
	// example:
	//
	// ze
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The name of the log field.
	//
	// example:
	//
	// aaa
	LogFieldName *string `json:"LogFieldName,omitempty" xml:"LogFieldName,omitempty"`
	// The value of the log field.
	//
	// example:
	//
	// bbb
	LogFieldValue *string `json:"LogFieldValue,omitempty" xml:"LogFieldValue,omitempty"`
	// The reason for the validation result. Valid values:
	//
	// - OperationDenied.TheValueIsRequired: A required parameter is empty.
	//
	// - OperationDenied.TheValueIsNull: The parameter value is null.
	//
	// - OperationDenied.TheEnumValueNotSupport: The field value is not within the valid enum values.
	//
	// - OperationDenied.TheValueLessThanMin: The field value is less than the specified minimum value.
	//
	// - OperationDenied.TheValueMoreThanMax: The field value is greater than the specified maximum value.
	//
	// - OperationDenied.TheValueNotMatchRegularExpression: The field value fails the regular expression match.
	//
	// - success: The validation is passed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The source of the normalized field. Valid values:
	//
	// - preset: Built-in.
	//
	// - custom: Custom.
	//
	// example:
	//
	// preset
	NormalizationFieldFrom *string `json:"NormalizationFieldFrom,omitempty" xml:"NormalizationFieldFrom,omitempty"`
	// The name of the normalized field.
	//
	// example:
	//
	// host
	NormalizationFieldName *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	// Specifies whether the normalized field is required.
	//
	// example:
	//
	// true
	NormalizationFieldRequired *bool `json:"NormalizationFieldRequired,omitempty" xml:"NormalizationFieldRequired,omitempty"`
	// Specifies whether the normalized field name is built into the system.
	NormalizationFieldReserved *bool `json:"NormalizationFieldReserved,omitempty" xml:"NormalizationFieldReserved,omitempty"`
	// The type of the normalized field. Valid values: text, long, double, and json.
	//
	// example:
	//
	// text
	NormalizationFieldType *string `json:"NormalizationFieldType,omitempty" xml:"NormalizationFieldType,omitempty"`
	// The reason why the normalized field fails the validation.
	//
	// example:
	//
	// OperationDenied.TheValueIsRequired
	NormalizationFieldValidationReason *string `json:"NormalizationFieldValidationReason,omitempty" xml:"NormalizationFieldValidationReason,omitempty"`
	// The validation status of the normalized field. Valid values: pass and fail.
	//
	// example:
	//
	// pass
	NormalizationFieldValidationStatus *string `json:"NormalizationFieldValidationStatus,omitempty" xml:"NormalizationFieldValidationStatus,omitempty"`
	// The validation result. Valid values:
	//
	// - 1: The validation is passed.
	//
	// - 0: A warning is returned.
	//
	// - 1: The validation fails.
	//
	// example:
	//
	// 1
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

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetLogFieldName() *string {
	return s.LogFieldName
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetLogFieldValue() *string {
	return s.LogFieldValue
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetMessage() *string {
	return s.Message
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetNormalizationFieldFrom() *string {
	return s.NormalizationFieldFrom
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetNormalizationFieldRequired() *bool {
	return s.NormalizationFieldRequired
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetNormalizationFieldReserved() *bool {
	return s.NormalizationFieldReserved
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetNormalizationFieldType() *string {
	return s.NormalizationFieldType
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetNormalizationFieldValidationReason() *string {
	return s.NormalizationFieldValidationReason
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) GetNormalizationFieldValidationStatus() *string {
	return s.NormalizationFieldValidationStatus
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

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetLogFieldName(v string) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.LogFieldName = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetLogFieldValue(v string) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.LogFieldValue = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetMessage(v string) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.Message = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetNormalizationFieldFrom(v string) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.NormalizationFieldFrom = &v
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

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetNormalizationFieldReserved(v bool) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.NormalizationFieldReserved = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetNormalizationFieldType(v string) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.NormalizationFieldType = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetNormalizationFieldValidationReason(v string) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.NormalizationFieldValidationReason = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetNormalizationFieldValidationStatus(v string) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.NormalizationFieldValidationStatus = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) SetResult(v int32) *ValidateNormalizationRuleResponseBodyValidateResult {
	s.Result = &v
	return s
}

func (s *ValidateNormalizationRuleResponseBodyValidateResult) Validate() error {
	return dara.Validate(s)
}
