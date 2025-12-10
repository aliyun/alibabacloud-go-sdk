// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryExperimentVisualizationDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*QueryExperimentVisualizationDataRequestBody) *QueryExperimentVisualizationDataRequest
	GetBody() []*QueryExperimentVisualizationDataRequestBody
}

type QueryExperimentVisualizationDataRequest struct {
	// This parameter is required.
	Body []*QueryExperimentVisualizationDataRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s QueryExperimentVisualizationDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryExperimentVisualizationDataRequest) GoString() string {
	return s.String()
}

func (s *QueryExperimentVisualizationDataRequest) GetBody() []*QueryExperimentVisualizationDataRequestBody {
	return s.Body
}

func (s *QueryExperimentVisualizationDataRequest) SetBody(v []*QueryExperimentVisualizationDataRequestBody) *QueryExperimentVisualizationDataRequest {
	s.Body = v
	return s
}

func (s *QueryExperimentVisualizationDataRequest) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryExperimentVisualizationDataRequestBody struct {
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// node-2dfd8xfjda
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	StartTime            *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	VisualizationDataIds []*string `json:"VisualizationDataIds,omitempty" xml:"VisualizationDataIds,omitempty" type:"Repeated"`
}

func (s QueryExperimentVisualizationDataRequestBody) String() string {
	return dara.Prettify(s)
}

func (s QueryExperimentVisualizationDataRequestBody) GoString() string {
	return s.String()
}

func (s *QueryExperimentVisualizationDataRequestBody) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryExperimentVisualizationDataRequestBody) GetNodeId() *string {
	return s.NodeId
}

func (s *QueryExperimentVisualizationDataRequestBody) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryExperimentVisualizationDataRequestBody) GetVisualizationDataIds() []*string {
	return s.VisualizationDataIds
}

func (s *QueryExperimentVisualizationDataRequestBody) SetEndTime(v string) *QueryExperimentVisualizationDataRequestBody {
	s.EndTime = &v
	return s
}

func (s *QueryExperimentVisualizationDataRequestBody) SetNodeId(v string) *QueryExperimentVisualizationDataRequestBody {
	s.NodeId = &v
	return s
}

func (s *QueryExperimentVisualizationDataRequestBody) SetStartTime(v string) *QueryExperimentVisualizationDataRequestBody {
	s.StartTime = &v
	return s
}

func (s *QueryExperimentVisualizationDataRequestBody) SetVisualizationDataIds(v []*string) *QueryExperimentVisualizationDataRequestBody {
	s.VisualizationDataIds = v
	return s
}

func (s *QueryExperimentVisualizationDataRequestBody) Validate() error {
	return dara.Validate(s)
}
