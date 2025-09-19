// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageEventOperationConditionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventType(v string) *DescribeImageEventOperationConditionRequest
	GetEventType() *string
	SetLang(v string) *DescribeImageEventOperationConditionRequest
	GetLang() *string
}

type DescribeImageEventOperationConditionRequest struct {
	// The alert type.
	//
	// 	- Set the value to **sensitiveFile**.
	//
	// example:
	//
	// sensitiveFile
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeImageEventOperationConditionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageEventOperationConditionRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageEventOperationConditionRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeImageEventOperationConditionRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageEventOperationConditionRequest) SetEventType(v string) *DescribeImageEventOperationConditionRequest {
	s.EventType = &v
	return s
}

func (s *DescribeImageEventOperationConditionRequest) SetLang(v string) *DescribeImageEventOperationConditionRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageEventOperationConditionRequest) Validate() error {
	return dara.Validate(s)
}
