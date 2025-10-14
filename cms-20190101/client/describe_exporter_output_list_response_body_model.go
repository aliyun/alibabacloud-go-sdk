// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExporterOutputListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeExporterOutputListResponseBody
	GetCode() *string
	SetDatapoints(v *DescribeExporterOutputListResponseBodyDatapoints) *DescribeExporterOutputListResponseBody
	GetDatapoints() *DescribeExporterOutputListResponseBodyDatapoints
	SetMessage(v string) *DescribeExporterOutputListResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeExporterOutputListResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeExporterOutputListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeExporterOutputListResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeExporterOutputListResponseBody
	GetTotal() *int32
}

type DescribeExporterOutputListResponseBody struct {
	// The HTTP status code.
	//
	// > The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration sets for exporting monitoring data.
	Datapoints *DescribeExporterOutputListResponseBodyDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// sucess
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0E657631-CD6C-4C24-9637-98D000B9272C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The total number of returned entries.
	//
	// example:
	//
	// 25
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeExporterOutputListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExporterOutputListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExporterOutputListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeExporterOutputListResponseBody) GetDatapoints() *DescribeExporterOutputListResponseBodyDatapoints {
	return s.Datapoints
}

func (s *DescribeExporterOutputListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeExporterOutputListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeExporterOutputListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExporterOutputListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeExporterOutputListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeExporterOutputListResponseBody) SetCode(v string) *DescribeExporterOutputListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExporterOutputListResponseBody) SetDatapoints(v *DescribeExporterOutputListResponseBodyDatapoints) *DescribeExporterOutputListResponseBody {
	s.Datapoints = v
	return s
}

func (s *DescribeExporterOutputListResponseBody) SetMessage(v string) *DescribeExporterOutputListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExporterOutputListResponseBody) SetPageNumber(v int32) *DescribeExporterOutputListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeExporterOutputListResponseBody) SetRequestId(v string) *DescribeExporterOutputListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExporterOutputListResponseBody) SetSuccess(v bool) *DescribeExporterOutputListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExporterOutputListResponseBody) SetTotal(v int32) *DescribeExporterOutputListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeExporterOutputListResponseBody) Validate() error {
	if s.Datapoints != nil {
		if err := s.Datapoints.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeExporterOutputListResponseBodyDatapoints struct {
	Datapoint []*DescribeExporterOutputListResponseBodyDatapointsDatapoint `json:"Datapoint,omitempty" xml:"Datapoint,omitempty" type:"Repeated"`
}

func (s DescribeExporterOutputListResponseBodyDatapoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeExporterOutputListResponseBodyDatapoints) GoString() string {
	return s.String()
}

func (s *DescribeExporterOutputListResponseBodyDatapoints) GetDatapoint() []*DescribeExporterOutputListResponseBodyDatapointsDatapoint {
	return s.Datapoint
}

func (s *DescribeExporterOutputListResponseBodyDatapoints) SetDatapoint(v []*DescribeExporterOutputListResponseBodyDatapointsDatapoint) *DescribeExporterOutputListResponseBodyDatapoints {
	s.Datapoint = v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapoints) Validate() error {
	if s.Datapoint != nil {
		for _, item := range s.Datapoint {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExporterOutputListResponseBodyDatapointsDatapoint struct {
	// The JSON object that contains the details about the destination to which the monitoring data is exported.
	ConfigJson *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson `json:"ConfigJson,omitempty" xml:"ConfigJson,omitempty" type:"Struct"`
	// The time when the configuration set was created. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1584016495498
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the configuration set.
	//
	// example:
	//
	// exporterOut
	DestName *string `json:"DestName,omitempty" xml:"DestName,omitempty"`
	// The service to which the monitoring data is exported.
	//
	// > Only Log Service is supported. More services will be supported in the future.
	//
	// example:
	//
	// SLS
	DestType *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
}

func (s DescribeExporterOutputListResponseBodyDatapointsDatapoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeExporterOutputListResponseBodyDatapointsDatapoint) GoString() string {
	return s.String()
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) GetConfigJson() *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson {
	return s.ConfigJson
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) GetDestName() *string {
	return s.DestName
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) GetDestType() *string {
	return s.DestType
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) SetConfigJson(v *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) *DescribeExporterOutputListResponseBodyDatapointsDatapoint {
	s.ConfigJson = v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) SetCreateTime(v int64) *DescribeExporterOutputListResponseBodyDatapointsDatapoint {
	s.CreateTime = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) SetDestName(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapoint {
	s.DestName = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) SetDestType(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapoint {
	s.DestType = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) Validate() error {
	if s.ConfigJson != nil {
		if err := s.ConfigJson.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson struct {
	// The AccessKey ID.
	//
	// example:
	//
	// LTAIpY33********
	Ak *string `json:"ak,omitempty" xml:"ak,omitempty"`
	// The Log Service endpoint to which the monitoring data is exported.
	//
	// example:
	//
	// http://cn-qingdao-share.log.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The Logstore.
	//
	// example:
	//
	// monitorlogstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The Log Service project to which the monitoring data is exported.
	//
	// example:
	//
	// exporter
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) String() string {
	return dara.Prettify(s)
}

func (s DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) GoString() string {
	return s.String()
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) GetAk() *string {
	return s.Ak
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) GetLogstore() *string {
	return s.Logstore
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) GetProject() *string {
	return s.Project
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) SetAk(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson {
	s.Ak = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) SetEndpoint(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson {
	s.Endpoint = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) SetLogstore(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson {
	s.Logstore = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) SetProject(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson {
	s.Project = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) Validate() error {
	return dara.Validate(s)
}
