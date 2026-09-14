// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpdateTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUpdateTaskResultResponseBody
	GetRequestId() *string
	SetResult(v *GetUpdateTaskResultResponseBodyResult) *GetUpdateTaskResultResponseBody
	GetResult() *GetUpdateTaskResultResponseBodyResult
}

type GetUpdateTaskResultResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF020E7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The node update result.
	Result *GetUpdateTaskResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetUpdateTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUpdateTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetUpdateTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUpdateTaskResultResponseBody) GetResult() *GetUpdateTaskResultResponseBodyResult {
	return s.Result
}

func (s *GetUpdateTaskResultResponseBody) SetRequestId(v string) *GetUpdateTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUpdateTaskResultResponseBody) SetResult(v *GetUpdateTaskResultResponseBodyResult) *GetUpdateTaskResultResponseBody {
	s.Result = v
	return s
}

func (s *GetUpdateTaskResultResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUpdateTaskResultResponseBodyResult struct {
	// The failure message. This field is returned if the update fails.
	//
	// example:
	//
	// Invalid Param xxx
	FailureMessage *string `json:"FailureMessage,omitempty" xml:"FailureMessage,omitempty"`
	// The update status. Valid values:
	//
	// - Updating: The node is being updated.
	//
	// - Updated: The node is updated.
	//
	// - UpdateFailed: The node failed to be updated.
	//
	// example:
	//
	// Updated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetUpdateTaskResultResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetUpdateTaskResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUpdateTaskResultResponseBodyResult) GetFailureMessage() *string {
	return s.FailureMessage
}

func (s *GetUpdateTaskResultResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetUpdateTaskResultResponseBodyResult) SetFailureMessage(v string) *GetUpdateTaskResultResponseBodyResult {
	s.FailureMessage = &v
	return s
}

func (s *GetUpdateTaskResultResponseBodyResult) SetStatus(v string) *GetUpdateTaskResultResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetUpdateTaskResultResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
