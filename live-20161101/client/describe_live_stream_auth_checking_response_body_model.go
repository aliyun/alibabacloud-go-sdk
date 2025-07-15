// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamAuthCheckingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeLiveStreamAuthCheckingResponseBody
	GetDescription() *string
	SetRequestId(v string) *DescribeLiveStreamAuthCheckingResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeLiveStreamAuthCheckingResponseBody
	GetStatus() *string
}

type DescribeLiveStreamAuthCheckingResponseBody struct {
	// The error message for failed authentication.
	//
	// example:
	//
	// INVALID AUTH_KEY
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16F08B4E-14FD-5D72-AB2F-BAFA4C4D57F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the signed URL is valid. Valid values:
	//
	// 	- pass: valid.
	//
	// 	- false: invalid.
	//
	// example:
	//
	// pass
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeLiveStreamAuthCheckingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamAuthCheckingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamAuthCheckingResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeLiveStreamAuthCheckingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamAuthCheckingResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLiveStreamAuthCheckingResponseBody) SetDescription(v string) *DescribeLiveStreamAuthCheckingResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLiveStreamAuthCheckingResponseBody) SetRequestId(v string) *DescribeLiveStreamAuthCheckingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamAuthCheckingResponseBody) SetStatus(v string) *DescribeLiveStreamAuthCheckingResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLiveStreamAuthCheckingResponseBody) Validate() error {
	return dara.Validate(s)
}
