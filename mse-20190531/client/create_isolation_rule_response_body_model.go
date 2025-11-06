// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIsolationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateIsolationRuleResponseBody
	GetCode() *string
	SetData(v *CreateIsolationRuleResponseBodyData) *CreateIsolationRuleResponseBody
	GetData() *CreateIsolationRuleResponseBodyData
	SetMessage(v string) *CreateIsolationRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateIsolationRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateIsolationRuleResponseBody
	GetSuccess() *bool
}

type CreateIsolationRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateIsolationRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4E9FDCFE-0738-493B-B801-82BDFBCB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateIsolationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIsolationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIsolationRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateIsolationRuleResponseBody) GetData() *CreateIsolationRuleResponseBodyData {
	return s.Data
}

func (s *CreateIsolationRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateIsolationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIsolationRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateIsolationRuleResponseBody) SetCode(v string) *CreateIsolationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIsolationRuleResponseBody) SetData(v *CreateIsolationRuleResponseBodyData) *CreateIsolationRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateIsolationRuleResponseBody) SetMessage(v string) *CreateIsolationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIsolationRuleResponseBody) SetRequestId(v string) *CreateIsolationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIsolationRuleResponseBody) SetSuccess(v bool) *CreateIsolationRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIsolationRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIsolationRuleResponseBodyData struct {
	// example:
	//
	// hpn9ac29kz@e31a4b871******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 1
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LimitApp *string `json:"LimitApp,omitempty" xml:"LimitApp,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// /a
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// 3
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateIsolationRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateIsolationRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIsolationRuleResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *CreateIsolationRuleResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *CreateIsolationRuleResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *CreateIsolationRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CreateIsolationRuleResponseBodyData) GetLimitApp() *string {
	return s.LimitApp
}

func (s *CreateIsolationRuleResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateIsolationRuleResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIsolationRuleResponseBodyData) GetResource() *string {
	return s.Resource
}

func (s *CreateIsolationRuleResponseBodyData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CreateIsolationRuleResponseBodyData) SetAppId(v string) *CreateIsolationRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetAppName(v string) *CreateIsolationRuleResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetEnable(v bool) *CreateIsolationRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetId(v int64) *CreateIsolationRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetLimitApp(v string) *CreateIsolationRuleResponseBodyData {
	s.LimitApp = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetNamespace(v string) *CreateIsolationRuleResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetRegionId(v string) *CreateIsolationRuleResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetResource(v string) *CreateIsolationRuleResponseBodyData {
	s.Resource = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) SetThreshold(v float32) *CreateIsolationRuleResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *CreateIsolationRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
