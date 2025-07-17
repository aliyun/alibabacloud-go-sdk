// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusGlobalViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetPrometheusGlobalViewResponseBody
	GetCode() *int32
	SetData(v string) *GetPrometheusGlobalViewResponseBody
	GetData() *string
	SetMessage(v string) *GetPrometheusGlobalViewResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPrometheusGlobalViewResponseBody
	GetRequestId() *string
}

type GetPrometheusGlobalViewResponseBody struct {
	// Status code. 200 is success, other status codes are exceptions.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct returned.
	//
	// example:
	//
	// { "clusterId":"The ID of the global aggregation instance.", "groupName":"The name of the global aggregation instance.", "dataSources":[ { "sourceName":"The name of the data source.- ArmsPrometheus No.1", "sourceType":"AlibabaPrometheus", "userId":"UserID", "clusterId":"ClusterId" }, // more datasources ] }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Returns a hint message for the result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 743AD493-D006-53BD-AAEC-DDCE7FB68EA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPrometheusGlobalViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrometheusGlobalViewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetPrometheusGlobalViewResponseBody) GetData() *string {
	return s.Data
}

func (s *GetPrometheusGlobalViewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPrometheusGlobalViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPrometheusGlobalViewResponseBody) SetCode(v int32) *GetPrometheusGlobalViewResponseBody {
	s.Code = &v
	return s
}

func (s *GetPrometheusGlobalViewResponseBody) SetData(v string) *GetPrometheusGlobalViewResponseBody {
	s.Data = &v
	return s
}

func (s *GetPrometheusGlobalViewResponseBody) SetMessage(v string) *GetPrometheusGlobalViewResponseBody {
	s.Message = &v
	return s
}

func (s *GetPrometheusGlobalViewResponseBody) SetRequestId(v string) *GetPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrometheusGlobalViewResponseBody) Validate() error {
	return dara.Validate(s)
}
