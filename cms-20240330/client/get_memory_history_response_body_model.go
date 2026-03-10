// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMemoryHistoryResponseBody
	GetRequestId() *string
	SetResults(v []*GetMemoryHistoryResponseBodyResults) *GetMemoryHistoryResponseBody
	GetResults() []*GetMemoryHistoryResponseBodyResults
}

type GetMemoryHistoryResponseBody struct {
	// example:
	//
	// 3B311FD9-A60B-55E0-A896-A0C73*********
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Results   []*GetMemoryHistoryResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s GetMemoryHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryHistoryResponseBody) GetResults() []*GetMemoryHistoryResponseBodyResults {
	return s.Results
}

func (s *GetMemoryHistoryResponseBody) SetRequestId(v string) *GetMemoryHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryHistoryResponseBody) SetResults(v []*GetMemoryHistoryResponseBodyResults) *GetMemoryHistoryResponseBody {
	s.Results = v
	return s
}

func (s *GetMemoryHistoryResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMemoryHistoryResponseBodyResults struct {
	// example:
	//
	// 1764556182850
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// ADD
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// b25d6ad7-306f-4040-9890-4dddd2505a2e
	Id    *string                                     `json:"id,omitempty" xml:"id,omitempty"`
	Input []*GetMemoryHistoryResponseBodyResultsInput `json:"input,omitempty" xml:"input,omitempty" type:"Repeated"`
	// example:
	//
	// 019cacf6-7b39-7f61-8314-548f07ca449a
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
	// example:
	//
	// {"sessionId":"test_session_001"}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// I really like Python.
	NewMemory *string `json:"newMemory,omitempty" xml:"newMemory,omitempty"`
	// example:
	//
	// I really don\\"t like Python at all.
	OldMemory *string `json:"oldMemory,omitempty" xml:"oldMemory,omitempty"`
	// example:
	//
	// 1771036123785
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// test_session_001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetMemoryHistoryResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryHistoryResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetMemoryHistoryResponseBodyResults) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetMemoryHistoryResponseBodyResults) GetEvent() *string {
	return s.Event
}

func (s *GetMemoryHistoryResponseBodyResults) GetId() *string {
	return s.Id
}

func (s *GetMemoryHistoryResponseBodyResults) GetInput() []*GetMemoryHistoryResponseBodyResultsInput {
	return s.Input
}

func (s *GetMemoryHistoryResponseBodyResults) GetMemoryId() *string {
	return s.MemoryId
}

func (s *GetMemoryHistoryResponseBodyResults) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *GetMemoryHistoryResponseBodyResults) GetNewMemory() *string {
	return s.NewMemory
}

func (s *GetMemoryHistoryResponseBodyResults) GetOldMemory() *string {
	return s.OldMemory
}

func (s *GetMemoryHistoryResponseBodyResults) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetMemoryHistoryResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *GetMemoryHistoryResponseBodyResults) SetCreatedAt(v string) *GetMemoryHistoryResponseBodyResults {
	s.CreatedAt = &v
	return s
}

func (s *GetMemoryHistoryResponseBodyResults) SetEvent(v string) *GetMemoryHistoryResponseBodyResults {
	s.Event = &v
	return s
}

func (s *GetMemoryHistoryResponseBodyResults) SetId(v string) *GetMemoryHistoryResponseBodyResults {
	s.Id = &v
	return s
}

func (s *GetMemoryHistoryResponseBodyResults) SetInput(v []*GetMemoryHistoryResponseBodyResultsInput) *GetMemoryHistoryResponseBodyResults {
	s.Input = v
	return s
}

func (s *GetMemoryHistoryResponseBodyResults) SetMemoryId(v string) *GetMemoryHistoryResponseBodyResults {
	s.MemoryId = &v
	return s
}

func (s *GetMemoryHistoryResponseBodyResults) SetMetadata(v map[string]interface{}) *GetMemoryHistoryResponseBodyResults {
	s.Metadata = v
	return s
}

func (s *GetMemoryHistoryResponseBodyResults) SetNewMemory(v string) *GetMemoryHistoryResponseBodyResults {
	s.NewMemory = &v
	return s
}

func (s *GetMemoryHistoryResponseBodyResults) SetOldMemory(v string) *GetMemoryHistoryResponseBodyResults {
	s.OldMemory = &v
	return s
}

func (s *GetMemoryHistoryResponseBodyResults) SetUpdatedAt(v string) *GetMemoryHistoryResponseBodyResults {
	s.UpdatedAt = &v
	return s
}

func (s *GetMemoryHistoryResponseBodyResults) SetUserId(v string) *GetMemoryHistoryResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *GetMemoryHistoryResponseBodyResults) Validate() error {
	if s.Input != nil {
		for _, item := range s.Input {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMemoryHistoryResponseBodyResultsInput struct {
	// example:
	//
	// My name is Zhang San and I live in Hangzhou.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetMemoryHistoryResponseBodyResultsInput) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryHistoryResponseBodyResultsInput) GoString() string {
	return s.String()
}

func (s *GetMemoryHistoryResponseBodyResultsInput) GetContent() *string {
	return s.Content
}

func (s *GetMemoryHistoryResponseBodyResultsInput) GetRole() *string {
	return s.Role
}

func (s *GetMemoryHistoryResponseBodyResultsInput) SetContent(v string) *GetMemoryHistoryResponseBodyResultsInput {
	s.Content = &v
	return s
}

func (s *GetMemoryHistoryResponseBodyResultsInput) SetRole(v string) *GetMemoryHistoryResponseBodyResultsInput {
	s.Role = &v
	return s
}

func (s *GetMemoryHistoryResponseBodyResultsInput) Validate() error {
	return dara.Validate(s)
}
