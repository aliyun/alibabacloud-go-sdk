// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataStreamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsManaged(v bool) *ListDataStreamsRequest
	GetIsManaged() *bool
	SetName(v string) *ListDataStreamsRequest
	GetName() *string
}

type ListDataStreamsRequest struct {
	// example:
	//
	// false
	IsManaged *bool `json:"isManaged,omitempty" xml:"isManaged,omitempty"`
	// example:
	//
	// Log1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListDataStreamsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataStreamsRequest) GoString() string {
	return s.String()
}

func (s *ListDataStreamsRequest) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListDataStreamsRequest) GetName() *string {
	return s.Name
}

func (s *ListDataStreamsRequest) SetIsManaged(v bool) *ListDataStreamsRequest {
	s.IsManaged = &v
	return s
}

func (s *ListDataStreamsRequest) SetName(v string) *ListDataStreamsRequest {
	s.Name = &v
	return s
}

func (s *ListDataStreamsRequest) Validate() error {
	return dara.Validate(s)
}
