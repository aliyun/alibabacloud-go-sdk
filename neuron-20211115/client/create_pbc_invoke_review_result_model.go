// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcInvokeReviewResult interface {
	dara.Model
	String() string
	GoString() string
	SetPbcInvokeReviewId(v int64) *CreatePbcInvokeReviewResult
	GetPbcInvokeReviewId() *int64
	SetRequestId(v string) *CreatePbcInvokeReviewResult
	GetRequestId() *string
}

type CreatePbcInvokeReviewResult struct {
	PbcInvokeReviewId *int64  `json:"pbcInvokeReviewId,omitempty" xml:"pbcInvokeReviewId,omitempty"`
	RequestId         *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePbcInvokeReviewResult) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcInvokeReviewResult) GoString() string {
	return s.String()
}

func (s *CreatePbcInvokeReviewResult) GetPbcInvokeReviewId() *int64 {
	return s.PbcInvokeReviewId
}

func (s *CreatePbcInvokeReviewResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePbcInvokeReviewResult) SetPbcInvokeReviewId(v int64) *CreatePbcInvokeReviewResult {
	s.PbcInvokeReviewId = &v
	return s
}

func (s *CreatePbcInvokeReviewResult) SetRequestId(v string) *CreatePbcInvokeReviewResult {
	s.RequestId = &v
	return s
}

func (s *CreatePbcInvokeReviewResult) Validate() error {
	return dara.Validate(s)
}
