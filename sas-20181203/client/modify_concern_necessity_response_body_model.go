// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConcernNecessityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyConcernNecessityResponseBody
	GetRequestId() *string
}

type ModifyConcernNecessityResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F059E190-A65B-5DF8-8709-2CC7791A5B65
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyConcernNecessityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyConcernNecessityResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConcernNecessityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyConcernNecessityResponseBody) SetRequestId(v string) *ModifyConcernNecessityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConcernNecessityResponseBody) Validate() error {
	return dara.Validate(s)
}
