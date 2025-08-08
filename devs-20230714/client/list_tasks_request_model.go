// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelSelector(v []*string) *ListTasksRequest
	GetLabelSelector() []*string
}

type ListTasksRequest struct {
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
}

func (s ListTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) GetLabelSelector() []*string {
	return s.LabelSelector
}

func (s *ListTasksRequest) SetLabelSelector(v []*string) *ListTasksRequest {
	s.LabelSelector = v
	return s
}

func (s *ListTasksRequest) Validate() error {
	return dara.Validate(s)
}
