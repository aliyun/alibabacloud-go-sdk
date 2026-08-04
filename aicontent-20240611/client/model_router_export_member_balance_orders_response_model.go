// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterExportMemberBalanceOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterExportMemberBalanceOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterExportMemberBalanceOrdersResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterExportMemberBalanceOrdersResponseBody) *ModelRouterExportMemberBalanceOrdersResponse
	GetBody() *ModelRouterExportMemberBalanceOrdersResponseBody
}

type ModelRouterExportMemberBalanceOrdersResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterExportMemberBalanceOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterExportMemberBalanceOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterExportMemberBalanceOrdersResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterExportMemberBalanceOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterExportMemberBalanceOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterExportMemberBalanceOrdersResponse) GetBody() *ModelRouterExportMemberBalanceOrdersResponseBody {
	return s.Body
}

func (s *ModelRouterExportMemberBalanceOrdersResponse) SetHeaders(v map[string]*string) *ModelRouterExportMemberBalanceOrdersResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterExportMemberBalanceOrdersResponse) SetStatusCode(v int32) *ModelRouterExportMemberBalanceOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterExportMemberBalanceOrdersResponse) SetBody(v *ModelRouterExportMemberBalanceOrdersResponseBody) *ModelRouterExportMemberBalanceOrdersResponse {
	s.Body = v
	return s
}

func (s *ModelRouterExportMemberBalanceOrdersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
