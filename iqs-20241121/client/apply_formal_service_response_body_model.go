// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyFormalServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyFormalServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyFormalServiceResponseBody
	GetSuccess() *bool
}

type ApplyFormalServiceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ApplyFormalServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyFormalServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyFormalServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyFormalServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyFormalServiceResponseBody) SetRequestId(v string) *ApplyFormalServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyFormalServiceResponseBody) SetSuccess(v bool) *ApplyFormalServiceResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyFormalServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
