// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFlowAliasResponseBody
	GetRequestId() *string
}

type DeleteFlowAliasResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 3A44E113-9962-5B0B-AB92-14060EFE3164
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowAliasResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlowAliasResponseBody) SetRequestId(v string) *DeleteFlowAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
