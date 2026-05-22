// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSitePauseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSitePauseResponseBody
	GetRequestId() *string
}

type UpdateSitePauseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSitePauseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSitePauseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSitePauseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSitePauseResponseBody) SetRequestId(v string) *UpdateSitePauseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSitePauseResponseBody) Validate() error {
	return dara.Validate(s)
}
