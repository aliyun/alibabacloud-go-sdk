// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewDataPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasets(v []*PreviewDataPipelineResponseBodyDatasets) *PreviewDataPipelineResponseBody
	GetDatasets() []*PreviewDataPipelineResponseBodyDatasets
	SetEffectiveScript(v string) *PreviewDataPipelineResponseBody
	GetEffectiveScript() *string
	SetRequestId(v string) *PreviewDataPipelineResponseBody
	GetRequestId() *string
}

type PreviewDataPipelineResponseBody struct {
	// The dataset preview results.
	Datasets []*PreviewDataPipelineResponseBodyDatasets `json:"datasets,omitempty" xml:"datasets,omitempty" type:"Repeated"`
	// The effective SPL.
	//
	// example:
	//
	// 	- | where status_code == "ERROR"
	EffectiveScript *string `json:"effectiveScript,omitempty" xml:"effectiveScript,omitempty"`
	// The request ID.
	//
	// example:
	//
	// req-01j2example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PreviewDataPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineResponseBody) GetDatasets() []*PreviewDataPipelineResponseBodyDatasets {
	return s.Datasets
}

func (s *PreviewDataPipelineResponseBody) GetEffectiveScript() *string {
	return s.EffectiveScript
}

func (s *PreviewDataPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewDataPipelineResponseBody) SetDatasets(v []*PreviewDataPipelineResponseBodyDatasets) *PreviewDataPipelineResponseBody {
	s.Datasets = v
	return s
}

func (s *PreviewDataPipelineResponseBody) SetEffectiveScript(v string) *PreviewDataPipelineResponseBody {
	s.EffectiveScript = &v
	return s
}

func (s *PreviewDataPipelineResponseBody) SetRequestId(v string) *PreviewDataPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewDataPipelineResponseBody) Validate() error {
	if s.Datasets != nil {
		for _, item := range s.Datasets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PreviewDataPipelineResponseBodyDatasets struct {
	// The preview data.
	Data []map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The field metadata.
	Meta []*PreviewDataPipelineResponseBodyDatasetsMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The dataset name.
	//
	// example:
	//
	// error_spans
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of samples.
	//
	// example:
	//
	// 1
	SampleCount *int64 `json:"sampleCount,omitempty" xml:"sampleCount,omitempty"`
}

func (s PreviewDataPipelineResponseBodyDatasets) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineResponseBodyDatasets) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineResponseBodyDatasets) GetData() []map[string]interface{} {
	return s.Data
}

func (s *PreviewDataPipelineResponseBodyDatasets) GetMeta() []*PreviewDataPipelineResponseBodyDatasetsMeta {
	return s.Meta
}

func (s *PreviewDataPipelineResponseBodyDatasets) GetName() *string {
	return s.Name
}

func (s *PreviewDataPipelineResponseBodyDatasets) GetSampleCount() *int64 {
	return s.SampleCount
}

func (s *PreviewDataPipelineResponseBodyDatasets) SetData(v []map[string]interface{}) *PreviewDataPipelineResponseBodyDatasets {
	s.Data = v
	return s
}

func (s *PreviewDataPipelineResponseBodyDatasets) SetMeta(v []*PreviewDataPipelineResponseBodyDatasetsMeta) *PreviewDataPipelineResponseBodyDatasets {
	s.Meta = v
	return s
}

func (s *PreviewDataPipelineResponseBodyDatasets) SetName(v string) *PreviewDataPipelineResponseBodyDatasets {
	s.Name = &v
	return s
}

func (s *PreviewDataPipelineResponseBodyDatasets) SetSampleCount(v int64) *PreviewDataPipelineResponseBodyDatasets {
	s.SampleCount = &v
	return s
}

func (s *PreviewDataPipelineResponseBodyDatasets) Validate() error {
	if s.Meta != nil {
		for _, item := range s.Meta {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PreviewDataPipelineResponseBodyDatasetsMeta struct {
	// The field name.
	//
	// example:
	//
	// trace_id
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The field type.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PreviewDataPipelineResponseBodyDatasetsMeta) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineResponseBodyDatasetsMeta) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineResponseBodyDatasetsMeta) GetName() *string {
	return s.Name
}

func (s *PreviewDataPipelineResponseBodyDatasetsMeta) GetType() *string {
	return s.Type
}

func (s *PreviewDataPipelineResponseBodyDatasetsMeta) SetName(v string) *PreviewDataPipelineResponseBodyDatasetsMeta {
	s.Name = &v
	return s
}

func (s *PreviewDataPipelineResponseBodyDatasetsMeta) SetType(v string) *PreviewDataPipelineResponseBodyDatasetsMeta {
	s.Type = &v
	return s
}

func (s *PreviewDataPipelineResponseBodyDatasetsMeta) Validate() error {
	return dara.Validate(s)
}
