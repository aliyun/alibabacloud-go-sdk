// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigSequenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateConfigSequenceResponseBody
	GetRequestId() *string
}

type UpdateConfigSequenceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigSequenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigSequenceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigSequenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConfigSequenceResponseBody) SetRequestId(v string) *UpdateConfigSequenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigSequenceResponseBody) Validate() error {
	return dara.Validate(s)
}
