// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckYikeAIAppJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PrecheckYikeAIAppJobResponseBody
	GetRequestId() *string
	SetResult(v []*PrecheckYikeAIAppJobResponseBodyResult) *PrecheckYikeAIAppJobResponseBody
	GetResult() []*PrecheckYikeAIAppJobResponseBodyResult
	SetStatus(v string) *PrecheckYikeAIAppJobResponseBody
	GetStatus() *string
}

type PrecheckYikeAIAppJobResponseBody struct {
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*PrecheckYikeAIAppJobResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PrecheckYikeAIAppJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PrecheckYikeAIAppJobResponseBody) GoString() string {
	return s.String()
}

func (s *PrecheckYikeAIAppJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PrecheckYikeAIAppJobResponseBody) GetResult() []*PrecheckYikeAIAppJobResponseBodyResult {
	return s.Result
}

func (s *PrecheckYikeAIAppJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *PrecheckYikeAIAppJobResponseBody) SetRequestId(v string) *PrecheckYikeAIAppJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *PrecheckYikeAIAppJobResponseBody) SetResult(v []*PrecheckYikeAIAppJobResponseBodyResult) *PrecheckYikeAIAppJobResponseBody {
	s.Result = v
	return s
}

func (s *PrecheckYikeAIAppJobResponseBody) SetStatus(v string) *PrecheckYikeAIAppJobResponseBody {
	s.Status = &v
	return s
}

func (s *PrecheckYikeAIAppJobResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PrecheckYikeAIAppJobResponseBodyResult struct {
	// example:
	//
	// ImageCheckFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// LoadImage.1.TargetImage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s PrecheckYikeAIAppJobResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s PrecheckYikeAIAppJobResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PrecheckYikeAIAppJobResponseBodyResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PrecheckYikeAIAppJobResponseBodyResult) GetKey() *string {
	return s.Key
}

func (s *PrecheckYikeAIAppJobResponseBodyResult) SetErrorCode(v string) *PrecheckYikeAIAppJobResponseBodyResult {
	s.ErrorCode = &v
	return s
}

func (s *PrecheckYikeAIAppJobResponseBodyResult) SetKey(v string) *PrecheckYikeAIAppJobResponseBodyResult {
	s.Key = &v
	return s
}

func (s *PrecheckYikeAIAppJobResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
