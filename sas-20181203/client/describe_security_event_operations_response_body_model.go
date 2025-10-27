// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventOperationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSecurityEventOperationsResponseBody
	GetRequestId() *string
	SetSecurityEventOperationsResponse(v []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) *DescribeSecurityEventOperationsResponseBody
	GetSecurityEventOperationsResponse() []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse
}

type DescribeSecurityEventOperationsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// B7A2000F-497E-5DA0-B14D-615CD410DD7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The operations that are performed to handle the alert.
	SecurityEventOperationsResponse []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse `json:"SecurityEventOperationsResponse,omitempty" xml:"SecurityEventOperationsResponse,omitempty" type:"Repeated"`
}

func (s DescribeSecurityEventOperationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityEventOperationsResponseBody) GetSecurityEventOperationsResponse() []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	return s.SecurityEventOperationsResponse
}

func (s *DescribeSecurityEventOperationsResponseBody) SetRequestId(v string) *DescribeSecurityEventOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBody) SetSecurityEventOperationsResponse(v []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) *DescribeSecurityEventOperationsResponseBody {
	s.SecurityEventOperationsResponse = v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBody) Validate() error {
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

type DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse struct {
	// The objects on which the operations are performed. This parameter is required when you add the alert to the whitelist by configuring precise defense rules.
	MappingMarkFields []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields `json:"MappingMarkFields,omitempty" xml:"MappingMarkFields,omitempty" type:"Repeated"`
	// The configurations that are used when the value of the OperationCode parameter is **advance_mark_mis_info**.
	MarkField []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField `json:"MarkField,omitempty" xml:"MarkField,omitempty" type:"Repeated"`
	// The configuration items that can be used when the value of the OperationCode parameter is advance_mark_mis_info.
	MarkFieldsSource []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource `json:"MarkFieldsSource,omitempty" xml:"MarkFieldsSource,omitempty" type:"Repeated"`
	// The operation that is performed to handle the alert. Valid values:
	//
	// 	- **block_ip**: blocks the source IP address.
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
	// 	- **kill_and_quara**: terminates the malicious process and quarantines the source file.
	//
	// 	- **disable_malicious_defense**: disables the malicious behavior defense feature.
	//
	// 	- **client_problem_check**: performs troubleshooting.
	//
	// 	- **quara**: quarantines the source file of the malicious process.
	//
	// 	- **defense_mark_mis_info**: enables the precise defense feature but disables the notification feature.
	//
	// 	- **rm_defense_mark_mis_info**: enables the notification feature.
	//
	// 	- **rm_mark_mis_info**: removes the alert from the whitelist.
	//
	// 	- **cancle_manual**: cancels marking the alert as manually handled.
	//
	// example:
	//
	// advance_mark_mis_info
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The configuration of the operation that is performed to handle the alert.
	//
	// >  If the value of the **OperationCode*	- parameter is **kill_and_quara*	- or **block_ip**, the OperationParams parameter is required. If the value of the **OperationCode*	- parameter is a different value, the OperationParams parameter can be left empty.
	//
	// example:
	//
	// {"expireTime":1641566807783}
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	// Indicates whether you can handle the alert in the current edition of Security Center. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	UserCanOperate *bool `json:"UserCanOperate,omitempty" xml:"UserCanOperate,omitempty"`
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GetMappingMarkFields() []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields {
	return s.MappingMarkFields
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GetMarkField() []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	return s.MarkField
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GetMarkFieldsSource() []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	return s.MarkFieldsSource
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GetOperationCode() *string {
	return s.OperationCode
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GetOperationParams() *string {
	return s.OperationParams
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GetUserCanOperate() *bool {
	return s.UserCanOperate
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetMappingMarkFields(v []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.MappingMarkFields = v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetMarkField(v []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.MarkField = v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetMarkFieldsSource(v []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.MarkFieldsSource = v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetOperationCode(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.OperationCode = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetOperationParams(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.OperationParams = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetUserCanOperate(v bool) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.UserCanOperate = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) Validate() error {
	if s.MappingMarkFields != nil {
		for _, item := range s.MappingMarkFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields struct {
	// The description of the field that is added to the whitelist.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the value of the field can be changed.
	//
	// 	- **CUSTOM**: The value of the field can be changed.
	//
	// 	- **SYSTEM**: The value of the field cannot be changed.
	//
	// example:
	//
	// CUSTOM
	FillType *string `json:"FillType,omitempty" xml:"FillType,omitempty"`
	// The maximum length of the field that is added to the whitelist.
	//
	// example:
	//
	// 2048
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// The minimum length of the field that is added to the whitelist.
	//
	// example:
	//
	// 1024
	MinLength *int32 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// The name of the field that is added to the whitelist.
	//
	// example:
	//
	// pid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// The display name of the field that can be used in the whitelist rule.
	//
	// example:
	//
	// pid
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The display name of the field that is added to the whitelist.
	//
	// example:
	//
	// 1791
	ShowValue *string `json:"ShowValue,omitempty" xml:"ShowValue,omitempty"`
	// The value of the field that is added to the whitelist.
	//
	// example:
	//
	// 1791
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) GetDescription() *string {
	return s.Description
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) GetFillType() *string {
	return s.FillType
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) GetMinLength() *int32 {
	return s.MinLength
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) GetName() *string {
	return s.Name
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) GetRequired() *bool {
	return s.Required
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) GetShowName() *string {
	return s.ShowName
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) GetShowValue() *string {
	return s.ShowValue
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) GetValue() *string {
	return s.Value
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) SetDescription(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields {
	s.Description = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) SetFillType(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields {
	s.FillType = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) SetMaxLength(v int32) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields {
	s.MaxLength = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) SetMinLength(v int32) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields {
	s.MinLength = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) SetName(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields {
	s.Name = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) SetRequired(v bool) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields {
	s.Required = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) SetShowName(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields {
	s.ShowName = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) SetShowValue(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields {
	s.ShowValue = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) SetValue(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields {
	s.Value = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMappingMarkFields) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField struct {
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
	// gmtModified
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
	// An array consisting of the operations that are supported by the method to add the alert event to the whitelist.
	SupportedMisType []*string `json:"SupportedMisType,omitempty" xml:"SupportedMisType,omitempty" type:"Repeated"`
	// The UUID of the server on which the alert event is detected.
	//
	// example:
	//
	// 3d6b4a75-c28f-447b-9142-38f6252c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GetFiledAliasName() *string {
	return s.FiledAliasName
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GetFiledName() *string {
	return s.FiledName
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GetMarkMisType() *string {
	return s.MarkMisType
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GetMarkMisValue() *string {
	return s.MarkMisValue
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GetSupportedMisType() []*string {
	return s.SupportedMisType
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetFiledAliasName(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.FiledAliasName = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetFiledName(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.FiledName = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetMarkMisType(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.MarkMisType = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetMarkMisValue(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.MarkMisValue = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetSupportedMisType(v []*string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.SupportedMisType = v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetUuid(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.Uuid = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource struct {
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
	// gmtModified
	FiledName *string `json:"FiledName,omitempty" xml:"FiledName,omitempty"`
	// The value of the field that can be used in the whitelist rule.
	//
	// example:
	//
	// contains
	MarkMisValue *string `json:"MarkMisValue,omitempty" xml:"MarkMisValue,omitempty"`
	// An array consisting of the operations that are supported by the method to add the alert event to the whitelist.
	SupportedMisType []*string `json:"SupportedMisType,omitempty" xml:"SupportedMisType,omitempty" type:"Repeated"`
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) GetFiledAliasName() *string {
	return s.FiledAliasName
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) GetFiledName() *string {
	return s.FiledName
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) GetMarkMisValue() *string {
	return s.MarkMisValue
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) GetSupportedMisType() []*string {
	return s.SupportedMisType
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) SetFiledAliasName(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	s.FiledAliasName = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) SetFiledName(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	s.FiledName = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) SetMarkMisValue(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	s.MarkMisValue = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) SetSupportedMisType(v []*string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	s.SupportedMisType = v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) Validate() error {
	return dara.Validate(s)
}
