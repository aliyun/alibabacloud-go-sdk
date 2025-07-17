// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLineageTask interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]*string) *LineageTask
	GetAttributes() map[string]*string
	SetId(v string) *LineageTask
	GetId() *string
	SetType(v string) *LineageTask
	GetType() *string
}

type LineageTask struct {
	Attributes map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// 12345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// custom-sql
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s LineageTask) String() string {
	return dara.Prettify(s)
}

func (s LineageTask) GoString() string {
	return s.String()
}

func (s *LineageTask) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *LineageTask) GetId() *string {
	return s.Id
}

func (s *LineageTask) GetType() *string {
	return s.Type
}

func (s *LineageTask) SetAttributes(v map[string]*string) *LineageTask {
	s.Attributes = v
	return s
}

func (s *LineageTask) SetId(v string) *LineageTask {
	s.Id = &v
	return s
}

func (s *LineageTask) SetType(v string) *LineageTask {
	s.Type = &v
	return s
}

func (s *LineageTask) Validate() error {
	return dara.Validate(s)
}
