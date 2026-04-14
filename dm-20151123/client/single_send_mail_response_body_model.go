// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleSendMailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnvId(v string) *SingleSendMailResponseBody
	GetEnvId() *string
	SetRequestId(v string) *SingleSendMailResponseBody
	GetRequestId() *string
}

type SingleSendMailResponseBody struct {
	// Event ID
	//
	// example:
	//
	// 600000xxxxxxxxxx642
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	// Request ID
	//
	// example:
	//
	// 2D086F6-xxxx-xxxx-xxxx-006DED011A85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SingleSendMailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailResponseBody) GoString() string {
	return s.String()
}

func (s *SingleSendMailResponseBody) GetEnvId() *string {
	return s.EnvId
}

func (s *SingleSendMailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SingleSendMailResponseBody) SetEnvId(v string) *SingleSendMailResponseBody {
	s.EnvId = &v
	return s
}

func (s *SingleSendMailResponseBody) SetRequestId(v string) *SingleSendMailResponseBody {
	s.RequestId = &v
	return s
}

func (s *SingleSendMailResponseBody) Validate() error {
	return dara.Validate(s)
}
