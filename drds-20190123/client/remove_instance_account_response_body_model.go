// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveInstanceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveInstanceAccountResponseBody
	GetSuccess() *bool
}

type RemoveInstanceAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A3140FC7-B78B-4D8E-B0C8-926D28******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveInstanceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInstanceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveInstanceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveInstanceAccountResponseBody) SetRequestId(v string) *RemoveInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInstanceAccountResponseBody) SetSuccess(v bool) *RemoveInstanceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveInstanceAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
