// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAScriptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateAScriptsResponseBody
	GetJobId() *string
	SetRequestId(v string) *UpdateAScriptsResponseBody
	GetRequestId() *string
}

type UpdateAScriptsResponseBody struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// 2e82b5f4-1ba9-4d20-89c8-1082ebaa****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ACA19FE1-C09E-53C7-8FDA-560F49D71891
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAScriptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAScriptsResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpdateAScriptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAScriptsResponseBody) SetJobId(v string) *UpdateAScriptsResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateAScriptsResponseBody) SetRequestId(v string) *UpdateAScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAScriptsResponseBody) Validate() error {
	return dara.Validate(s)
}
