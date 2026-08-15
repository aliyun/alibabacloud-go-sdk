// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRunId(v string) *CreateRunResponseBody
	GetRunId() *string
	SetRequestId(v string) *CreateRunResponseBody
	GetRequestId() *string
}

type CreateRunResponseBody struct {
	// The run ID.
	//
	// example:
	//
	// run-1meoz7VJd2C6f****
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *CreateRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRunResponseBody) SetRunId(v string) *CreateRunResponseBody {
	s.RunId = &v
	return s
}

func (s *CreateRunResponseBody) SetRequestId(v string) *CreateRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRunResponseBody) Validate() error {
	return dara.Validate(s)
}
