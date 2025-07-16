// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateLiveResponseBody
	GetSuccess() *bool
}

type UpdateLiveResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateLiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateLiveResponseBody) SetRequestId(v string) *UpdateLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveResponseBody) SetSuccess(v bool) *UpdateLiveResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLiveResponseBody) Validate() error {
	return dara.Validate(s)
}
