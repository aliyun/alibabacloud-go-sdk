// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseSortScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseSortScriptResponseBody
	GetRequestId() *string
	SetResult(v *ReleaseSortScriptResponseBodyResult) *ReleaseSortScriptResponseBody
	GetResult() *ReleaseSortScriptResponseBodyResult
}

type ReleaseSortScriptResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ReleaseSortScriptResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ReleaseSortScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseSortScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseSortScriptResponseBody) GetResult() *ReleaseSortScriptResponseBodyResult {
	return s.Result
}

func (s *ReleaseSortScriptResponseBody) SetRequestId(v string) *ReleaseSortScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseSortScriptResponseBody) SetResult(v *ReleaseSortScriptResponseBodyResult) *ReleaseSortScriptResponseBody {
	s.Result = v
	return s
}

func (s *ReleaseSortScriptResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReleaseSortScriptResponseBodyResult struct {
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ReleaseSortScriptResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSortScriptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ReleaseSortScriptResponseBodyResult) GetVersion() *int64 {
	return s.Version
}

func (s *ReleaseSortScriptResponseBodyResult) SetVersion(v int64) *ReleaseSortScriptResponseBodyResult {
	s.Version = &v
	return s
}

func (s *ReleaseSortScriptResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
