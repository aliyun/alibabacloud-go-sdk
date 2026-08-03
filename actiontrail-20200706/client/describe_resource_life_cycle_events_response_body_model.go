// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLifeCycleEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DescribeResourceLifeCycleEventsResponseBody
	GetData() *string
	SetRequestId(v string) *DescribeResourceLifeCycleEventsResponseBody
	GetRequestId() *string
}

type DescribeResourceLifeCycleEventsResponseBody struct {
	// The lifecycle events.<br>This field is returned as a JSON-serialized string. The string contains the hierarchical data for lifecycle event categories. Use a standard JSON deserialization tool for your programming language to parse the string into an array of objects.
	//
	// example:
	//
	// [{"children":[{"children":[{"label":"Create Events","labelEn":"Create Events","value":"Create,CreateInstance,RunInstances"},{"label":"Delete Events","labelEn":"Delete Events","value":"DeleteInstance,DeleteInstances,Release"}],"label":"ECS Instance","labelEn":"ECS Instance","value":"ACS::ECS::Instance"}],"label":"Elastic Compute Service","labelEn":"Elastic Compute Service","value":"Ecs"}]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B10969CF-C743-55F8-9710-F0711504****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeResourceLifeCycleEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLifeCycleEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceLifeCycleEventsResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeResourceLifeCycleEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceLifeCycleEventsResponseBody) SetData(v string) *DescribeResourceLifeCycleEventsResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeResourceLifeCycleEventsResponseBody) SetRequestId(v string) *DescribeResourceLifeCycleEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceLifeCycleEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
