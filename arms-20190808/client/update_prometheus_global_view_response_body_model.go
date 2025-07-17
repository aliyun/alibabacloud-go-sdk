// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusGlobalViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdatePrometheusGlobalViewResponseBody
	GetCode() *int32
	SetData(v *UpdatePrometheusGlobalViewResponseBodyData) *UpdatePrometheusGlobalViewResponseBody
	GetData() *UpdatePrometheusGlobalViewResponseBodyData
	SetMessage(v string) *UpdatePrometheusGlobalViewResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePrometheusGlobalViewResponseBody
	GetRequestId() *string
}

type UpdatePrometheusGlobalViewResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *UpdatePrometheusGlobalViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// E9C9DA3D-10FE-472E-9EEF-2D0A3E41****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrometheusGlobalViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusGlobalViewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdatePrometheusGlobalViewResponseBody) GetData() *UpdatePrometheusGlobalViewResponseBodyData {
	return s.Data
}

func (s *UpdatePrometheusGlobalViewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePrometheusGlobalViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrometheusGlobalViewResponseBody) SetCode(v int32) *UpdatePrometheusGlobalViewResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePrometheusGlobalViewResponseBody) SetData(v *UpdatePrometheusGlobalViewResponseBodyData) *UpdatePrometheusGlobalViewResponseBody {
	s.Data = v
	return s
}

func (s *UpdatePrometheusGlobalViewResponseBody) SetMessage(v string) *UpdatePrometheusGlobalViewResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusGlobalViewResponseBody) SetRequestId(v string) *UpdatePrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrometheusGlobalViewResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdatePrometheusGlobalViewResponseBodyData struct {
	// The data sources that failed to be updated.
	FailedInstances []*UpdatePrometheusGlobalViewResponseBodyDataFailedInstances `json:"FailedInstances,omitempty" xml:"FailedInstances,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePrometheusGlobalViewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusGlobalViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusGlobalViewResponseBodyData) GetFailedInstances() []*UpdatePrometheusGlobalViewResponseBodyDataFailedInstances {
	return s.FailedInstances
}

func (s *UpdatePrometheusGlobalViewResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePrometheusGlobalViewResponseBodyData) SetFailedInstances(v []*UpdatePrometheusGlobalViewResponseBodyDataFailedInstances) *UpdatePrometheusGlobalViewResponseBodyData {
	s.FailedInstances = v
	return s
}

func (s *UpdatePrometheusGlobalViewResponseBodyData) SetSuccess(v bool) *UpdatePrometheusGlobalViewResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdatePrometheusGlobalViewResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type UpdatePrometheusGlobalViewResponseBodyDataFailedInstances struct {
	// The ID of the Prometheus instance.
	//
	// example:
	//
	// cdb65ed2d527345*********
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// sourcename-test
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	// The type of the data source. AlibabaPrometheus MetricStore CustomPrometheus
	//
	// example:
	//
	// AlibabaPrometheus
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 23784673825*******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdatePrometheusGlobalViewResponseBodyDataFailedInstances) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusGlobalViewResponseBodyDataFailedInstances) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances) GetSourceName() *string {
	return s.SourceName
}

func (s *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances) GetUserId() *string {
	return s.UserId
}

func (s *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances) SetClusterId(v string) *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances) SetSourceName(v string) *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances {
	s.SourceName = &v
	return s
}

func (s *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances) SetSourceType(v string) *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances {
	s.SourceType = &v
	return s
}

func (s *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances) SetUserId(v string) *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances {
	s.UserId = &v
	return s
}

func (s *UpdatePrometheusGlobalViewResponseBodyDataFailedInstances) Validate() error {
	return dara.Validate(s)
}
