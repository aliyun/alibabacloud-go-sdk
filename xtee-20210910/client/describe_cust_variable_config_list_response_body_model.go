// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustVariableConfigListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCustVariableConfigListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeCustVariableConfigListResponseBodyResultObject) *DescribeCustVariableConfigListResponseBody
	GetResultObject() []*DescribeCustVariableConfigListResponseBodyResultObject
}

type DescribeCustVariableConfigListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeCustVariableConfigListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeCustVariableConfigListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustVariableConfigListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustVariableConfigListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustVariableConfigListResponseBody) GetResultObject() []*DescribeCustVariableConfigListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeCustVariableConfigListResponseBody) SetRequestId(v string) *DescribeCustVariableConfigListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustVariableConfigListResponseBody) SetResultObject(v []*DescribeCustVariableConfigListResponseBodyResultObject) *DescribeCustVariableConfigListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeCustVariableConfigListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCustVariableConfigListResponseBodyResultObject struct {
	// Configuration key
	//
	// example:
	//
	// COUNT
	ConfigKey *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
	// Configuration value
	//
	// example:
	//
	// 累计
	ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty"`
}

func (s DescribeCustVariableConfigListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustVariableConfigListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeCustVariableConfigListResponseBodyResultObject) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DescribeCustVariableConfigListResponseBodyResultObject) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *DescribeCustVariableConfigListResponseBodyResultObject) SetConfigKey(v string) *DescribeCustVariableConfigListResponseBodyResultObject {
	s.ConfigKey = &v
	return s
}

func (s *DescribeCustVariableConfigListResponseBodyResultObject) SetConfigValue(v string) *DescribeCustVariableConfigListResponseBodyResultObject {
	s.ConfigValue = &v
	return s
}

func (s *DescribeCustVariableConfigListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
