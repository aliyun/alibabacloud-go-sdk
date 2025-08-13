// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRuleReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetConfigRuleReportResponseBodyData) *GetConfigRuleReportResponseBody
	GetData() *GetConfigRuleReportResponseBodyData
	SetHttpStatusCode(v int32) *GetConfigRuleReportResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetConfigRuleReportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConfigRuleReportResponseBody
	GetSuccess() *bool
}

type GetConfigRuleReportResponseBody struct {
	// The basic information of the resource non-compliance report that is last generated.
	Data *GetConfigRuleReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A68BD5BC-5B12-5A9B-8AE9-77884886BE10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConfigRuleReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigRuleReportResponseBody) GetData() *GetConfigRuleReportResponseBodyData {
	return s.Data
}

func (s *GetConfigRuleReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetConfigRuleReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConfigRuleReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConfigRuleReportResponseBody) SetData(v *GetConfigRuleReportResponseBodyData) *GetConfigRuleReportResponseBody {
	s.Data = v
	return s
}

func (s *GetConfigRuleReportResponseBody) SetHttpStatusCode(v int32) *GetConfigRuleReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConfigRuleReportResponseBody) SetRequestId(v string) *GetConfigRuleReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigRuleReportResponseBody) SetSuccess(v bool) *GetConfigRuleReportResponseBody {
	s.Success = &v
	return s
}

func (s *GetConfigRuleReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetConfigRuleReportResponseBodyData struct {
	// The time when the report was generated. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1655089159000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the report.
	//
	// example:
	//
	// crp-ao0786618088006c****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the object.
	//
	// >  This parameter is returned if you set the `TargetType` and `TargetId` parameters in the current request to the same values as the parameters that are configured when you call the [GenerateConfigRuleReport](https://help.aliyun.com/document_detail/433313.html) operation to generate the report.
	//
	// example:
	//
	// 154950938137****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	//
	// 	- ROOT: the Root folder in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- FOLDER: a folder other than the Root folder in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- ACCOUNT: a member in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// >  This parameter is returned if you set the `TargetType` and `TargetId` parameters in the current request to the same values as the parameters that are configured when you call the [GenerateConfigRuleReport](https://help.aliyun.com/document_detail/433313.html) operation to generate the report.
	//
	// example:
	//
	// ACCOUNT
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetConfigRuleReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConfigRuleReportResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetConfigRuleReportResponseBodyData) GetReportId() *string {
	return s.ReportId
}

func (s *GetConfigRuleReportResponseBodyData) GetTargetId() *string {
	return s.TargetId
}

func (s *GetConfigRuleReportResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *GetConfigRuleReportResponseBodyData) SetCreatedTime(v int64) *GetConfigRuleReportResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetConfigRuleReportResponseBodyData) SetReportId(v string) *GetConfigRuleReportResponseBodyData {
	s.ReportId = &v
	return s
}

func (s *GetConfigRuleReportResponseBodyData) SetTargetId(v string) *GetConfigRuleReportResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *GetConfigRuleReportResponseBodyData) SetTargetType(v string) *GetConfigRuleReportResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *GetConfigRuleReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
