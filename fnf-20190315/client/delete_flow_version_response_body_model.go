// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFlowVersionResponseBody
	GetRequestId() *string
}

type DeleteFlowVersionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 3A44E113-9962-5B0B-AB92-14060EFE3164
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlowVersionResponseBody) SetRequestId(v string) *DeleteFlowVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
