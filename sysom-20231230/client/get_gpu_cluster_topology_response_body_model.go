// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGpuClusterTopologyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGpuClusterTopologyResponseBody
	GetCode() *string
	SetData(v *GetGpuClusterTopologyResponseBodyData) *GetGpuClusterTopologyResponseBody
	GetData() *GetGpuClusterTopologyResponseBodyData
	SetMessage(v string) *GetGpuClusterTopologyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGpuClusterTopologyResponseBody
	GetRequestId() *string
}

type GetGpuClusterTopologyResponseBody struct {
	// The response code. Valid values: Success, Sysom.ServerError, and Sysom.InvalidParameter.
	//
	// example:
	//
	// Fail/Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned by the operation.
	Data *GetGpuClusterTopologyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message returned when the code is not Success.
	//
	// example:
	//
	// some error msg if not success，“” if success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// return real request ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGpuClusterTopologyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGpuClusterTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *GetGpuClusterTopologyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGpuClusterTopologyResponseBody) GetData() *GetGpuClusterTopologyResponseBodyData {
	return s.Data
}

func (s *GetGpuClusterTopologyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGpuClusterTopologyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGpuClusterTopologyResponseBody) SetCode(v string) *GetGpuClusterTopologyResponseBody {
	s.Code = &v
	return s
}

func (s *GetGpuClusterTopologyResponseBody) SetData(v *GetGpuClusterTopologyResponseBodyData) *GetGpuClusterTopologyResponseBody {
	s.Data = v
	return s
}

func (s *GetGpuClusterTopologyResponseBody) SetMessage(v string) *GetGpuClusterTopologyResponseBody {
	s.Message = &v
	return s
}

func (s *GetGpuClusterTopologyResponseBody) SetRequestId(v string) *GetGpuClusterTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGpuClusterTopologyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGpuClusterTopologyResponseBodyData struct {
	// The cluster topology data structure.
	Data *GetGpuClusterTopologyResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The total number of GPUs returned.
	//
	// example:
	//
	// 36
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetGpuClusterTopologyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGpuClusterTopologyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGpuClusterTopologyResponseBodyData) GetData() *GetGpuClusterTopologyResponseBodyDataData {
	return s.Data
}

func (s *GetGpuClusterTopologyResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *GetGpuClusterTopologyResponseBodyData) SetData(v *GetGpuClusterTopologyResponseBodyDataData) *GetGpuClusterTopologyResponseBodyData {
	s.Data = v
	return s
}

func (s *GetGpuClusterTopologyResponseBodyData) SetTotal(v int32) *GetGpuClusterTopologyResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetGpuClusterTopologyResponseBodyData) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGpuClusterTopologyResponseBodyDataData struct {
	// The edge data in the topology structure.
	//
	// example:
	//
	// "edges": [
	//
	//         {
	//
	//           "type": "GPU",
	//
	//           "src_id": "00534a88",
	//
	//           "dst_id": "260525e9",
	//
	//           "values": {
	//
	//             "Exception": [
	//
	//               {
	//
	//                 "key": "NCCL HANG",
	//
	//                 "value": "error0"
	//
	//               }
	//
	//             ],
	//
	//             "Metric": [
	//
	//               {
	//
	//                 "key": "a",
	//
	//                 "value": "372201199012283428"
	//
	//               }
	//
	//             ],
	//
	//             "Basic": [
	//
	//               {
	//
	//                 "key": "NCCL Version",
	//
	//                 "value": "v2.22.3"
	//
	//               }
	//
	//             ]
	//
	//           },
	//
	//           "style": {
	//
	//             "lineWidth": "0.01",
	//
	//             "opacity": "1",
	//
	//             "selectedStyle": "\\"{\\\\\\"opacity\\\\\\": 1, \\\\\\"lineWidth\\\\\\": 1}\\"",
	//
	//             "stroke": "#e94242\\n"
	//
	//           }
	//
	//         }
	//
	//       ]
	Edges interface{} `json:"edges,omitempty" xml:"edges,omitempty"`
	// The node data in the topology structure.
	//
	// example:
	//
	// "nodes": [
	//
	//         {
	//
	//           "id": "i-xxxxxx",
	//
	//           "name": "i-xxxxxx",
	//
	//           "gpu_list": [
	//
	//             {
	//
	//               "id": "6cc37fbc",
	//
	//               "name": "GPU0",
	//
	//               "style": {
	//
	//                 "fill": "#e94242",
	//
	//                 "opacity": "0.1"
	//
	//               },
	//
	//               "values": {
	//
	//                 "Metric": [
	//
	//                   {
	//
	//                     "key": "0",
	//
	//                     "value": "gpu_util(%)"
	//
	//                   }
	//
	//                 ],
	//
	//                 "Basic": [
	//
	//                   {
	//
	//                     "key": "1",
	//
	//                     "value": "Gpu Number"
	//
	//                   }
	//
	//                 ],
	//
	//                 "Exception": [
	//
	//                   {
	//
	//                     "key": "NCCL HANG",
	//
	//                     "value": "error0"
	//
	//                   }
	//
	//                 ]
	//
	//               }
	//
	//             }
	//
	//           ],
	//
	//           "style": {
	//
	//             "fillOpacity": "0.5",
	//
	//             "fill": "#e94242"
	//
	//           },
	//
	//           "values": {
	//
	//             "Exception": [
	//
	//               {
	//
	//                 "key": "NCCL HANG",
	//
	//                 "value": "error0"
	//
	//               }
	//
	//             ],
	//
	//             "Basic": [
	//
	//               {
	//
	//                 "key": "podList",
	//
	//                 "value": "podname1"
	//
	//               }
	//
	//             ],
	//
	//             "Metric": [
	//
	//               {
	//
	//                 "key": "instance",
	//
	//                 "value": "i-xxxxxx"
	//
	//               }
	//
	//             ]
	//
	//           }
	//
	//         }
	//
	//       ],
	Nodes interface{} `json:"nodes,omitempty" xml:"nodes,omitempty"`
}

func (s GetGpuClusterTopologyResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s GetGpuClusterTopologyResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *GetGpuClusterTopologyResponseBodyDataData) GetEdges() interface{} {
	return s.Edges
}

func (s *GetGpuClusterTopologyResponseBodyDataData) GetNodes() interface{} {
	return s.Nodes
}

func (s *GetGpuClusterTopologyResponseBodyDataData) SetEdges(v interface{}) *GetGpuClusterTopologyResponseBodyDataData {
	s.Edges = v
	return s
}

func (s *GetGpuClusterTopologyResponseBodyDataData) SetNodes(v interface{}) *GetGpuClusterTopologyResponseBodyDataData {
	s.Nodes = v
	return s
}

func (s *GetGpuClusterTopologyResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}
