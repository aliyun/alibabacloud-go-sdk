// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentVisualizationMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetVisualizationMeta(v []*GetExperimentVisualizationMetaResponseBodyVisualizationMeta) *GetExperimentVisualizationMetaResponseBody
	GetVisualizationMeta() []*GetExperimentVisualizationMetaResponseBodyVisualizationMeta
	SetRequestId(v string) *GetExperimentVisualizationMetaResponseBody
	GetRequestId() *string
}

type GetExperimentVisualizationMetaResponseBody struct {
	VisualizationMeta []*GetExperimentVisualizationMetaResponseBodyVisualizationMeta `json:"VisualizationMeta,omitempty" xml:"VisualizationMeta,omitempty" type:"Repeated"`
	// example:
	//
	// A84A1282-D3E7-5198-9E8E-2AD09C78C6C1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetExperimentVisualizationMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentVisualizationMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentVisualizationMetaResponseBody) GetVisualizationMeta() []*GetExperimentVisualizationMetaResponseBodyVisualizationMeta {
	return s.VisualizationMeta
}

func (s *GetExperimentVisualizationMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExperimentVisualizationMetaResponseBody) SetVisualizationMeta(v []*GetExperimentVisualizationMetaResponseBodyVisualizationMeta) *GetExperimentVisualizationMetaResponseBody {
	s.VisualizationMeta = v
	return s
}

func (s *GetExperimentVisualizationMetaResponseBody) SetRequestId(v string) *GetExperimentVisualizationMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentVisualizationMetaResponseBody) Validate() error {
	if s.VisualizationMeta != nil {
		for _, item := range s.VisualizationMeta {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetExperimentVisualizationMetaResponseBodyVisualizationMeta struct {
	// example:
	//
	// {
	//
	// 	"locations": [{
	//
	// 		"id": "result_table",
	//
	// 		"location": {
	//
	// 			"project": "mulan_test_pre_1",
	//
	// 			"endpoint": "http://service.cn.maxcompute.aliyun-inc.com/api",
	//
	// 			"table": "pai_temp_flow_qzkkjqic95olnrel1w_node_7hc1rdsa99gy2msbvc_visualizationTable"
	//
	// 		},
	//
	// 		"locationType": "MaxComputeTable"
	//
	// 	}],
	//
	// 	"components": [{
	//
	// 		"id": "histogram-chart",
	//
	// 		"dataId": "histogram_result"
	//
	// 	}],
	//
	// 	"dataInfos": [{
	//
	// 		"id": "histogram_result",
	//
	// 		"locationId": "result_table",
	//
	// 		"dataType": "json"
	//
	// 	}]
	//
	// }
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// node_id1
	NodeId   *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetExperimentVisualizationMetaResponseBodyVisualizationMeta) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentVisualizationMetaResponseBodyVisualizationMeta) GoString() string {
	return s.String()
}

func (s *GetExperimentVisualizationMetaResponseBodyVisualizationMeta) GetMeta() *string {
	return s.Meta
}

func (s *GetExperimentVisualizationMetaResponseBodyVisualizationMeta) GetNodeId() *string {
	return s.NodeId
}

func (s *GetExperimentVisualizationMetaResponseBodyVisualizationMeta) GetNodeName() *string {
	return s.NodeName
}

func (s *GetExperimentVisualizationMetaResponseBodyVisualizationMeta) SetMeta(v string) *GetExperimentVisualizationMetaResponseBodyVisualizationMeta {
	s.Meta = &v
	return s
}

func (s *GetExperimentVisualizationMetaResponseBodyVisualizationMeta) SetNodeId(v string) *GetExperimentVisualizationMetaResponseBodyVisualizationMeta {
	s.NodeId = &v
	return s
}

func (s *GetExperimentVisualizationMetaResponseBodyVisualizationMeta) SetNodeName(v string) *GetExperimentVisualizationMetaResponseBodyVisualizationMeta {
	s.NodeName = &v
	return s
}

func (s *GetExperimentVisualizationMetaResponseBodyVisualizationMeta) Validate() error {
	return dara.Validate(s)
}
