// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcReviewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPbcReviewsResponseBody
	GetRequestId() *string
	SetResult(v *PbcReviewListResult) *ListPbcReviewsResponseBody
	GetResult() *PbcReviewListResult
}

type ListPbcReviewsResponseBody struct {
	RequestId *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *PbcReviewListResult `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListPbcReviewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPbcReviewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPbcReviewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPbcReviewsResponseBody) GetResult() *PbcReviewListResult {
	return s.Result
}

func (s *ListPbcReviewsResponseBody) SetRequestId(v string) *ListPbcReviewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPbcReviewsResponseBody) SetResult(v *PbcReviewListResult) *ListPbcReviewsResponseBody {
	s.Result = v
	return s
}

func (s *ListPbcReviewsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
