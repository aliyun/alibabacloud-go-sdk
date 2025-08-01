// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySlbSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QuerySlbSpecResponseBody
	GetCode() *int32
	SetData(v []*QuerySlbSpecResponseBodyData) *QuerySlbSpecResponseBody
	GetData() []*QuerySlbSpecResponseBodyData
	SetHttpStatusCode(v int32) *QuerySlbSpecResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QuerySlbSpecResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySlbSpecResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySlbSpecResponseBody
	GetSuccess() *bool
}

type QuerySlbSpecResponseBody struct {
	// The return value.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data entries returned.
	Data []*QuerySlbSpecResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned. If the request is successful, a success message is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySlbSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySlbSpecResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySlbSpecResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QuerySlbSpecResponseBody) GetData() []*QuerySlbSpecResponseBodyData {
	return s.Data
}

func (s *QuerySlbSpecResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QuerySlbSpecResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySlbSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySlbSpecResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySlbSpecResponseBody) SetCode(v int32) *QuerySlbSpecResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySlbSpecResponseBody) SetData(v []*QuerySlbSpecResponseBodyData) *QuerySlbSpecResponseBody {
	s.Data = v
	return s
}

func (s *QuerySlbSpecResponseBody) SetHttpStatusCode(v int32) *QuerySlbSpecResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySlbSpecResponseBody) SetMessage(v string) *QuerySlbSpecResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySlbSpecResponseBody) SetRequestId(v string) *QuerySlbSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySlbSpecResponseBody) SetSuccess(v bool) *QuerySlbSpecResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySlbSpecResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySlbSpecResponseBodyData struct {
	// The ID of the returned data.
	//
	// example:
	//
	// 2
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The maximum number of connections.
	//
	// example:
	//
	// 50,000
	MaxConnection *string `json:"MaxConnection,omitempty" xml:"MaxConnection,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// Standard I
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of connections per second.
	//
	// example:
	//
	// 50,000
	NewConnectionPerSecond *string `json:"NewConnectionPerSecond,omitempty" xml:"NewConnectionPerSecond,omitempty"`
	// The number of queries per second (QPS).
	//
	// example:
	//
	// 50,000
	Qps *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The specification of the instance.
	//
	// example:
	//
	// slb.s2.small
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s QuerySlbSpecResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySlbSpecResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySlbSpecResponseBodyData) GetId() *int32 {
	return s.Id
}

func (s *QuerySlbSpecResponseBodyData) GetMaxConnection() *string {
	return s.MaxConnection
}

func (s *QuerySlbSpecResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QuerySlbSpecResponseBodyData) GetNewConnectionPerSecond() *string {
	return s.NewConnectionPerSecond
}

func (s *QuerySlbSpecResponseBodyData) GetQps() *string {
	return s.Qps
}

func (s *QuerySlbSpecResponseBodyData) GetSpec() *string {
	return s.Spec
}

func (s *QuerySlbSpecResponseBodyData) SetId(v int32) *QuerySlbSpecResponseBodyData {
	s.Id = &v
	return s
}

func (s *QuerySlbSpecResponseBodyData) SetMaxConnection(v string) *QuerySlbSpecResponseBodyData {
	s.MaxConnection = &v
	return s
}

func (s *QuerySlbSpecResponseBodyData) SetName(v string) *QuerySlbSpecResponseBodyData {
	s.Name = &v
	return s
}

func (s *QuerySlbSpecResponseBodyData) SetNewConnectionPerSecond(v string) *QuerySlbSpecResponseBodyData {
	s.NewConnectionPerSecond = &v
	return s
}

func (s *QuerySlbSpecResponseBodyData) SetQps(v string) *QuerySlbSpecResponseBodyData {
	s.Qps = &v
	return s
}

func (s *QuerySlbSpecResponseBodyData) SetSpec(v string) *QuerySlbSpecResponseBodyData {
	s.Spec = &v
	return s
}

func (s *QuerySlbSpecResponseBodyData) Validate() error {
	return dara.Validate(s)
}
