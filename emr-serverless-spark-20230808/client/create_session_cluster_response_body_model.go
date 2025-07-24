// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSessionClusterResponseBody
	GetRequestId() *string
	SetSessionClusterId(v string) *CreateSessionClusterResponseBody
	GetSessionClusterId() *string
}

type CreateSessionClusterResponseBody struct {
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

func (s CreateSessionClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSessionClusterResponseBody) GetSessionClusterId() *string {
	return s.SessionClusterId
}

func (s *CreateSessionClusterResponseBody) SetRequestId(v string) *CreateSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSessionClusterResponseBody) SetSessionClusterId(v string) *CreateSessionClusterResponseBody {
	s.SessionClusterId = &v
	return s
}

func (s *CreateSessionClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
