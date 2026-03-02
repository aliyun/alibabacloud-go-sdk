// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLibrarysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListLibrarysResponseBody
	GetRequestId() *string
	SetResult(v *LibraryListResult) *ListLibrarysResponseBody
	GetResult() *LibraryListResult
}

type ListLibrarysResponseBody struct {
	RequestId *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *LibraryListResult `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListLibrarysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLibrarysResponseBody) GoString() string {
	return s.String()
}

func (s *ListLibrarysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLibrarysResponseBody) GetResult() *LibraryListResult {
	return s.Result
}

func (s *ListLibrarysResponseBody) SetRequestId(v string) *ListLibrarysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLibrarysResponseBody) SetResult(v *LibraryListResult) *ListLibrarysResponseBody {
	s.Result = v
	return s
}

func (s *ListLibrarysResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
