// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOnlineTestResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *DescribeOnlineTestResultRequest
	GetResourceType() *string
	SetServiceCode(v string) *DescribeOnlineTestResultRequest
	GetServiceCode() *string
	SetTaskId(v string) *DescribeOnlineTestResultRequest
	GetTaskId() *string
}

type DescribeOnlineTestResultRequest struct {
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// VideoModeration
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeOnlineTestResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnlineTestResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnlineTestResultRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeOnlineTestResultRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeOnlineTestResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeOnlineTestResultRequest) SetResourceType(v string) *DescribeOnlineTestResultRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeOnlineTestResultRequest) SetServiceCode(v string) *DescribeOnlineTestResultRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeOnlineTestResultRequest) SetTaskId(v string) *DescribeOnlineTestResultRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeOnlineTestResultRequest) Validate() error {
	return dara.Validate(s)
}
