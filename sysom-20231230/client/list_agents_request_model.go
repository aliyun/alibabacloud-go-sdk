// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int64) *ListAgentsRequest
	GetCurrent() *int64
	SetName(v string) *ListAgentsRequest
	GetName() *string
	SetPageSize(v int64) *ListAgentsRequest
	GetPageSize() *int64
	SetType(v string) *ListAgentsRequest
	GetType() *string
}

type ListAgentsRequest struct {
	Current  *int64  `json:"current,omitempty" xml:"current,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	PageSize *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentsRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListAgentsRequest) GetName() *string {
	return s.Name
}

func (s *ListAgentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAgentsRequest) GetType() *string {
	return s.Type
}

func (s *ListAgentsRequest) SetCurrent(v int64) *ListAgentsRequest {
	s.Current = &v
	return s
}

func (s *ListAgentsRequest) SetName(v string) *ListAgentsRequest {
	s.Name = &v
	return s
}

func (s *ListAgentsRequest) SetPageSize(v int64) *ListAgentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentsRequest) SetType(v string) *ListAgentsRequest {
	s.Type = &v
	return s
}

func (s *ListAgentsRequest) Validate() error {
	return dara.Validate(s)
}
