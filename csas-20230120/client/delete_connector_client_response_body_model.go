// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectorClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConnectorClientResponseBody
	GetRequestId() *string
}

type DeleteConnectorClientResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FEF5CFA-14CC-5DE5-BD1F-AFFE0996E71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConnectorClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectorClientResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnectorClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConnectorClientResponseBody) SetRequestId(v string) *DeleteConnectorClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConnectorClientResponseBody) Validate() error {
	return dara.Validate(s)
}
