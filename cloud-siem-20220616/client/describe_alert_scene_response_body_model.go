// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAlertSceneResponseBody
	GetCode() *int32
	SetData(v []*DescribeAlertSceneResponseBodyData) *DescribeAlertSceneResponseBody
	GetData() []*DescribeAlertSceneResponseBodyData
	SetMessage(v string) *DescribeAlertSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertSceneResponseBody
	GetSuccess() *bool
}

type DescribeAlertSceneResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeAlertSceneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAlertSceneResponseBody) GetData() []*DescribeAlertSceneResponseBodyData {
	return s.Data
}

func (s *DescribeAlertSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertSceneResponseBody) SetCode(v int32) *DescribeAlertSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertSceneResponseBody) SetData(v []*DescribeAlertSceneResponseBodyData) *DescribeAlertSceneResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertSceneResponseBody) SetMessage(v string) *DescribeAlertSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertSceneResponseBody) SetRequestId(v string) *DescribeAlertSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertSceneResponseBody) SetSuccess(v bool) *DescribeAlertSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertSceneResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertSceneResponseBodyData struct {
	// The name of the alert. The value varies based on the display language (Chinese or English) of the Security Center console.
	//
	// example:
	//
	// login_common_ip
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The ID of the alert name.
	//
	// example:
	//
	// login_common_ip
	AlertNameId *string `json:"AlertNameId,omitempty" xml:"AlertNameId,omitempty"`
	// The title of the alert notification. The value varies based on the display language (Chinese or English) of the Security Center console.
	//
	// example:
	//
	// unusual login-login_common_ip
	AlertTile *string `json:"AlertTile,omitempty" xml:"AlertTile,omitempty"`
	// The ID of the alert title.
	//
	// example:
	//
	// unusual login-login_common_ip
	AlertTileId *string `json:"AlertTileId,omitempty" xml:"AlertTileId,omitempty"`
	// The type of the alert. The value varies based on the display language (Chinese or English) of the Security Center console.
	//
	// example:
	//
	// unusual login
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The ID of the alert type.
	//
	// example:
	//
	// unusual login
	AlertTypeId *string `json:"AlertTypeId,omitempty" xml:"AlertTypeId,omitempty"`
	// The information about the entities for which you need to add the alert to the whitelist.
	//
	// example:
	//
	// [{"Type": "host_uuid","Value": "441862da-a539-4cc0-a00d-473955826881","Values": ["441862da-a539-4cc0-a00d-473955826881"],"Name": "${aliyun.siem.entity.host_uuid}"}]
	Targets []*DescribeAlertSceneResponseBodyDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s DescribeAlertSceneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSceneResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneResponseBodyData) GetAlertName() *string {
	return s.AlertName
}

func (s *DescribeAlertSceneResponseBodyData) GetAlertNameId() *string {
	return s.AlertNameId
}

func (s *DescribeAlertSceneResponseBodyData) GetAlertTile() *string {
	return s.AlertTile
}

func (s *DescribeAlertSceneResponseBodyData) GetAlertTileId() *string {
	return s.AlertTileId
}

func (s *DescribeAlertSceneResponseBodyData) GetAlertType() *string {
	return s.AlertType
}

func (s *DescribeAlertSceneResponseBodyData) GetAlertTypeId() *string {
	return s.AlertTypeId
}

func (s *DescribeAlertSceneResponseBodyData) GetTargets() []*DescribeAlertSceneResponseBodyDataTargets {
	return s.Targets
}

func (s *DescribeAlertSceneResponseBodyData) SetAlertName(v string) *DescribeAlertSceneResponseBodyData {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) SetAlertNameId(v string) *DescribeAlertSceneResponseBodyData {
	s.AlertNameId = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) SetAlertTile(v string) *DescribeAlertSceneResponseBodyData {
	s.AlertTile = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) SetAlertTileId(v string) *DescribeAlertSceneResponseBodyData {
	s.AlertTileId = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) SetAlertType(v string) *DescribeAlertSceneResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) SetAlertTypeId(v string) *DescribeAlertSceneResponseBodyData {
	s.AlertTypeId = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) SetTargets(v []*DescribeAlertSceneResponseBodyDataTargets) *DescribeAlertSceneResponseBodyData {
	s.Targets = v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertSceneResponseBodyDataTargets struct {
	// The display name of the attribute for the entity.
	//
	// example:
	//
	// HOST UUID
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The attribute of the entity.
	//
	// example:
	//
	// host_uuid
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The right operand that is displayed by default in the whitelist rule.
	//
	// example:
	//
	// 441862da-a539-4cc0-a00d-47395582****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The right operands supported by the whitelist rule.
	//
	// example:
	//
	// ["441862da-a539-4cc0-a00d-473955826881"]
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeAlertSceneResponseBodyDataTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSceneResponseBodyDataTargets) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneResponseBodyDataTargets) GetName() *string {
	return s.Name
}

func (s *DescribeAlertSceneResponseBodyDataTargets) GetType() *string {
	return s.Type
}

func (s *DescribeAlertSceneResponseBodyDataTargets) GetValue() *string {
	return s.Value
}

func (s *DescribeAlertSceneResponseBodyDataTargets) GetValues() []*string {
	return s.Values
}

func (s *DescribeAlertSceneResponseBodyDataTargets) SetName(v string) *DescribeAlertSceneResponseBodyDataTargets {
	s.Name = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyDataTargets) SetType(v string) *DescribeAlertSceneResponseBodyDataTargets {
	s.Type = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyDataTargets) SetValue(v string) *DescribeAlertSceneResponseBodyDataTargets {
	s.Value = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyDataTargets) SetValues(v []*string) *DescribeAlertSceneResponseBodyDataTargets {
	s.Values = v
	return s
}

func (s *DescribeAlertSceneResponseBodyDataTargets) Validate() error {
	return dara.Validate(s)
}
