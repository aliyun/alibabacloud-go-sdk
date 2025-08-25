// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*ScanTextRequestLabels) *ScanTextRequest
	GetLabels() []*ScanTextRequestLabels
	SetTasks(v []*ScanTextRequestTasks) *ScanTextRequest
	GetTasks() []*ScanTextRequestTasks
}

type ScanTextRequest struct {
	// 1
	//
	// This parameter is required.
	Labels []*ScanTextRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// 1
	//
	// This parameter is required.
	Tasks []*ScanTextRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ScanTextRequest) String() string {
	return dara.Prettify(s)
}

func (s ScanTextRequest) GoString() string {
	return s.String()
}

func (s *ScanTextRequest) GetLabels() []*ScanTextRequestLabels {
	return s.Labels
}

func (s *ScanTextRequest) GetTasks() []*ScanTextRequestTasks {
	return s.Tasks
}

func (s *ScanTextRequest) SetLabels(v []*ScanTextRequestLabels) *ScanTextRequest {
	s.Labels = v
	return s
}

func (s *ScanTextRequest) SetTasks(v []*ScanTextRequestTasks) *ScanTextRequest {
	s.Tasks = v
	return s
}

func (s *ScanTextRequest) Validate() error {
	return dara.Validate(s)
}

type ScanTextRequestLabels struct {
	// This parameter is required.
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ScanTextRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s ScanTextRequestLabels) GoString() string {
	return s.String()
}

func (s *ScanTextRequestLabels) GetLabel() *string {
	return s.Label
}

func (s *ScanTextRequestLabels) SetLabel(v string) *ScanTextRequestLabels {
	s.Label = &v
	return s
}

func (s *ScanTextRequestLabels) Validate() error {
	return dara.Validate(s)
}

type ScanTextRequestTasks struct {
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ScanTextRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s ScanTextRequestTasks) GoString() string {
	return s.String()
}

func (s *ScanTextRequestTasks) GetContent() *string {
	return s.Content
}

func (s *ScanTextRequestTasks) SetContent(v string) *ScanTextRequestTasks {
	s.Content = &v
	return s
}

func (s *ScanTextRequestTasks) Validate() error {
	return dara.Validate(s)
}
