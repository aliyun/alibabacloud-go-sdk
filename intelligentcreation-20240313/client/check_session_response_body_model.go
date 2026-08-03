// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckSessionResponseBody
	GetRequestId() *string
	SetStatus(v string) *CheckSessionResponseBody
	GetStatus() *string
}

type CheckSessionResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CheckSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckSessionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CheckSessionResponseBody) SetRequestId(v string) *CheckSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSessionResponseBody) SetStatus(v string) *CheckSessionResponseBody {
	s.Status = &v
	return s
}

func (s *CheckSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
