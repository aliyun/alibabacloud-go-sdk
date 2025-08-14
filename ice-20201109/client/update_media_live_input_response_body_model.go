// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLiveInputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMediaLiveInputResponseBody
	GetRequestId() *string
}

type UpdateMediaLiveInputResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaLiveInputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveInputResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveInputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaLiveInputResponseBody) SetRequestId(v string) *UpdateMediaLiveInputResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaLiveInputResponseBody) Validate() error {
	return dara.Validate(s)
}
