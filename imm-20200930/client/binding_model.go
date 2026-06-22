// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBinding interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *Binding
	GetCreateTime() *string
	SetDatasetName(v string) *Binding
	GetDatasetName() *string
	SetPhase(v string) *Binding
	GetPhase() *string
	SetProjectName(v string) *Binding
	GetProjectName() *string
	SetReason(v string) *Binding
	GetReason() *string
	SetState(v string) *Binding
	GetState() *string
	SetURI(v string) *Binding
	GetURI() *string
	SetUpdateTime(v string) *Binding
	GetUpdateTime() *string
}

type Binding struct {
	// The timestamp when the binding between the dataset and the OSS bucket was created. The format is RFC3339Nano.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The scan type. Valid values:
	//
	// - FullScanning: A full scan is in progress.
	//
	// - IncrementalScanning: An incremental scan is in progress.
	//
	// example:
	//
	// FullScanning
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Reason
	//
	// example:
	//
	// pause usage
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The state of the binding between the dataset and the OSS bucket. Valid values:
	//
	// - Ready: The binding is being prepared after it is created.
	//
	// - Stopped: The binding is paused.
	//
	// - Running: The binding is running.
	//
	// - Retrying: The binding is being retried after it is created.
	//
	// - Failed: The binding failed to be created.
	//
	// - Deleted: The binding is deleted.
	//
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The URI of the Object Storage Service (OSS) bucket attached to the dataset.
	//
	// The format of an OSS bucket URI is `oss://${bucketname}`. The `bucketname` is the name of an OSS bucket that is in the same region as the current project.
	//
	// example:
	//
	// oss://examplebucket
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// The timestamp when the binding between the dataset and the OSS bucket was last modified. The format is RFC3339Nano.
	//
	// > After a binding is created, if the binding has not been paused or restarted, this timestamp is the same as the creation timestamp.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Binding) String() string {
	return dara.Prettify(s)
}

func (s Binding) GoString() string {
	return s.String()
}

func (s *Binding) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Binding) GetDatasetName() *string {
	return s.DatasetName
}

func (s *Binding) GetPhase() *string {
	return s.Phase
}

func (s *Binding) GetProjectName() *string {
	return s.ProjectName
}

func (s *Binding) GetReason() *string {
	return s.Reason
}

func (s *Binding) GetState() *string {
	return s.State
}

func (s *Binding) GetURI() *string {
	return s.URI
}

func (s *Binding) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *Binding) SetCreateTime(v string) *Binding {
	s.CreateTime = &v
	return s
}

func (s *Binding) SetDatasetName(v string) *Binding {
	s.DatasetName = &v
	return s
}

func (s *Binding) SetPhase(v string) *Binding {
	s.Phase = &v
	return s
}

func (s *Binding) SetProjectName(v string) *Binding {
	s.ProjectName = &v
	return s
}

func (s *Binding) SetReason(v string) *Binding {
	s.Reason = &v
	return s
}

func (s *Binding) SetState(v string) *Binding {
	s.State = &v
	return s
}

func (s *Binding) SetURI(v string) *Binding {
	s.URI = &v
	return s
}

func (s *Binding) SetUpdateTime(v string) *Binding {
	s.UpdateTime = &v
	return s
}

func (s *Binding) Validate() error {
	return dara.Validate(s)
}
