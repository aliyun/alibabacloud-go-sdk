// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApisecEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApisecEventsResponseBody
	GetRequestId() *string
}

type DeleteApisecEventsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApisecEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApisecEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApisecEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApisecEventsResponseBody) SetRequestId(v string) *DeleteApisecEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApisecEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
