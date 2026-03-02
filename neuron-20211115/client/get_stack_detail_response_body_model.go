// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetStackDetailResponseBody
	GetRequestId() *string
	SetResult(v *StackDetailResult) *GetStackDetailResponseBody
	GetResult() *StackDetailResult
}

type GetStackDetailResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// sdadawqewe
	RequestId *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *StackDetailResult `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetStackDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackDetailResponseBody) GetResult() *StackDetailResult {
	return s.Result
}

func (s *GetStackDetailResponseBody) SetRequestId(v string) *GetStackDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackDetailResponseBody) SetResult(v *StackDetailResult) *GetStackDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetStackDetailResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
