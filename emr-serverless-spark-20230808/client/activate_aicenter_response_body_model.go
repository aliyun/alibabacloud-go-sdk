// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateAICenterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ActivateAICenterResponseBody
	GetRequestId() *string
}

type ActivateAICenterResponseBody struct {
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ActivateAICenterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateAICenterResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateAICenterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateAICenterResponseBody) SetRequestId(v string) *ActivateAICenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateAICenterResponseBody) Validate() error {
	return dara.Validate(s)
}
