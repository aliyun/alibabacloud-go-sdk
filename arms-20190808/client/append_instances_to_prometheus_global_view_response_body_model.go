// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendInstancesToPrometheusGlobalViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AppendInstancesToPrometheusGlobalViewResponseBody
	GetCode() *int32
	SetData(v *AppendInstancesToPrometheusGlobalViewResponseBodyData) *AppendInstancesToPrometheusGlobalViewResponseBody
	GetData() *AppendInstancesToPrometheusGlobalViewResponseBodyData
	SetMessage(v string) *AppendInstancesToPrometheusGlobalViewResponseBody
	GetMessage() *string
	SetRequestId(v string) *AppendInstancesToPrometheusGlobalViewResponseBody
	GetRequestId() *string
}

type AppendInstancesToPrometheusGlobalViewResponseBody struct {
	// Status code. 200 means success, other status codes are exceptions.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the array object.
	Data *AppendInstancesToPrometheusGlobalViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Additional message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 27E653FA-5958-45BE-8AA9-14D884DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppendInstancesToPrometheusGlobalViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AppendInstancesToPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBody) GetData() *AppendInstancesToPrometheusGlobalViewResponseBodyData {
	return s.Data
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBody) SetCode(v int32) *AppendInstancesToPrometheusGlobalViewResponseBody {
	s.Code = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBody) SetData(v *AppendInstancesToPrometheusGlobalViewResponseBodyData) *AppendInstancesToPrometheusGlobalViewResponseBody {
	s.Data = v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBody) SetMessage(v string) *AppendInstancesToPrometheusGlobalViewResponseBody {
	s.Message = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBody) SetRequestId(v string) *AppendInstancesToPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBody) Validate() error {
	return dara.Validate(s)
}

type AppendInstancesToPrometheusGlobalViewResponseBodyData struct {
	// The Info-level information.
	//
	// example:
	//
	// {regionId: the region where the aggregation instance resides. globalViewClusterId: the ID of the aggregation instance. failedInstances: the ID of the object that failed to be added.}
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// The additional information.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- `true`: The call was successful.
	//
	// 	- `false`: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AppendInstancesToPrometheusGlobalViewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AppendInstancesToPrometheusGlobalViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBodyData) GetInfo() *string {
	return s.Info
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBodyData) SetInfo(v string) *AppendInstancesToPrometheusGlobalViewResponseBodyData {
	s.Info = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBodyData) SetMsg(v string) *AppendInstancesToPrometheusGlobalViewResponseBodyData {
	s.Msg = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBodyData) SetSuccess(v bool) *AppendInstancesToPrometheusGlobalViewResponseBodyData {
	s.Success = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBodyData) Validate() error {
	return dara.Validate(s)
}
