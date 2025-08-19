// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTaskOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListAsyncTaskOutput
	GetNextToken() *string
	SetTasks(v []*AsyncTask) *ListAsyncTaskOutput
	GetTasks() []*AsyncTask
}

type ListAsyncTaskOutput struct {
	NextToken *string      `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Tasks     []*AsyncTask `json:"tasks" xml:"tasks" type:"Repeated"`
}

func (s ListAsyncTaskOutput) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTaskOutput) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAsyncTaskOutput) GetTasks() []*AsyncTask {
	return s.Tasks
}

func (s *ListAsyncTaskOutput) SetNextToken(v string) *ListAsyncTaskOutput {
	s.NextToken = &v
	return s
}

func (s *ListAsyncTaskOutput) SetTasks(v []*AsyncTask) *ListAsyncTaskOutput {
	s.Tasks = v
	return s
}

func (s *ListAsyncTaskOutput) Validate() error {
	return dara.Validate(s)
}
