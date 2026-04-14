// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSendMailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnvId(v string) *BatchSendMailResponseBody
	GetEnvId() *string
	SetRequestId(v string) *BatchSendMailResponseBody
	GetRequestId() *string
}

type BatchSendMailResponseBody struct {
	// Event ID
	//
	// example:
	//
	// xxx
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	// Request ID
	//
	// example:
	//
	// 12D086F6-8F31-4658-84C1-006DED011A85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSendMailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMailResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendMailResponseBody) GetEnvId() *string {
	return s.EnvId
}

func (s *BatchSendMailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSendMailResponseBody) SetEnvId(v string) *BatchSendMailResponseBody {
	s.EnvId = &v
	return s
}

func (s *BatchSendMailResponseBody) SetRequestId(v string) *BatchSendMailResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSendMailResponseBody) Validate() error {
	return dara.Validate(s)
}
