// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFailoverTestJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFailoverTestJobResponseBody
	GetRequestId() *string
}

type DeleteFailoverTestJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C44F62BE-9CE7-4277-B117-69243F3988BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFailoverTestJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFailoverTestJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFailoverTestJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFailoverTestJobResponseBody) SetRequestId(v string) *DeleteFailoverTestJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFailoverTestJobResponseBody) Validate() error {
	return dara.Validate(s)
}
