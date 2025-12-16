// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSearchStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSearchStrategyResponseBody
	GetRequestId() *string
}

type GetSearchStrategyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5C1C1C45-C64A-AD30-565F-140871D57E5E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSearchStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSearchStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetSearchStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSearchStrategyResponseBody) SetRequestId(v string) *GetSearchStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSearchStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
