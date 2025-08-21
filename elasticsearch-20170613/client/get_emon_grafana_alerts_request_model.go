// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonGrafanaAlertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *GetEmonGrafanaAlertsRequest
	GetBody() *string
}

type GetEmonGrafanaAlertsRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmonGrafanaAlertsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEmonGrafanaAlertsRequest) GoString() string {
	return s.String()
}

func (s *GetEmonGrafanaAlertsRequest) GetBody() *string {
	return s.Body
}

func (s *GetEmonGrafanaAlertsRequest) SetBody(v string) *GetEmonGrafanaAlertsRequest {
	s.Body = &v
	return s
}

func (s *GetEmonGrafanaAlertsRequest) Validate() error {
	return dara.Validate(s)
}
