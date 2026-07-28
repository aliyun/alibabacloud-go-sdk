// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateResourceControlResponseBody
	GetRequestId() *string
}

type CreateResourceControlResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// C61892A4-0850-4516-9E26-44D96C1782DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceControlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceControlResponseBody) SetRequestId(v string) *CreateResourceControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceControlResponseBody) Validate() error {
	return dara.Validate(s)
}
