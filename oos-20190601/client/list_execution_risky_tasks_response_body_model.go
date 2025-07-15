// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutionRiskyTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListExecutionRiskyTasksResponseBody
	GetRequestId() *string
	SetRiskyTasks(v []*ListExecutionRiskyTasksResponseBodyRiskyTasks) *ListExecutionRiskyTasksResponseBody
	GetRiskyTasks() []*ListExecutionRiskyTasksResponseBodyRiskyTasks
}

type ListExecutionRiskyTasksResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C04B668D-D2DD-4B40-B6E9-0E3C4F53D5B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about high-risk tasks.
	RiskyTasks []*ListExecutionRiskyTasksResponseBodyRiskyTasks `json:"RiskyTasks,omitempty" xml:"RiskyTasks,omitempty" type:"Repeated"`
}

func (s ListExecutionRiskyTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionRiskyTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutionRiskyTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExecutionRiskyTasksResponseBody) GetRiskyTasks() []*ListExecutionRiskyTasksResponseBodyRiskyTasks {
	return s.RiskyTasks
}

func (s *ListExecutionRiskyTasksResponseBody) SetRequestId(v string) *ListExecutionRiskyTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutionRiskyTasksResponseBody) SetRiskyTasks(v []*ListExecutionRiskyTasksResponseBodyRiskyTasks) *ListExecutionRiskyTasksResponseBody {
	s.RiskyTasks = v
	return s
}

func (s *ListExecutionRiskyTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListExecutionRiskyTasksResponseBodyRiskyTasks struct {
	// The name of the operation that the high-risk task calls.
	//
	// example:
	//
	// DeleteInstance
	API *string `json:"API,omitempty" xml:"API,omitempty"`
	// The cloud service in which the high-risk task runs.
	//
	// example:
	//
	// ECS
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The details of the high-risk task.
	Task []*string `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
	// The details of templates to which the high-risk task belongs.
	Template []*string `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s ListExecutionRiskyTasksResponseBodyRiskyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionRiskyTasksResponseBodyRiskyTasks) GoString() string {
	return s.String()
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) GetAPI() *string {
	return s.API
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) GetService() *string {
	return s.Service
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) GetTask() []*string {
	return s.Task
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) GetTemplate() []*string {
	return s.Template
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) SetAPI(v string) *ListExecutionRiskyTasksResponseBodyRiskyTasks {
	s.API = &v
	return s
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) SetService(v string) *ListExecutionRiskyTasksResponseBodyRiskyTasks {
	s.Service = &v
	return s
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) SetTask(v []*string) *ListExecutionRiskyTasksResponseBodyRiskyTasks {
	s.Task = v
	return s
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) SetTemplate(v []*string) *ListExecutionRiskyTasksResponseBodyRiskyTasks {
	s.Template = v
	return s
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) Validate() error {
	return dara.Validate(s)
}
