// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVbrFlowTopNResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVbrFlowTopNResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVbrFlowTopNResponse
	GetStatusCode() *int32
	SetBody(v *GetVbrFlowTopNResponseBody) *GetVbrFlowTopNResponse
	GetBody() *GetVbrFlowTopNResponseBody
}

type GetVbrFlowTopNResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVbrFlowTopNResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVbrFlowTopNResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVbrFlowTopNResponse) GoString() string {
	return s.String()
}

func (s *GetVbrFlowTopNResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVbrFlowTopNResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVbrFlowTopNResponse) GetBody() *GetVbrFlowTopNResponseBody {
	return s.Body
}

func (s *GetVbrFlowTopNResponse) SetHeaders(v map[string]*string) *GetVbrFlowTopNResponse {
	s.Headers = v
	return s
}

func (s *GetVbrFlowTopNResponse) SetStatusCode(v int32) *GetVbrFlowTopNResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVbrFlowTopNResponse) SetBody(v *GetVbrFlowTopNResponseBody) *GetVbrFlowTopNResponse {
	s.Body = v
	return s
}

func (s *GetVbrFlowTopNResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
