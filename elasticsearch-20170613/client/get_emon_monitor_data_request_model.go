// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *GetEmonMonitorDataRequest
	GetBody() *string
}

type GetEmonMonitorDataRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmonMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEmonMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *GetEmonMonitorDataRequest) GetBody() *string {
	return s.Body
}

func (s *GetEmonMonitorDataRequest) SetBody(v string) *GetEmonMonitorDataRequest {
	s.Body = &v
	return s
}

func (s *GetEmonMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
