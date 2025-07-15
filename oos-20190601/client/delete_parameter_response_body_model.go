// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteParameterResponseBody
	GetRequestId() *string
}

type DeleteParameterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0D02BDF-77F6-49F2-95C9-8E87121D2979
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteParameterResponseBody) SetRequestId(v string) *DeleteParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteParameterResponseBody) Validate() error {
	return dara.Validate(s)
}
