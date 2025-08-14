// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProgramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProgramResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteProgramResponseBody
	GetSuccess() *bool
}

type DeleteProgramResponseBody struct {
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteProgramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProgramResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProgramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProgramResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteProgramResponseBody) SetRequestId(v string) *DeleteProgramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProgramResponseBody) SetSuccess(v bool) *DeleteProgramResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteProgramResponseBody) Validate() error {
	return dara.Validate(s)
}
