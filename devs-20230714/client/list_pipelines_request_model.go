// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelSelector(v []*string) *ListPipelinesRequest
	GetLabelSelector() []*string
}

type ListPipelinesRequest struct {
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
}

func (s ListPipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequest) GetLabelSelector() []*string {
	return s.LabelSelector
}

func (s *ListPipelinesRequest) SetLabelSelector(v []*string) *ListPipelinesRequest {
	s.LabelSelector = v
	return s
}

func (s *ListPipelinesRequest) Validate() error {
	return dara.Validate(s)
}
