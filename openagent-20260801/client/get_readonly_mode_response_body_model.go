// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReadonlyModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetReadonlyModeResponseBodyData) *GetReadonlyModeResponseBody
	GetData() *GetReadonlyModeResponseBodyData
	SetRequestId(v string) *GetReadonlyModeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetReadonlyModeResponseBody
	GetSuccess() *bool
}

type GetReadonlyModeResponseBody struct {
	// The returned data.
	Data *GetReadonlyModeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetReadonlyModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReadonlyModeResponseBody) GoString() string {
	return s.String()
}

func (s *GetReadonlyModeResponseBody) GetData() *GetReadonlyModeResponseBodyData {
	return s.Data
}

func (s *GetReadonlyModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetReadonlyModeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetReadonlyModeResponseBody) SetData(v *GetReadonlyModeResponseBodyData) *GetReadonlyModeResponseBody {
	s.Data = v
	return s
}

func (s *GetReadonlyModeResponseBody) SetRequestId(v string) *GetReadonlyModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetReadonlyModeResponseBody) SetSuccess(v bool) *GetReadonlyModeResponseBody {
	s.Success = &v
	return s
}

func (s *GetReadonlyModeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetReadonlyModeResponseBodyData struct {
	// The configuration name. The value is fixed to readonly.
	//
	// example:
	//
	// readonly
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// Indicates whether a configuration record exists.
	//
	// example:
	//
	// true
	Exists *bool `json:"exists,omitempty" xml:"exists,omitempty"`
	// The last update time in ISO 8601 format. The value is null if no configuration record exists.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	GmtUpdate *string `json:"gmtUpdate,omitempty" xml:"gmtUpdate,omitempty"`
	// The current effective read-only mode, returned as a string. A value of true indicates that the mode is enabled, and a value of false indicates that the mode is disabled. The value is true if not configured.
	//
	// example:
	//
	// true
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// The configuration value type. The value is fixed to BOOL.
	//
	// example:
	//
	// BOOL
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s GetReadonlyModeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetReadonlyModeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetReadonlyModeResponseBodyData) GetConfigName() *string {
	return s.ConfigName
}

func (s *GetReadonlyModeResponseBodyData) GetExists() *bool {
	return s.Exists
}

func (s *GetReadonlyModeResponseBodyData) GetGmtUpdate() *string {
	return s.GmtUpdate
}

func (s *GetReadonlyModeResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *GetReadonlyModeResponseBodyData) GetValueType() *string {
	return s.ValueType
}

func (s *GetReadonlyModeResponseBodyData) SetConfigName(v string) *GetReadonlyModeResponseBodyData {
	s.ConfigName = &v
	return s
}

func (s *GetReadonlyModeResponseBodyData) SetExists(v bool) *GetReadonlyModeResponseBodyData {
	s.Exists = &v
	return s
}

func (s *GetReadonlyModeResponseBodyData) SetGmtUpdate(v string) *GetReadonlyModeResponseBodyData {
	s.GmtUpdate = &v
	return s
}

func (s *GetReadonlyModeResponseBodyData) SetValue(v string) *GetReadonlyModeResponseBodyData {
	s.Value = &v
	return s
}

func (s *GetReadonlyModeResponseBodyData) SetValueType(v string) *GetReadonlyModeResponseBodyData {
	s.ValueType = &v
	return s
}

func (s *GetReadonlyModeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
