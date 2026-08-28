// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAgentSpecVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SubmitAgentSpecVersionResponseBody
	GetData() *string
	SetRequestId(v string) *SubmitAgentSpecVersionResponseBody
	GetRequestId() *string
}

type SubmitAgentSpecVersionResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// 1.0.0
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SubmitAgentSpecVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAgentSpecVersionResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAgentSpecVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitAgentSpecVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAgentSpecVersionResponseBody) SetData(v string) *SubmitAgentSpecVersionResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitAgentSpecVersionResponseBody) SetRequestId(v string) *SubmitAgentSpecVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAgentSpecVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
