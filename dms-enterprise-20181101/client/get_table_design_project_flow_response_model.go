// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableDesignProjectFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableDesignProjectFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableDesignProjectFlowResponse
	GetStatusCode() *int32
	SetBody(v *GetTableDesignProjectFlowResponseBody) *GetTableDesignProjectFlowResponse
	GetBody() *GetTableDesignProjectFlowResponseBody
}

type GetTableDesignProjectFlowResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableDesignProjectFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableDesignProjectFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableDesignProjectFlowResponse) GoString() string {
	return s.String()
}

func (s *GetTableDesignProjectFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableDesignProjectFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableDesignProjectFlowResponse) GetBody() *GetTableDesignProjectFlowResponseBody {
	return s.Body
}

func (s *GetTableDesignProjectFlowResponse) SetHeaders(v map[string]*string) *GetTableDesignProjectFlowResponse {
	s.Headers = v
	return s
}

func (s *GetTableDesignProjectFlowResponse) SetStatusCode(v int32) *GetTableDesignProjectFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableDesignProjectFlowResponse) SetBody(v *GetTableDesignProjectFlowResponseBody) *GetTableDesignProjectFlowResponse {
	s.Body = v
	return s
}

func (s *GetTableDesignProjectFlowResponse) Validate() error {
	return dara.Validate(s)
}
