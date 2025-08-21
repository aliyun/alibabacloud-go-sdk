// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonGrafanaDashboardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *GetEmonGrafanaDashboardsRequest
	GetBody() *string
}

type GetEmonGrafanaDashboardsRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmonGrafanaDashboardsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEmonGrafanaDashboardsRequest) GoString() string {
	return s.String()
}

func (s *GetEmonGrafanaDashboardsRequest) GetBody() *string {
	return s.Body
}

func (s *GetEmonGrafanaDashboardsRequest) SetBody(v string) *GetEmonGrafanaDashboardsRequest {
	s.Body = &v
	return s
}

func (s *GetEmonGrafanaDashboardsRequest) Validate() error {
	return dara.Validate(s)
}
