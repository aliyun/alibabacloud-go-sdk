// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSessionClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopSessionClusterResponseBody
	GetRequestId() *string
	SetSessionClusterId(v string) *StopSessionClusterResponseBody
	GetSessionClusterId() *string
}

type StopSessionClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// w-******
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
}

func (s StopSessionClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopSessionClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopSessionClusterResponseBody) GetSessionClusterId() *string {
	return s.SessionClusterId
}

func (s *StopSessionClusterResponseBody) SetRequestId(v string) *StopSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSessionClusterResponseBody) SetSessionClusterId(v string) *StopSessionClusterResponseBody {
	s.SessionClusterId = &v
	return s
}

func (s *StopSessionClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
