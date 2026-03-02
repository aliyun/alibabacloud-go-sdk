// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcInvokesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPbcInvokesResponseBody
	GetRequestId() *string
	SetResult(v *PbcListResult) *ListPbcInvokesResponseBody
	GetResult() *PbcListResult
}

type ListPbcInvokesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// sdadawqewe
	RequestId *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *PbcListResult `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListPbcInvokesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPbcInvokesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPbcInvokesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPbcInvokesResponseBody) GetResult() *PbcListResult {
	return s.Result
}

func (s *ListPbcInvokesResponseBody) SetRequestId(v string) *ListPbcInvokesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPbcInvokesResponseBody) SetResult(v *PbcListResult) *ListPbcInvokesResponseBody {
	s.Result = v
	return s
}

func (s *ListPbcInvokesResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
