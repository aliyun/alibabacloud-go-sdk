// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdvanceSecurityEventOperationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AdvanceSecurityEventOperationsResponseBody
	GetRequestId() *string
	SetSecurityEventOperationsResponse(v []*AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) *AdvanceSecurityEventOperationsResponseBody
	GetSecurityEventOperationsResponse() []*AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse
}

type AdvanceSecurityEventOperationsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-XXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The operation performed on the alert event.
	SecurityEventOperationsResponse []*AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse `json:"SecurityEventOperationsResponse,omitempty" xml:"SecurityEventOperationsResponse,omitempty" type:"Repeated"`
}

func (s AdvanceSecurityEventOperationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AdvanceSecurityEventOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *AdvanceSecurityEventOperationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AdvanceSecurityEventOperationsResponseBody) GetSecurityEventOperationsResponse() []*AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	return s.SecurityEventOperationsResponse
}

func (s *AdvanceSecurityEventOperationsResponseBody) SetRequestId(v string) *AdvanceSecurityEventOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBody) SetSecurityEventOperationsResponse(v []*AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) *AdvanceSecurityEventOperationsResponseBody {
	s.SecurityEventOperationsResponse = v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBody) Validate() error {
	if s.SecurityEventOperationsResponse != nil {
		for _, item := range s.SecurityEventOperationsResponse {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse struct {
	// The object on which the operation is performed. This parameter is required when you set the OperationCode parameter to **advance_mark_mis_info**.
	MarkField []*AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField `json:"MarkField,omitempty" xml:"MarkField,omitempty" type:"Repeated"`
	// The metadata configuration returned by the advanced whitelist rule.
	MarkFieldsSource []*AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource `json:"MarkFieldsSource,omitempty" xml:"MarkFieldsSource,omitempty" type:"Repeated"`
	// The operation performed to handle the alert. Valid values:
	//
	// 	- **block_ip**: blocks the alert.
	//
	// 	- **advance_mark_mis_info**: adds the alert to the whitelist.
	//
	// 	- **ignore**: ignores the alert.
	//
	// 	- **manual_handled**: marks the alert as manually handled.
	//
	// 	- **kill_process**: terminates the malicious process.
	//
	// 	- **cleanup**: performs in-depth virus detection and removal.
	//
	// 	- **kill_and_quara**: performs virus detection and removal.
	//
	// 	- **disable_malicious_defense**: turns off malicious defense behavior.
	//
	// 	- **client_problem_check**: performs troubleshooting.
	//
	// 	- **quara**: performs quarantine operations.
	//
	// example:
	//
	// advance_mark_mis_info
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The configuration of the operation performed to handle the alert event.
	//
	// example:
	//
	// {\\"subOperation\\":\\"killByMd5andPath\\"}
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	// Indicates whether the operation can be performed.
	//
	// 	- **true**: The operation can be performed.
	//
	// 	- **false**: The operation cannot be performed.
	//
	// example:
	//
	// false
	UserCanOperate *bool `json:"UserCanOperate,omitempty" xml:"UserCanOperate,omitempty"`
}

func (s AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) String() string {
	return dara.Prettify(s)
}

func (s AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GoString() string {
	return s.String()
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GetMarkField() []*AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	return s.MarkField
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GetMarkFieldsSource() []*AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	return s.MarkFieldsSource
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GetOperationCode() *string {
	return s.OperationCode
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GetOperationParams() *string {
	return s.OperationParams
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GetUserCanOperate() *bool {
	return s.UserCanOperate
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetMarkField(v []*AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.MarkField = v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetMarkFieldsSource(v []*AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.MarkFieldsSource = v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetOperationCode(v string) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.OperationCode = &v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetOperationParams(v string) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.OperationParams = &v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetUserCanOperate(v bool) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.UserCanOperate = &v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponse) Validate() error {
	if s.MarkField != nil {
		for _, item := range s.MarkField {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MarkFieldsSource != nil {
		for _, item := range s.MarkFieldsSource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField struct {
	// The alias of the field that is used in the whitelist rule.
	//
	// example:
	//
	// file path
	FiledAliasName *string `json:"FiledAliasName,omitempty" xml:"FiledAliasName,omitempty"`
	// The field that is used in the whitelist rule.
	//
	// example:
	//
	// filePath
	FiledName *string `json:"FiledName,omitempty" xml:"FiledName,omitempty"`
	// The operation that is used in the whitelist rule. Valid values:
	//
	// 	- **contains**: contains
	//
	// 	- **notContains**: does not contain
	//
	// 	- **regex**: regular expression
	//
	// 	- **strEqual**: equals
	//
	// 	- **strNotEqual**: does not equal
	//
	// example:
	//
	// contains
	MarkMisType *string `json:"MarkMisType,omitempty" xml:"MarkMisType,omitempty"`
	// The value of the field that is used in the whitelist rule.
	//
	// example:
	//
	// 2022-04-25 10:11:04
	MarkMisValue *string `json:"MarkMisValue,omitempty" xml:"MarkMisValue,omitempty"`
	// The operation that is used and can be modified in the whitelist rule. Valid values:
	//
	// 	- **contains**: contains
	//
	// 	- **notContains**: does not contain
	//
	// 	- **regex**: regular expression
	//
	// 	- **strEqual**: equals
	//
	// 	- **strNotEqual**: does not equal
	SupportedMisType []*string `json:"SupportedMisType,omitempty" xml:"SupportedMisType,omitempty" type:"Repeated"`
}

func (s AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) String() string {
	return dara.Prettify(s)
}

func (s AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GoString() string {
	return s.String()
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GetFiledAliasName() *string {
	return s.FiledAliasName
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GetFiledName() *string {
	return s.FiledName
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GetMarkMisType() *string {
	return s.MarkMisType
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GetMarkMisValue() *string {
	return s.MarkMisValue
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GetSupportedMisType() []*string {
	return s.SupportedMisType
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetFiledAliasName(v string) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.FiledAliasName = &v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetFiledName(v string) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.FiledName = &v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetMarkMisType(v string) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.MarkMisType = &v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetMarkMisValue(v string) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.MarkMisValue = &v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetSupportedMisType(v []*string) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.SupportedMisType = v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) Validate() error {
	return dara.Validate(s)
}

type AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource struct {
	// The alias of the field that can be used in the whitelist rule.
	//
	// example:
	//
	// file path
	FiledAliasName *string `json:"FiledAliasName,omitempty" xml:"FiledAliasName,omitempty"`
	// The field that can be used in the whitelist rule.
	//
	// example:
	//
	// filePath
	FiledName *string `json:"FiledName,omitempty" xml:"FiledName,omitempty"`
	// The value of the field that can be used in the whitelist rule.
	//
	// example:
	//
	// contains
	MarkMisValue *string `json:"MarkMisValue,omitempty" xml:"MarkMisValue,omitempty"`
	// The operation that is supported in the whitelist rule. Valid values:
	//
	// 	- **contains**: contains
	//
	// 	- **notContains**: does not contain
	//
	// 	- **regex**: regular expression
	//
	// 	- **strEqual**: equals
	//
	// 	- **strNotEqual**: does not equal
	SupportedMisType []*string `json:"SupportedMisType,omitempty" xml:"SupportedMisType,omitempty" type:"Repeated"`
}

func (s AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) String() string {
	return dara.Prettify(s)
}

func (s AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) GoString() string {
	return s.String()
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) GetFiledAliasName() *string {
	return s.FiledAliasName
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) GetFiledName() *string {
	return s.FiledName
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) GetMarkMisValue() *string {
	return s.MarkMisValue
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) GetSupportedMisType() []*string {
	return s.SupportedMisType
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) SetFiledAliasName(v string) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	s.FiledAliasName = &v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) SetFiledName(v string) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	s.FiledName = &v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) SetMarkMisValue(v string) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	s.MarkMisValue = &v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) SetSupportedMisType(v []*string) *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	s.SupportedMisType = v
	return s
}

func (s *AdvanceSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) Validate() error {
	return dara.Validate(s)
}
