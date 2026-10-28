// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReadonlyModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateReadonlyModeResponseBodyData) *UpdateReadonlyModeResponseBody
	GetData() *UpdateReadonlyModeResponseBodyData
	SetRequestId(v string) *UpdateReadonlyModeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateReadonlyModeResponseBody
	GetSuccess() *bool
}

type UpdateReadonlyModeResponseBody struct {
	// The currently active configuration information.
	Data *UpdateReadonlyModeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateReadonlyModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateReadonlyModeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateReadonlyModeResponseBody) GetData() *UpdateReadonlyModeResponseBodyData {
	return s.Data
}

func (s *UpdateReadonlyModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateReadonlyModeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateReadonlyModeResponseBody) SetData(v *UpdateReadonlyModeResponseBodyData) *UpdateReadonlyModeResponseBody {
	s.Data = v
	return s
}

func (s *UpdateReadonlyModeResponseBody) SetRequestId(v string) *UpdateReadonlyModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateReadonlyModeResponseBody) SetSuccess(v bool) *UpdateReadonlyModeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateReadonlyModeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateReadonlyModeResponseBodyData struct {
	// The name of the configuration. The value is fixed to readonly.
	//
	// example:
	//
	// readonly
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// Indicates whether a configuration record already exists.
	//
	// example:
	//
	// true
	Exists *bool `json:"exists,omitempty" xml:"exists,omitempty"`
	// The last update time in ISO 8601 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	GmtUpdate *string `json:"gmtUpdate,omitempty" xml:"gmtUpdate,omitempty"`
	// The currently active configuration value, returned as the string true or false.
	//
	// example:
	//
	// true
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// The type of the configuration value. The value is BOOL for this API operation.
	//
	// example:
	//
	// BOOL
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s UpdateReadonlyModeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateReadonlyModeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateReadonlyModeResponseBodyData) GetConfigName() *string {
	return s.ConfigName
}

func (s *UpdateReadonlyModeResponseBodyData) GetExists() *bool {
	return s.Exists
}

func (s *UpdateReadonlyModeResponseBodyData) GetGmtUpdate() *string {
	return s.GmtUpdate
}

func (s *UpdateReadonlyModeResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *UpdateReadonlyModeResponseBodyData) GetValueType() *string {
	return s.ValueType
}

func (s *UpdateReadonlyModeResponseBodyData) SetConfigName(v string) *UpdateReadonlyModeResponseBodyData {
	s.ConfigName = &v
	return s
}

func (s *UpdateReadonlyModeResponseBodyData) SetExists(v bool) *UpdateReadonlyModeResponseBodyData {
	s.Exists = &v
	return s
}

func (s *UpdateReadonlyModeResponseBodyData) SetGmtUpdate(v string) *UpdateReadonlyModeResponseBodyData {
	s.GmtUpdate = &v
	return s
}

func (s *UpdateReadonlyModeResponseBodyData) SetValue(v string) *UpdateReadonlyModeResponseBodyData {
	s.Value = &v
	return s
}

func (s *UpdateReadonlyModeResponseBodyData) SetValueType(v string) *UpdateReadonlyModeResponseBodyData {
	s.ValueType = &v
	return s
}

func (s *UpdateReadonlyModeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
