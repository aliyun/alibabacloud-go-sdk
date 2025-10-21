// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelation interface {
	dara.Model
	String() string
	GoString() string
	SetDestination(v string) *Relation
	GetDestination() *string
	SetJobId(v string) *Relation
	GetJobId() *string
	SetSource(v string) *Relation
	GetSource() *string
}

type Relation struct {
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty"`
	JobId       *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	Source      *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s Relation) String() string {
	return dara.Prettify(s)
}

func (s Relation) GoString() string {
	return s.String()
}

func (s *Relation) GetDestination() *string {
	return s.Destination
}

func (s *Relation) GetJobId() *string {
	return s.JobId
}

func (s *Relation) GetSource() *string {
	return s.Source
}

func (s *Relation) SetDestination(v string) *Relation {
	s.Destination = &v
	return s
}

func (s *Relation) SetJobId(v string) *Relation {
	s.JobId = &v
	return s
}

func (s *Relation) SetSource(v string) *Relation {
	s.Source = &v
	return s
}

func (s *Relation) Validate() error {
	return dara.Validate(s)
}
