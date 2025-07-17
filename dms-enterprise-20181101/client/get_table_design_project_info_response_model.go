// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableDesignProjectInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableDesignProjectInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableDesignProjectInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetTableDesignProjectInfoResponseBody) *GetTableDesignProjectInfoResponse
	GetBody() *GetTableDesignProjectInfoResponseBody
}

type GetTableDesignProjectInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableDesignProjectInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableDesignProjectInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableDesignProjectInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTableDesignProjectInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableDesignProjectInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableDesignProjectInfoResponse) GetBody() *GetTableDesignProjectInfoResponseBody {
	return s.Body
}

func (s *GetTableDesignProjectInfoResponse) SetHeaders(v map[string]*string) *GetTableDesignProjectInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTableDesignProjectInfoResponse) SetStatusCode(v int32) *GetTableDesignProjectInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableDesignProjectInfoResponse) SetBody(v *GetTableDesignProjectInfoResponseBody) *GetTableDesignProjectInfoResponse {
	s.Body = v
	return s
}

func (s *GetTableDesignProjectInfoResponse) Validate() error {
	return dara.Validate(s)
}
