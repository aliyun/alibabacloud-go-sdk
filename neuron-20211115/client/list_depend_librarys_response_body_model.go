// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDependLibrarysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDependLibrarysResponseBody
	GetRequestId() *string
	SetResult(v *LibraryListResult) *ListDependLibrarysResponseBody
	GetResult() *LibraryListResult
}

type ListDependLibrarysResponseBody struct {
	RequestId *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *LibraryListResult `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListDependLibrarysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDependLibrarysResponseBody) GoString() string {
	return s.String()
}

func (s *ListDependLibrarysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDependLibrarysResponseBody) GetResult() *LibraryListResult {
	return s.Result
}

func (s *ListDependLibrarysResponseBody) SetRequestId(v string) *ListDependLibrarysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDependLibrarysResponseBody) SetResult(v *LibraryListResult) *ListDependLibrarysResponseBody {
	s.Result = v
	return s
}

func (s *ListDependLibrarysResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
