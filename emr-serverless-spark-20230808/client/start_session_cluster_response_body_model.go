// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSessionClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartSessionClusterResponseBody
	GetRequestId() *string
	SetSessionClusterId(v string) *StartSessionClusterResponseBody
	GetSessionClusterId() *string
}

type StartSessionClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-******
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
}

func (s StartSessionClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StartSessionClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSessionClusterResponseBody) GetSessionClusterId() *string {
	return s.SessionClusterId
}

func (s *StartSessionClusterResponseBody) SetRequestId(v string) *StartSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSessionClusterResponseBody) SetSessionClusterId(v string) *StartSessionClusterResponseBody {
	s.SessionClusterId = &v
	return s
}

func (s *StartSessionClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
