// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertSceneByEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAlertSceneByEventResponseBody
	GetCode() *int32
	SetData(v []*DescribeAlertSceneByEventResponseBodyData) *DescribeAlertSceneByEventResponseBody
	GetData() []*DescribeAlertSceneByEventResponseBodyData
	SetMessage(v string) *DescribeAlertSceneByEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertSceneByEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertSceneByEventResponseBody
	GetSuccess() *bool
}

type DescribeAlertSceneByEventResponseBody struct {
	// The HTTP status code.
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
	Data []*DescribeAlertSceneByEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeAlertSceneByEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSceneByEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAlertSceneByEventResponseBody) GetData() []*DescribeAlertSceneByEventResponseBodyData {
	return s.Data
}

func (s *DescribeAlertSceneByEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertSceneByEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertSceneByEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertSceneByEventResponseBody) SetCode(v int32) *DescribeAlertSceneByEventResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) SetData(v []*DescribeAlertSceneByEventResponseBodyData) *DescribeAlertSceneByEventResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) SetMessage(v string) *DescribeAlertSceneByEventResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) SetRequestId(v string) *DescribeAlertSceneByEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) SetSuccess(v bool) *DescribeAlertSceneByEventResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) Validate() error {
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

type DescribeAlertSceneByEventResponseBodyData struct {
	// The alert name. The display name of the alert name varies based on the language of the system, such as Chinese and English.
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
	// The alert title. The display name of the alert title varies based on the language of the system, such as Chinese and English.
	//
	// example:
	//
	// Unusual Logon-login_common_ip
	AlertTile *string `json:"AlertTile,omitempty" xml:"AlertTile,omitempty"`
	// The ID of the alert title.
	//
	// example:
	//
	// Unusual Logon-login_common_ip
	AlertTileId *string `json:"AlertTileId,omitempty" xml:"AlertTileId,omitempty"`
	// The alert type. The display name of the alert type varies based on the language of the system, such as Chinese and English.
	//
	// example:
	//
	// Unusual Logon
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The ID of the alert type.
	//
	// example:
	//
	// Unusual Logon
	AlertTypeId *string `json:"AlertTypeId,omitempty" xml:"AlertTypeId,omitempty"`
	// The objects that can be added to the whitelist.
	//
	// example:
	//
	// [{"Type": "host_uuid","Value": "441862da-a539-4cc0-a00d-473955826881","Values": ["441862da-a539-4cc0-a00d-473955826881"],"Name": "${aliyun.siem.entity.host_uuid}"}]
	Targets []*DescribeAlertSceneByEventResponseBodyDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s DescribeAlertSceneByEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSceneByEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventResponseBodyData) GetAlertName() *string {
	return s.AlertName
}

func (s *DescribeAlertSceneByEventResponseBodyData) GetAlertNameId() *string {
	return s.AlertNameId
}

func (s *DescribeAlertSceneByEventResponseBodyData) GetAlertTile() *string {
	return s.AlertTile
}

func (s *DescribeAlertSceneByEventResponseBodyData) GetAlertTileId() *string {
	return s.AlertTileId
}

func (s *DescribeAlertSceneByEventResponseBodyData) GetAlertType() *string {
	return s.AlertType
}

func (s *DescribeAlertSceneByEventResponseBodyData) GetAlertTypeId() *string {
	return s.AlertTypeId
}

func (s *DescribeAlertSceneByEventResponseBodyData) GetTargets() []*DescribeAlertSceneByEventResponseBodyDataTargets {
	return s.Targets
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertName(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertNameId(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertNameId = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertTile(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertTile = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertTileId(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertTileId = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertType(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertTypeId(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertTypeId = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetTargets(v []*DescribeAlertSceneByEventResponseBodyDataTargets) *DescribeAlertSceneByEventResponseBodyData {
	s.Targets = v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) Validate() error {
	if s.Targets != nil {
		for _, item := range s.Targets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertSceneByEventResponseBodyDataTargets struct {
	// The display name of the entity attribute field that can be added to the whitelist.
	//
	// example:
	//
	// host uuid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The entity attribute field that can be added to the whitelist.
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
	// The supported right operands of the whitelist rule.
	//
	// example:
	//
	// ["441862da-a539-4cc0-a00d-473955826881"]
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeAlertSceneByEventResponseBodyDataTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSceneByEventResponseBodyDataTargets) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) GetName() *string {
	return s.Name
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) GetType() *string {
	return s.Type
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) GetValue() *string {
	return s.Value
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) GetValues() []*string {
	return s.Values
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) SetName(v string) *DescribeAlertSceneByEventResponseBodyDataTargets {
	s.Name = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) SetType(v string) *DescribeAlertSceneByEventResponseBodyDataTargets {
	s.Type = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) SetValue(v string) *DescribeAlertSceneByEventResponseBodyDataTargets {
	s.Value = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) SetValues(v []*string) *DescribeAlertSceneByEventResponseBodyDataTargets {
	s.Values = v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) Validate() error {
	return dara.Validate(s)
}
