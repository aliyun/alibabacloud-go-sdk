// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAScriptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DeleteAScriptsResponseBody
	GetJobId() *string
	SetRequestId(v string) *DeleteAScriptsResponseBody
	GetRequestId() *string
}

type DeleteAScriptsResponseBody struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// 03cf3fe1-ab37-479b-92a6-b481d762****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1B7B7695-3596-50C8-B739-030C6C685E61
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAScriptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAScriptsResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *DeleteAScriptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAScriptsResponseBody) SetJobId(v string) *DeleteAScriptsResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteAScriptsResponseBody) SetRequestId(v string) *DeleteAScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAScriptsResponseBody) Validate() error {
	return dara.Validate(s)
}
