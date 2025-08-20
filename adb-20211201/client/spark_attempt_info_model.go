// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSparkAttemptInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAttemptId(v string) *SparkAttemptInfo
	GetAttemptId() *string
	SetDetail(v *Detail) *SparkAttemptInfo
	GetDetail() *Detail
	SetMessage(v string) *SparkAttemptInfo
	GetMessage() *string
	SetPriority(v string) *SparkAttemptInfo
	GetPriority() *string
	SetState(v string) *SparkAttemptInfo
	GetState() *string
}

type SparkAttemptInfo struct {
	// example:
	//
	// s202207151211hz0cb4200*****-0001
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	Detail    *Detail `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// WARN: Disk is full
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// NORMAL
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// RUNNING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SparkAttemptInfo) String() string {
	return dara.Prettify(s)
}

func (s SparkAttemptInfo) GoString() string {
	return s.String()
}

func (s *SparkAttemptInfo) GetAttemptId() *string {
	return s.AttemptId
}

func (s *SparkAttemptInfo) GetDetail() *Detail {
	return s.Detail
}

func (s *SparkAttemptInfo) GetMessage() *string {
	return s.Message
}

func (s *SparkAttemptInfo) GetPriority() *string {
	return s.Priority
}

func (s *SparkAttemptInfo) GetState() *string {
	return s.State
}

func (s *SparkAttemptInfo) SetAttemptId(v string) *SparkAttemptInfo {
	s.AttemptId = &v
	return s
}

func (s *SparkAttemptInfo) SetDetail(v *Detail) *SparkAttemptInfo {
	s.Detail = v
	return s
}

func (s *SparkAttemptInfo) SetMessage(v string) *SparkAttemptInfo {
	s.Message = &v
	return s
}

func (s *SparkAttemptInfo) SetPriority(v string) *SparkAttemptInfo {
	s.Priority = &v
	return s
}

func (s *SparkAttemptInfo) SetState(v string) *SparkAttemptInfo {
	s.State = &v
	return s
}

func (s *SparkAttemptInfo) Validate() error {
	return dara.Validate(s)
}
