// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDeletionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelDeletionResponseBody
	GetRequestId() *string
	SetResult(v bool) *CancelDeletionResponseBody
	GetResult() *bool
}

type CancelDeletionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D682B6B3-B425-46DA-A5FC-5F5C60553622
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the instance is recovered. Valid values:
	//
	// - true: The instance is recovered.
	//
	// - false: The instance failed to be recovered.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CancelDeletionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelDeletionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDeletionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelDeletionResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CancelDeletionResponseBody) SetRequestId(v string) *CancelDeletionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDeletionResponseBody) SetResult(v bool) *CancelDeletionResponseBody {
	s.Result = &v
	return s
}

func (s *CancelDeletionResponseBody) Validate() error {
	return dara.Validate(s)
}
