// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSearchStrategyResponseBody
	GetRequestId() *string
}

type CreateSearchStrategyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// "abc123"
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateSearchStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSearchStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSearchStrategyResponseBody) SetRequestId(v string) *CreateSearchStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSearchStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
