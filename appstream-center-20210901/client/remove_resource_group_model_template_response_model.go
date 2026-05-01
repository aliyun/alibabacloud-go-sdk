// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveResourceGroupModelTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveResourceGroupModelTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveResourceGroupModelTemplateResponse
	GetStatusCode() *int32
	SetBody(v *RemoveResourceGroupModelTemplateResponseBody) *RemoveResourceGroupModelTemplateResponse
	GetBody() *RemoveResourceGroupModelTemplateResponseBody
}

type RemoveResourceGroupModelTemplateResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveResourceGroupModelTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveResourceGroupModelTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveResourceGroupModelTemplateResponse) GoString() string {
	return s.String()
}

func (s *RemoveResourceGroupModelTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveResourceGroupModelTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveResourceGroupModelTemplateResponse) GetBody() *RemoveResourceGroupModelTemplateResponseBody {
	return s.Body
}

func (s *RemoveResourceGroupModelTemplateResponse) SetHeaders(v map[string]*string) *RemoveResourceGroupModelTemplateResponse {
	s.Headers = v
	return s
}

func (s *RemoveResourceGroupModelTemplateResponse) SetStatusCode(v int32) *RemoveResourceGroupModelTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveResourceGroupModelTemplateResponse) SetBody(v *RemoveResourceGroupModelTemplateResponseBody) *RemoveResourceGroupModelTemplateResponse {
	s.Body = v
	return s
}

func (s *RemoveResourceGroupModelTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
