// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEntityStoreDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v [][]*string) *GetEntityStoreDataResponseBody
	GetData() [][]*string
	SetHeader(v []*string) *GetEntityStoreDataResponseBody
	GetHeader() []*string
	SetRequestId(v string) *GetEntityStoreDataResponseBody
	GetRequestId() *string
	SetResponseStatus(v *GetEntityStoreDataResponseBodyResponseStatus) *GetEntityStoreDataResponseBody
	GetResponseStatus() *GetEntityStoreDataResponseBodyResponseStatus
}

type GetEntityStoreDataResponseBody struct {
	Data   [][]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Header []*string   `json:"header,omitempty" xml:"header,omitempty" type:"Repeated"`
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId      *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResponseStatus *GetEntityStoreDataResponseBodyResponseStatus `json:"responseStatus,omitempty" xml:"responseStatus,omitempty" type:"Struct"`
}

func (s GetEntityStoreDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEntityStoreDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityStoreDataResponseBody) GetData() [][]*string {
	return s.Data
}

func (s *GetEntityStoreDataResponseBody) GetHeader() []*string {
	return s.Header
}

func (s *GetEntityStoreDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEntityStoreDataResponseBody) GetResponseStatus() *GetEntityStoreDataResponseBodyResponseStatus {
	return s.ResponseStatus
}

func (s *GetEntityStoreDataResponseBody) SetData(v [][]*string) *GetEntityStoreDataResponseBody {
	s.Data = v
	return s
}

func (s *GetEntityStoreDataResponseBody) SetHeader(v []*string) *GetEntityStoreDataResponseBody {
	s.Header = v
	return s
}

func (s *GetEntityStoreDataResponseBody) SetRequestId(v string) *GetEntityStoreDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntityStoreDataResponseBody) SetResponseStatus(v *GetEntityStoreDataResponseBodyResponseStatus) *GetEntityStoreDataResponseBody {
	s.ResponseStatus = v
	return s
}

func (s *GetEntityStoreDataResponseBody) Validate() error {
	if s.ResponseStatus != nil {
		if err := s.ResponseStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEntityStoreDataResponseBodyResponseStatus struct {
	ExecutionStates *string                                                   `json:"executionStates,omitempty" xml:"executionStates,omitempty"`
	Level           *string                                                   `json:"level,omitempty" xml:"level,omitempty"`
	Result          *string                                                   `json:"result,omitempty" xml:"result,omitempty"`
	RetryPolicy     *string                                                   `json:"retryPolicy,omitempty" xml:"retryPolicy,omitempty"`
	StatusItem      []*GetEntityStoreDataResponseBodyResponseStatusStatusItem `json:"statusItem,omitempty" xml:"statusItem,omitempty" type:"Repeated"`
}

func (s GetEntityStoreDataResponseBodyResponseStatus) String() string {
	return dara.Prettify(s)
}

func (s GetEntityStoreDataResponseBodyResponseStatus) GoString() string {
	return s.String()
}

func (s *GetEntityStoreDataResponseBodyResponseStatus) GetExecutionStates() *string {
	return s.ExecutionStates
}

func (s *GetEntityStoreDataResponseBodyResponseStatus) GetLevel() *string {
	return s.Level
}

func (s *GetEntityStoreDataResponseBodyResponseStatus) GetResult() *string {
	return s.Result
}

func (s *GetEntityStoreDataResponseBodyResponseStatus) GetRetryPolicy() *string {
	return s.RetryPolicy
}

func (s *GetEntityStoreDataResponseBodyResponseStatus) GetStatusItem() []*GetEntityStoreDataResponseBodyResponseStatusStatusItem {
	return s.StatusItem
}

func (s *GetEntityStoreDataResponseBodyResponseStatus) SetExecutionStates(v string) *GetEntityStoreDataResponseBodyResponseStatus {
	s.ExecutionStates = &v
	return s
}

func (s *GetEntityStoreDataResponseBodyResponseStatus) SetLevel(v string) *GetEntityStoreDataResponseBodyResponseStatus {
	s.Level = &v
	return s
}

func (s *GetEntityStoreDataResponseBodyResponseStatus) SetResult(v string) *GetEntityStoreDataResponseBodyResponseStatus {
	s.Result = &v
	return s
}

func (s *GetEntityStoreDataResponseBodyResponseStatus) SetRetryPolicy(v string) *GetEntityStoreDataResponseBodyResponseStatus {
	s.RetryPolicy = &v
	return s
}

func (s *GetEntityStoreDataResponseBodyResponseStatus) SetStatusItem(v []*GetEntityStoreDataResponseBodyResponseStatusStatusItem) *GetEntityStoreDataResponseBodyResponseStatus {
	s.StatusItem = v
	return s
}

func (s *GetEntityStoreDataResponseBodyResponseStatus) Validate() error {
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

type GetEntityStoreDataResponseBodyResponseStatusStatusItem struct {
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	Level      *string `json:"level,omitempty" xml:"level,omitempty"`
	Message    *string `json:"message,omitempty" xml:"message,omitempty"`
	Suggestion *string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
}

func (s GetEntityStoreDataResponseBodyResponseStatusStatusItem) String() string {
	return dara.Prettify(s)
}

func (s GetEntityStoreDataResponseBodyResponseStatusStatusItem) GoString() string {
	return s.String()
}

func (s *GetEntityStoreDataResponseBodyResponseStatusStatusItem) GetCode() *string {
	return s.Code
}

func (s *GetEntityStoreDataResponseBodyResponseStatusStatusItem) GetLevel() *string {
	return s.Level
}

func (s *GetEntityStoreDataResponseBodyResponseStatusStatusItem) GetMessage() *string {
	return s.Message
}

func (s *GetEntityStoreDataResponseBodyResponseStatusStatusItem) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetEntityStoreDataResponseBodyResponseStatusStatusItem) SetCode(v string) *GetEntityStoreDataResponseBodyResponseStatusStatusItem {
	s.Code = &v
	return s
}

func (s *GetEntityStoreDataResponseBodyResponseStatusStatusItem) SetLevel(v string) *GetEntityStoreDataResponseBodyResponseStatusStatusItem {
	s.Level = &v
	return s
}

func (s *GetEntityStoreDataResponseBodyResponseStatusStatusItem) SetMessage(v string) *GetEntityStoreDataResponseBodyResponseStatusStatusItem {
	s.Message = &v
	return s
}

func (s *GetEntityStoreDataResponseBodyResponseStatusStatusItem) SetSuggestion(v string) *GetEntityStoreDataResponseBodyResponseStatusStatusItem {
	s.Suggestion = &v
	return s
}

func (s *GetEntityStoreDataResponseBodyResponseStatusStatusItem) Validate() error {
	return dara.Validate(s)
}
