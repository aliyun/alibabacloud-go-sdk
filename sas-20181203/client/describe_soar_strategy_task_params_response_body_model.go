// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyTaskParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSoarStrategyTaskParamsResponseBody
	GetRequestId() *string
	SetTaskDetail(v *DescribeSoarStrategyTaskParamsResponseBodyTaskDetail) *DescribeSoarStrategyTaskParamsResponseBody
	GetTaskDetail() *DescribeSoarStrategyTaskParamsResponseBodyTaskDetail
}

type DescribeSoarStrategyTaskParamsResponseBody struct {
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// ACF97412-FD09-4D1F-994F-34DF12BR****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task details.
	TaskDetail *DescribeSoarStrategyTaskParamsResponseBodyTaskDetail `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty" type:"Struct"`
}

func (s DescribeSoarStrategyTaskParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTaskParamsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTaskParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSoarStrategyTaskParamsResponseBody) GetTaskDetail() *DescribeSoarStrategyTaskParamsResponseBodyTaskDetail {
	return s.TaskDetail
}

func (s *DescribeSoarStrategyTaskParamsResponseBody) SetRequestId(v string) *DescribeSoarStrategyTaskParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarStrategyTaskParamsResponseBody) SetTaskDetail(v *DescribeSoarStrategyTaskParamsResponseBodyTaskDetail) *DescribeSoarStrategyTaskParamsResponseBody {
	s.TaskDetail = v
	return s
}

func (s *DescribeSoarStrategyTaskParamsResponseBody) Validate() error {
	if s.TaskDetail != nil {
		if err := s.TaskDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSoarStrategyTaskParamsResponseBodyTaskDetail struct {
	// The task parameters.
	//
	// example:
	//
	// {"summary":[{"name":"email","type":"String","isRequired":false,"fromProperty":"notifyConfig.email"}]}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The task name.
	//
	// example:
	//
	// task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeSoarStrategyTaskParamsResponseBodyTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTaskParamsResponseBodyTaskDetail) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTaskParamsResponseBodyTaskDetail) GetParams() *string {
	return s.Params
}

func (s *DescribeSoarStrategyTaskParamsResponseBodyTaskDetail) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeSoarStrategyTaskParamsResponseBodyTaskDetail) SetParams(v string) *DescribeSoarStrategyTaskParamsResponseBodyTaskDetail {
	s.Params = &v
	return s
}

func (s *DescribeSoarStrategyTaskParamsResponseBodyTaskDetail) SetTaskName(v string) *DescribeSoarStrategyTaskParamsResponseBodyTaskDetail {
	s.TaskName = &v
	return s
}

func (s *DescribeSoarStrategyTaskParamsResponseBodyTaskDetail) Validate() error {
	return dara.Validate(s)
}
