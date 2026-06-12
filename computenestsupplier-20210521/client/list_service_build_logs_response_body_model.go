// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceBuildLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildLogs(v []*ListServiceBuildLogsResponseBodyBuildLogs) *ListServiceBuildLogsResponseBody
	GetBuildLogs() []*ListServiceBuildLogsResponseBodyBuildLogs
	SetNextToken(v string) *ListServiceBuildLogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceBuildLogsResponseBody
	GetRequestId() *string
}

type ListServiceBuildLogsResponseBody struct {
	// An array of build logs.
	BuildLogs []*ListServiceBuildLogsResponseBodyBuildLogs `json:"BuildLogs,omitempty" xml:"BuildLogs,omitempty" type:"Repeated"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// AAAAAW8kZY+u1sYOaYf5JmgmDQQ=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListServiceBuildLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceBuildLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceBuildLogsResponseBody) GetBuildLogs() []*ListServiceBuildLogsResponseBodyBuildLogs {
	return s.BuildLogs
}

func (s *ListServiceBuildLogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceBuildLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceBuildLogsResponseBody) SetBuildLogs(v []*ListServiceBuildLogsResponseBodyBuildLogs) *ListServiceBuildLogsResponseBody {
	s.BuildLogs = v
	return s
}

func (s *ListServiceBuildLogsResponseBody) SetNextToken(v string) *ListServiceBuildLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceBuildLogsResponseBody) SetRequestId(v string) *ListServiceBuildLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceBuildLogsResponseBody) Validate() error {
	if s.BuildLogs != nil {
		for _, item := range s.BuildLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceBuildLogsResponseBodyBuildLogs struct {
	// The name of the build step.
	//
	// example:
	//
	// BuildArtifact
	BuildStep *string `json:"BuildStep,omitempty" xml:"BuildStep,omitempty"`
	// The content of the log.
	//
	// example:
	//
	// build log
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The timestamp of the log entry.
	//
	// example:
	//
	// yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\"
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListServiceBuildLogsResponseBodyBuildLogs) String() string {
	return dara.Prettify(s)
}

func (s ListServiceBuildLogsResponseBodyBuildLogs) GoString() string {
	return s.String()
}

func (s *ListServiceBuildLogsResponseBodyBuildLogs) GetBuildStep() *string {
	return s.BuildStep
}

func (s *ListServiceBuildLogsResponseBodyBuildLogs) GetContent() *string {
	return s.Content
}

func (s *ListServiceBuildLogsResponseBodyBuildLogs) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListServiceBuildLogsResponseBodyBuildLogs) SetBuildStep(v string) *ListServiceBuildLogsResponseBodyBuildLogs {
	s.BuildStep = &v
	return s
}

func (s *ListServiceBuildLogsResponseBodyBuildLogs) SetContent(v string) *ListServiceBuildLogsResponseBodyBuildLogs {
	s.Content = &v
	return s
}

func (s *ListServiceBuildLogsResponseBodyBuildLogs) SetTimestamp(v string) *ListServiceBuildLogsResponseBodyBuildLogs {
	s.Timestamp = &v
	return s
}

func (s *ListServiceBuildLogsResponseBodyBuildLogs) Validate() error {
	return dara.Validate(s)
}
