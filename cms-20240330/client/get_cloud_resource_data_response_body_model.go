// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudResourceDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v [][]*string) *GetCloudResourceDataResponseBody
	GetData() [][]*string
	SetHeader(v []*string) *GetCloudResourceDataResponseBody
	GetHeader() []*string
	SetRequestId(v string) *GetCloudResourceDataResponseBody
	GetRequestId() *string
	SetResponseStatus(v *GetCloudResourceDataResponseBodyResponseStatus) *GetCloudResourceDataResponseBody
	GetResponseStatus() *GetCloudResourceDataResponseBodyResponseStatus
}

type GetCloudResourceDataResponseBody struct {
	// The total list of returned data.
	Data [][]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The list of headers.
	Header []*string `json:"header,omitempty" xml:"header,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result status.
	ResponseStatus *GetCloudResourceDataResponseBodyResponseStatus `json:"responseStatus,omitempty" xml:"responseStatus,omitempty" type:"Struct"`
}

func (s GetCloudResourceDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudResourceDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudResourceDataResponseBody) GetData() [][]*string {
	return s.Data
}

func (s *GetCloudResourceDataResponseBody) GetHeader() []*string {
	return s.Header
}

func (s *GetCloudResourceDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudResourceDataResponseBody) GetResponseStatus() *GetCloudResourceDataResponseBodyResponseStatus {
	return s.ResponseStatus
}

func (s *GetCloudResourceDataResponseBody) SetData(v [][]*string) *GetCloudResourceDataResponseBody {
	s.Data = v
	return s
}

func (s *GetCloudResourceDataResponseBody) SetHeader(v []*string) *GetCloudResourceDataResponseBody {
	s.Header = v
	return s
}

func (s *GetCloudResourceDataResponseBody) SetRequestId(v string) *GetCloudResourceDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudResourceDataResponseBody) SetResponseStatus(v *GetCloudResourceDataResponseBodyResponseStatus) *GetCloudResourceDataResponseBody {
	s.ResponseStatus = v
	return s
}

func (s *GetCloudResourceDataResponseBody) Validate() error {
	if s.ResponseStatus != nil {
		if err := s.ResponseStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudResourceDataResponseBodyResponseStatus struct {
	// The information during execution.
	//
	// example:
	//
	// {}
	ExecutionStates *string `json:"executionStates,omitempty" xml:"executionStates,omitempty"`
	// The status level.
	//
	// example:
	//
	// Info,Warn,Error
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The execution result.
	//
	// example:
	//
	// Success,PartialSuccess,Error
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// The retry policy.
	//
	// example:
	//
	// None,Once,Continuous
	RetryPolicy *string `json:"retryPolicy,omitempty" xml:"retryPolicy,omitempty"`
	// The detailed status information.
	StatusItem []*GetCloudResourceDataResponseBodyResponseStatusStatusItem `json:"statusItem,omitempty" xml:"statusItem,omitempty" type:"Repeated"`
}

func (s GetCloudResourceDataResponseBodyResponseStatus) String() string {
	return dara.Prettify(s)
}

func (s GetCloudResourceDataResponseBodyResponseStatus) GoString() string {
	return s.String()
}

func (s *GetCloudResourceDataResponseBodyResponseStatus) GetExecutionStates() *string {
	return s.ExecutionStates
}

func (s *GetCloudResourceDataResponseBodyResponseStatus) GetLevel() *string {
	return s.Level
}

func (s *GetCloudResourceDataResponseBodyResponseStatus) GetResult() *string {
	return s.Result
}

func (s *GetCloudResourceDataResponseBodyResponseStatus) GetRetryPolicy() *string {
	return s.RetryPolicy
}

func (s *GetCloudResourceDataResponseBodyResponseStatus) GetStatusItem() []*GetCloudResourceDataResponseBodyResponseStatusStatusItem {
	return s.StatusItem
}

func (s *GetCloudResourceDataResponseBodyResponseStatus) SetExecutionStates(v string) *GetCloudResourceDataResponseBodyResponseStatus {
	s.ExecutionStates = &v
	return s
}

func (s *GetCloudResourceDataResponseBodyResponseStatus) SetLevel(v string) *GetCloudResourceDataResponseBodyResponseStatus {
	s.Level = &v
	return s
}

func (s *GetCloudResourceDataResponseBodyResponseStatus) SetResult(v string) *GetCloudResourceDataResponseBodyResponseStatus {
	s.Result = &v
	return s
}

func (s *GetCloudResourceDataResponseBodyResponseStatus) SetRetryPolicy(v string) *GetCloudResourceDataResponseBodyResponseStatus {
	s.RetryPolicy = &v
	return s
}

func (s *GetCloudResourceDataResponseBodyResponseStatus) SetStatusItem(v []*GetCloudResourceDataResponseBodyResponseStatusStatusItem) *GetCloudResourceDataResponseBodyResponseStatus {
	s.StatusItem = v
	return s
}

func (s *GetCloudResourceDataResponseBodyResponseStatus) Validate() error {
	if s.StatusItem != nil {
		for _, item := range s.StatusItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCloudResourceDataResponseBodyResponseStatusStatusItem struct {
	// The status code.
	//
	// example:
	//
	// Success,ExecuteTimeout,UModelNotExist
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status level.
	//
	// example:
	//
	// Info,Warn,Error
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The message content.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The suggestion when an execution error occurs.
	//
	// example:
	//
	// Try to reduce the query scope or increase timeout limit, then retry
	Suggestion *string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
}

func (s GetCloudResourceDataResponseBodyResponseStatusStatusItem) String() string {
	return dara.Prettify(s)
}

func (s GetCloudResourceDataResponseBodyResponseStatusStatusItem) GoString() string {
	return s.String()
}

func (s *GetCloudResourceDataResponseBodyResponseStatusStatusItem) GetCode() *string {
	return s.Code
}

func (s *GetCloudResourceDataResponseBodyResponseStatusStatusItem) GetLevel() *string {
	return s.Level
}

func (s *GetCloudResourceDataResponseBodyResponseStatusStatusItem) GetMessage() *string {
	return s.Message
}

func (s *GetCloudResourceDataResponseBodyResponseStatusStatusItem) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetCloudResourceDataResponseBodyResponseStatusStatusItem) SetCode(v string) *GetCloudResourceDataResponseBodyResponseStatusStatusItem {
	s.Code = &v
	return s
}

func (s *GetCloudResourceDataResponseBodyResponseStatusStatusItem) SetLevel(v string) *GetCloudResourceDataResponseBodyResponseStatusStatusItem {
	s.Level = &v
	return s
}

func (s *GetCloudResourceDataResponseBodyResponseStatusStatusItem) SetMessage(v string) *GetCloudResourceDataResponseBodyResponseStatusStatusItem {
	s.Message = &v
	return s
}

func (s *GetCloudResourceDataResponseBodyResponseStatusStatusItem) SetSuggestion(v string) *GetCloudResourceDataResponseBodyResponseStatusStatusItem {
	s.Suggestion = &v
	return s
}

func (s *GetCloudResourceDataResponseBodyResponseStatusStatusItem) Validate() error {
	return dara.Validate(s)
}
