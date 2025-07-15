// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartFailoverTestJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartFailoverTestJobResponseBody
	GetRequestId() *string
}

type StartFailoverTestJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C44F62BE-9CE7-4277-B117-69243F3988BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartFailoverTestJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartFailoverTestJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartFailoverTestJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartFailoverTestJobResponseBody) SetRequestId(v string) *StartFailoverTestJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartFailoverTestJobResponseBody) Validate() error {
	return dara.Validate(s)
}
