// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DebugModelResponseBody
	GetRequestId() *string
}

type DebugModelResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DebugModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DebugModelResponseBody) GoString() string {
	return s.String()
}

func (s *DebugModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DebugModelResponseBody) SetRequestId(v string) *DebugModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugModelResponseBody) Validate() error {
	return dara.Validate(s)
}
