// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PauseClientResponseBody
	GetRequestId() *string
}

type PauseClientResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6673D49C-A9AB-40DD-B4A2-B92306701AE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PauseClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseClientResponseBody) GoString() string {
	return s.String()
}

func (s *PauseClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseClientResponseBody) SetRequestId(v string) *PauseClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseClientResponseBody) Validate() error {
	return dara.Validate(s)
}
