// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDestinationCidrBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDestinationCidrBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDestinationCidrBlockResponse
	GetStatusCode() *int32
	SetBody(v *GetDestinationCidrBlockResponseBody) *GetDestinationCidrBlockResponse
	GetBody() *GetDestinationCidrBlockResponseBody
}

type GetDestinationCidrBlockResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDestinationCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDestinationCidrBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDestinationCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *GetDestinationCidrBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDestinationCidrBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDestinationCidrBlockResponse) GetBody() *GetDestinationCidrBlockResponseBody {
	return s.Body
}

func (s *GetDestinationCidrBlockResponse) SetHeaders(v map[string]*string) *GetDestinationCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *GetDestinationCidrBlockResponse) SetStatusCode(v int32) *GetDestinationCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDestinationCidrBlockResponse) SetBody(v *GetDestinationCidrBlockResponseBody) *GetDestinationCidrBlockResponse {
	s.Body = v
	return s
}

func (s *GetDestinationCidrBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
