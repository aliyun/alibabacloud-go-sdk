// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommonLogDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *GetCommonLogDetailResponseBody
	GetAction() *string
	SetClusterId(v string) *GetCommonLogDetailResponseBody
	GetClusterId() *string
	SetLogDetail(v []*GetCommonLogDetailResponseBodyLogDetail) *GetCommonLogDetailResponseBody
	GetLogDetail() []*GetCommonLogDetailResponseBodyLogDetail
	SetLogType(v string) *GetCommonLogDetailResponseBody
	GetLogType() *string
	SetOperatorUid(v string) *GetCommonLogDetailResponseBody
	GetOperatorUid() *string
	SetRequestId(v string) *GetCommonLogDetailResponseBody
	GetRequestId() *string
	SetUid(v string) *GetCommonLogDetailResponseBody
	GetUid() *string
}

type GetCommonLogDetailResponseBody struct {
	// The action name.
	//
	// example:
	//
	// CreateCluster
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-abc***
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the logs.
	LogDetail []*GetCommonLogDetailResponseBodyLogDetail `json:"LogDetail,omitempty" xml:"LogDetail,omitempty" type:"Repeated"`
	// The log type.
	//
	// example:
	//
	// operation
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The account ID of the operator.
	//
	// example:
	//
	// 239***
	OperatorUid *string `json:"OperatorUid,omitempty" xml:"OperatorUid,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 464E9919-D04F-4D1D-B375-15989492****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 137***
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetCommonLogDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCommonLogDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommonLogDetailResponseBody) GetAction() *string {
	return s.Action
}

func (s *GetCommonLogDetailResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetCommonLogDetailResponseBody) GetLogDetail() []*GetCommonLogDetailResponseBodyLogDetail {
	return s.LogDetail
}

func (s *GetCommonLogDetailResponseBody) GetLogType() *string {
	return s.LogType
}

func (s *GetCommonLogDetailResponseBody) GetOperatorUid() *string {
	return s.OperatorUid
}

func (s *GetCommonLogDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCommonLogDetailResponseBody) GetUid() *string {
	return s.Uid
}

func (s *GetCommonLogDetailResponseBody) SetAction(v string) *GetCommonLogDetailResponseBody {
	s.Action = &v
	return s
}

func (s *GetCommonLogDetailResponseBody) SetClusterId(v string) *GetCommonLogDetailResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetCommonLogDetailResponseBody) SetLogDetail(v []*GetCommonLogDetailResponseBodyLogDetail) *GetCommonLogDetailResponseBody {
	s.LogDetail = v
	return s
}

func (s *GetCommonLogDetailResponseBody) SetLogType(v string) *GetCommonLogDetailResponseBody {
	s.LogType = &v
	return s
}

func (s *GetCommonLogDetailResponseBody) SetOperatorUid(v string) *GetCommonLogDetailResponseBody {
	s.OperatorUid = &v
	return s
}

func (s *GetCommonLogDetailResponseBody) SetRequestId(v string) *GetCommonLogDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCommonLogDetailResponseBody) SetUid(v string) *GetCommonLogDetailResponseBody {
	s.Uid = &v
	return s
}

func (s *GetCommonLogDetailResponseBody) Validate() error {
	if s.LogDetail != nil {
		for _, item := range s.LogDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCommonLogDetailResponseBodyLogDetail struct {
	// The stage name of the log.
	//
	// example:
	//
	// ConfigNetwork
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The information about the log stages.
	Stages []*GetCommonLogDetailResponseBodyLogDetailStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
}

func (s GetCommonLogDetailResponseBodyLogDetail) String() string {
	return dara.Prettify(s)
}

func (s GetCommonLogDetailResponseBodyLogDetail) GoString() string {
	return s.String()
}

func (s *GetCommonLogDetailResponseBodyLogDetail) GetStageName() *string {
	return s.StageName
}

func (s *GetCommonLogDetailResponseBodyLogDetail) GetStages() []*GetCommonLogDetailResponseBodyLogDetailStages {
	return s.Stages
}

func (s *GetCommonLogDetailResponseBodyLogDetail) SetStageName(v string) *GetCommonLogDetailResponseBodyLogDetail {
	s.StageName = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetail) SetStages(v []*GetCommonLogDetailResponseBodyLogDetailStages) *GetCommonLogDetailResponseBodyLogDetail {
	s.Stages = v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetail) Validate() error {
	if s.Stages != nil {
		for _, item := range s.Stages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCommonLogDetailResponseBodyLogDetailStages struct {
	// The log level.
	//
	// Valid values:
	//
	// 	- ERROR
	//
	// 	- INFO
	//
	// 	- WARN
	//
	// example:
	//
	// INFO
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// The output message of the log.
	//
	// example:
	//
	// Successfully created security group sg-bcd***
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The method involved in the log.
	//
	// example:
	//
	// CreateSecurityGroup
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The request ID associated with the log.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The action state involved in the log. Valid values:
	//
	// 	- InProgress: The action is being performed.
	//
	// 	- Finished: The action is completed.
	//
	// 	- Failed: The action failed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The resource involved in the log.
	//
	// example:
	//
	// sg-bcd***
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The time when the log was generated.
	//
	// example:
	//
	// 2024-08-22 14:21:54
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetCommonLogDetailResponseBodyLogDetailStages) String() string {
	return dara.Prettify(s)
}

func (s GetCommonLogDetailResponseBodyLogDetailStages) GoString() string {
	return s.String()
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) GetLogLevel() *string {
	return s.LogLevel
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) GetMessage() *string {
	return s.Message
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) GetMethod() *string {
	return s.Method
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) GetStatus() *string {
	return s.Status
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) GetTarget() *string {
	return s.Target
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) GetTime() *string {
	return s.Time
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetLogLevel(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.LogLevel = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetMessage(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.Message = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetMethod(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.Method = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetRequestId(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.RequestId = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetStatus(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.Status = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetTarget(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.Target = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetTime(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.Time = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) Validate() error {
	return dara.Validate(s)
}
