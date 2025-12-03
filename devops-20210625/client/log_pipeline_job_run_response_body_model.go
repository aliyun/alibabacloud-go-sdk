// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogPipelineJobRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *LogPipelineJobRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *LogPipelineJobRunResponseBody
	GetErrorMessage() *string
	SetLog(v *LogPipelineJobRunResponseBodyLog) *LogPipelineJobRunResponseBody
	GetLog() *LogPipelineJobRunResponseBodyLog
	SetRequestId(v string) *LogPipelineJobRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *LogPipelineJobRunResponseBody
	GetSuccess() *bool
}

type LogPipelineJobRunResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string                           `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Log          *LogPipelineJobRunResponseBodyLog `json:"log,omitempty" xml:"log,omitempty" type:"Struct"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LogPipelineJobRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LogPipelineJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *LogPipelineJobRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *LogPipelineJobRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *LogPipelineJobRunResponseBody) GetLog() *LogPipelineJobRunResponseBodyLog {
	return s.Log
}

func (s *LogPipelineJobRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LogPipelineJobRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *LogPipelineJobRunResponseBody) SetErrorCode(v string) *LogPipelineJobRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *LogPipelineJobRunResponseBody) SetErrorMessage(v string) *LogPipelineJobRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *LogPipelineJobRunResponseBody) SetLog(v *LogPipelineJobRunResponseBodyLog) *LogPipelineJobRunResponseBody {
	s.Log = v
	return s
}

func (s *LogPipelineJobRunResponseBody) SetRequestId(v string) *LogPipelineJobRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *LogPipelineJobRunResponseBody) SetSuccess(v bool) *LogPipelineJobRunResponseBody {
	s.Success = &v
	return s
}

func (s *LogPipelineJobRunResponseBody) Validate() error {
	if s.Log != nil {
		if err := s.Log.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LogPipelineJobRunResponseBodyLog struct {
	// example:
	//
	// success
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// true
	More *bool `json:"more,omitempty" xml:"more,omitempty"`
}

func (s LogPipelineJobRunResponseBodyLog) String() string {
	return dara.Prettify(s)
}

func (s LogPipelineJobRunResponseBodyLog) GoString() string {
	return s.String()
}

func (s *LogPipelineJobRunResponseBodyLog) GetContent() *string {
	return s.Content
}

func (s *LogPipelineJobRunResponseBodyLog) GetMore() *bool {
	return s.More
}

func (s *LogPipelineJobRunResponseBodyLog) SetContent(v string) *LogPipelineJobRunResponseBodyLog {
	s.Content = &v
	return s
}

func (s *LogPipelineJobRunResponseBodyLog) SetMore(v bool) *LogPipelineJobRunResponseBodyLog {
	s.More = &v
	return s
}

func (s *LogPipelineJobRunResponseBodyLog) Validate() error {
	return dara.Validate(s)
}
