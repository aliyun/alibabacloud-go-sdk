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
	// The attempt ID of the Spark application.
	//
	// example:
	//
	// s202207151211hz****-0001
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	// The information about the Spark application.
	Detail *Detail `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The alert message returned, such as task execution failure or insufficient resources. If no alert occurs, null is returned.
	//
	// example:
	//
	// WARN: Disk is full
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The attempt priority of the Spark application.
	//
	// example:
	//
	// NORMAL
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The state of the Spark application. Valid values:
	//
	// 	- **SUBMITTED**
	//
	// 	- **STARTING**
	//
	// 	- **RUNNING**
	//
	// 	- **FAILING**
	//
	// 	- **FAILED**
	//
	// 	- **KILLING**
	//
	// 	- **KILLED**
	//
	// 	- **SUCCEEDING**
	//
	// 	- **COMPLETED**
	//
	// 	- **FATAL**
	//
	// 	- **UNKNOWN**
	//
	// example:
	//
	// SUBMITTED
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
	if s.Detail != nil {
		if err := s.Detail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
