// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryExperimentVisualizationDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetVisualizationData(v []*QueryExperimentVisualizationDataResponseBodyVisualizationData) *QueryExperimentVisualizationDataResponseBody
	GetVisualizationData() []*QueryExperimentVisualizationDataResponseBodyVisualizationData
	SetRequestId(v string) *QueryExperimentVisualizationDataResponseBody
	GetRequestId() *string
}

type QueryExperimentVisualizationDataResponseBody struct {
	VisualizationData []*QueryExperimentVisualizationDataResponseBodyVisualizationData `json:"VisualizationData,omitempty" xml:"VisualizationData,omitempty" type:"Repeated"`
	// example:
	//
	// FFB1D4B4-B253-540A-9B3B-AA711C48A1B7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryExperimentVisualizationDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryExperimentVisualizationDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryExperimentVisualizationDataResponseBody) GetVisualizationData() []*QueryExperimentVisualizationDataResponseBodyVisualizationData {
	return s.VisualizationData
}

func (s *QueryExperimentVisualizationDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryExperimentVisualizationDataResponseBody) SetVisualizationData(v []*QueryExperimentVisualizationDataResponseBodyVisualizationData) *QueryExperimentVisualizationDataResponseBody {
	s.VisualizationData = v
	return s
}

func (s *QueryExperimentVisualizationDataResponseBody) SetRequestId(v string) *QueryExperimentVisualizationDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryExperimentVisualizationDataResponseBody) Validate() error {
	if s.VisualizationData != nil {
		for _, item := range s.VisualizationData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryExperimentVisualizationDataResponseBodyVisualizationData struct {
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1,2,3,4,5
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// dataId1
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// node-ux55ier8axzo2xelcc
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s QueryExperimentVisualizationDataResponseBodyVisualizationData) String() string {
	return dara.Prettify(s)
}

func (s QueryExperimentVisualizationDataResponseBodyVisualizationData) GoString() string {
	return s.String()
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) GetData() *string {
	return s.Data
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) GetDataId() *string {
	return s.DataId
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) GetNodeId() *string {
	return s.NodeId
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) SetCreateTime(v string) *QueryExperimentVisualizationDataResponseBodyVisualizationData {
	s.CreateTime = &v
	return s
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) SetData(v string) *QueryExperimentVisualizationDataResponseBodyVisualizationData {
	s.Data = &v
	return s
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) SetDataId(v string) *QueryExperimentVisualizationDataResponseBodyVisualizationData {
	s.DataId = &v
	return s
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) SetNodeId(v string) *QueryExperimentVisualizationDataResponseBodyVisualizationData {
	s.NodeId = &v
	return s
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) Validate() error {
	return dara.Validate(s)
}
