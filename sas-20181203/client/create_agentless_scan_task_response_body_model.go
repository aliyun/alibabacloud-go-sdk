// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentlessScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAgentlessScanTaskResponseBody
	GetRequestId() *string
}

type CreateAgentlessScanTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E90DE229-9FC6-58F6-BF4B-03AD6179****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAgentlessScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentlessScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentlessScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentlessScanTaskResponseBody) SetRequestId(v string) *CreateAgentlessScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentlessScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
