// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobInfo interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *JobInfo
	GetId() *string
	SetProperties(v map[string]interface{}) *JobInfo
	GetProperties() map[string]interface{}
}

type JobInfo struct {
	Id         *string                `json:"id,omitempty" xml:"id,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
}

func (s JobInfo) String() string {
	return dara.Prettify(s)
}

func (s JobInfo) GoString() string {
	return s.String()
}

func (s *JobInfo) GetId() *string {
	return s.Id
}

func (s *JobInfo) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *JobInfo) SetId(v string) *JobInfo {
	s.Id = &v
	return s
}

func (s *JobInfo) SetProperties(v map[string]interface{}) *JobInfo {
	s.Properties = v
	return s
}

func (s *JobInfo) Validate() error {
	return dara.Validate(s)
}
